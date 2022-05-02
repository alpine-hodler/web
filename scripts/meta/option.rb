# frozen_string_literal: true

require_relative 'protomodel'

# Model will generate the go files with structs that inherit the protomodel
# fields and methods, but is open to modification.
module Option
  PKG = 'option'
  PKG_DEC = "package #{PKG}"
  MSG = "\n// * This is a generated file, do not edit\n"

  def o_model_comment(endpoint)
    go_model_name = endpoint.go_model_name
    "\n" + format_go_comment("#{go_model_name} are options for API requests.")
  end

  def o_model_struct(endpoint)
    "\ntype #{endpoint.go_model_name} struct {#{Protomodel::PKG}.#{endpoint.graphql_model_name}}"
  end

  def o_from_prototype(endpoint)
    f_nam = "#{endpoint.go_model_name}FromPrototype"
    f_arg = "(proto *#{Protomodel::PKG}.#{endpoint.graphql_model_name})"
    f_ret = "(opts *#{endpoint.go_model_name})"
    f_sig = "#{f_nam}#{f_arg} #{f_ret}"

		logic = "if proto != nil {\n opts.#{endpoint.graphql_model_name} = *proto\n}\n"
		logic += "return"
    comment = format_go_comment("#{f_nam} will initialize the exportable options struct from the internal prototype.")

    "\n#{comment}\nfunc #{f_sig} {\n#{logic}\n}\n"
  end

	def o_setters(endpoint)
    return unless endpoint.query_params?

    endpoint.query_params.dup.map do |query_param|
      variable_name = query_param.go_variable_name
      variable_name = 't' if variable_name == 'type'

      go_type = if query_param.go_type.include?('[]')
                  "[]*#{query_param.go_type.dup.gsub('[]', '')}"
                else
                  query_param.go_type
                end

      r_ptr = "func (opts *#{endpoint.go_model_name})"
      r_name = "Set#{query_param.go_protofield_name}"
      r_sig = "#{r_name}(#{variable_name} #{go_type})"
      r_ret = "*#{endpoint.go_model_name}"
      comment = "#{r_name} sets the #{query_param.go_protofield_name} field on #{endpoint.go_model_name}."

      if query_param.required
        logic = "opts.#{query_param.go_protofield_name} = #{variable_name};"
      else
        prefix = go_type.include?('[]') ? '' : '&'
        logic = "opts.#{query_param.go_protofield_name} = #{prefix}#{variable_name};"
      end
      logic += 'return opts'

      "\n#{format_go_comment(comment)}\n#{r_ptr} #{r_sig} #{r_ret} {\n#{logic}\n}\n"
    end.join('')
  end


  def write_option
    Dir.chdir(PARENT_DIR + '/option') do
      endpoints.each do |endpoint|
        next unless endpoint.query_params?

        File.open(endpoint.go_query_param_filename, 'w') do |f|
          f.write(PKG_DEC)
          f.write(MSG)
          f.write('import "github.com/alpine-hodler/sdk/internal/protomodel";')
					f.write('import "github.com/alpine-hodler/sdk/pkg/scalar";')
					f.write('import "time";')
          f.write(o_model_comment(endpoint))
          f.write(o_model_struct(endpoint))
          f.write(o_from_prototype(endpoint))
					f.write(o_setters(endpoint))
        end

        # In addition to fixing imports, goimports also formats your code in the
        # same style as gofmt so it can be used as a replacement for your editor's
        # gofmt-on-save hook.
        `/go/bin/goimports -w #{endpoint.go_query_param_filename}`
      end
    end
  end

  # def self.protomodel_accesor_comment(scheme, field)
  #   scheme.format_go_comment("#{field.go_field_name} converts the native protomodel"\
  #     " #{field.go_protofield_name} to a local #{field.go_type}.\n")
  # end

  # def self.protomodel_accessor(scheme)
  #   reciever = "m *#{scheme.go_model_name}"
  #   scheme.fields.dup.map do |field|
  #     next unless field.go_struct?

  #     return_type = ''
  #     logic = ''
  #     if field.go_type.include?('[]')
  #       return_type = "(slice #{field.go_type})"
  #       base_model = field.go_type.dup.gsub('[]', '').gsub('*', '')
  #       append_slice = "slice = append(slice, &#{base_model}{#{base_model}: *v})"
  #       logic += "for _, v := range m.#{field.go_protofield_name} {#{append_slice}};"
  #       logic += 'return'
  #     else
  #       return_type = "*#{field.go_type}"
  #       logic = "return &#{field.go_type}{#{field.go_type}: m.#{field.go_protofield_name}}"
  #     end

  #     comment = Model.protomodel_accesor_comment(scheme, field)
  #     "\n#{comment}\nfunc (#{reciever}) #{field.go_field_name}() #{return_type} {#{logic}};"
  #   end.join('')
  # end

  # def self.protomodel_accessors(schema)
  #   schema.dup.map { |scheme| Model.protomodel_accessor(scheme) }.join("\n\n")
  # end

  # def self.write_protomodel_accessors(schema)
  #   Dir.chdir(PARENT_DIR + '/model') do
  #     File.open('protomodel_accessors.go', 'w') do |f|
  #       f.write(PKG_DEC)
  #       f.write(Protomodel::MSG)
  #       f.write(Model.protomodel_accessors(schema))
  #     end

  #     # In addition to fixing imports, goimports also formats your code in the
  #     # same style as gofmt so it can be used as a replacement for your editor's
  #     # gofmt-on-save hook.
  #     `/go/bin/goimports -w protomodel_accessors.go`
  #   end
  # end

  # def self.endpoint_setters(scheme, endpoint)
  #   return unless endpoint.query_params?

  #   endpoint.query_params.dup.map do |query_param|
  #     variable_name = query_param.go_variable_name
  #     variable_name = 't' if variable_name == 'type'

  #     go_type = if query_param.go_type.include?('[]')
  #                 "[]*#{query_param.go_type.dup.gsub('[]', '')}"
  #               else
  #                 query_param.go_type
  #               end

  #     r_ptr = "func (opts *#{endpoint.graphql_model_name})"
  #     r_name = "Set#{query_param.go_protofield_name}"
  #     r_sig = "#{r_name}(#{variable_name} #{go_type})"
  #     r_ret = "*#{endpoint.graphql_model_name}"
  #     comment = "#{r_name} sets the #{query_param.go_protofield_name} field on #{endpoint.graphql_model_name}."

  #     if query_param.required
  #       logic = "opts.#{query_param.go_protofield_name} = #{variable_name};"
  #     else
  #       prefix = go_type.include?('[]') ? '' : '&'
  #       logic = "opts.#{query_param.go_protofield_name} = #{prefix}#{variable_name};"
  #     end
  #     logic += 'return opts'

  #     "\n#{scheme.format_go_comment(comment)}\n#{r_ptr} #{r_sig} #{r_ret} {\n#{logic}\n}\n"
  #   end.join('')
  # end

  # def self.schema_setters(scheme)
  #   scheme.endpoints.dup.map do |endpoint|
  #     "\n#{Model.endpoint_setters(scheme, endpoint)}\n"
  #   end.join('')
  # end

  # def self.setters(schema)
  #   schema.dup.map { |scheme| Model.schema_setters(scheme) }.join("\n\n")
  # end

  # def self.write_query_param_setters(schema)
  #   Dir.chdir(PARENT_DIR + '/model') do
  #     File.open('setters.go', 'w') do |f|
  #       f.write(PKG_DEC)
  #       f.write("\nimport \"github.com/alpine-hodler/sdk/pkg/scalar\";")
  #       f.write("\nimport \"time\";")
  #       f.write(Protomodel::MSG)
  #       f.write(Model.setters(schema))
  #     end

  #     # In addition to fixing imports, goimports also formats your code in the
  #     # same style as gofmt so it can be used as a replacement for your editor's
  #     # gofmt-on-save hook.
  #     `/go/bin/goimports -w setters.go`
  #   end
  # end
end
