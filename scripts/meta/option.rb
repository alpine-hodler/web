# frozen_string_literal: true

require_relative 'protomodel'

# Model will generate the go files with structs that inherit the protomodel
# fields and methods, but is open to modification.
module Option
  PKG = 'option'
  PKG_DEC = "package #{PKG}"
  MSG = "\n// * This is a generated file, do not edit\n"

  def opinion_model_comment(endpoint)
    go_model_name = endpoint.go_model_name
    "\n" + format_go_comment("#{go_model_name} are options for API requests.")
  end

  def opinion_model_struct(endpoint)
    literals = []
    endpoint.query_params.each do |field|
      literal = field.description.empty? ? '' : "\n// #{field.description}\n"
      literal += "#{field.go_protofield_name} #{field.ptr_go_type}"
      literal += "`json:\"#{field.identifier}\" bson:\"#{field.identifier}\"`"
      literals << literal
    end
    literals = literals.sort.join("\n")
    "\ntype #{endpoint.go_model_name} struct {#{literals}}"
  end

  def opinion_from_prototype(endpoint)
    f_nam = "#{endpoint.go_model_name}FromPrototype"
    f_arg = "(proto *#{Protomodel::PKG}.#{endpoint.graphql_model_name})"
    f_ret = "(opts *#{endpoint.go_model_name})"
    f_sig = "#{f_nam}#{f_arg} #{f_ret}"

    logic = "if proto != nil {\n opts.#{endpoint.graphql_model_name} = *proto\n}\n"
    logic += 'return'
    comment = format_go_comment("#{f_nam} will initialize the exportable options struct from the internal prototype.")

    "\n#{comment}\nfunc #{f_sig} {\n#{logic}\n}\n"
  end

  def opinion_setters(endpoint)
    return unless endpoint.query_params?

    endpoint.query_params.dup.map do |field|
      variable_name = field.go_variable_name
      variable_name = 't' if variable_name == 'type'
      sig_go_type = field.ptr_go_type.include?('[]*') ? "[]*#{field.go_type.dup.gsub('[]', '')}" : field.go_type

      r_ptr = "func (opts *#{endpoint.go_model_name})"
      r_name = "Set#{field.go_protofield_name}"
      r_sig = "#{r_name}(#{variable_name} #{sig_go_type})"
      r_ret = "*#{endpoint.go_model_name}"
      comment = "#{r_name} sets the #{field.go_protofield_name} field on #{endpoint.go_model_name}."

      logic = "opts.#{field.go_protofield_name} = #{field.ptr_go_variable};"
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
          f.write(opinion_model_comment(endpoint))
          f.write(opinion_model_struct(endpoint))
          # f.write(opinion_from_prototype(endpoint))
          f.write(opinion_setters(endpoint))
        end

        # In addition to fixing imports, goimports also formats your code in the
        # same style as gofmt so it can be used as a replacement for your editor's
        # gofmt-on-save hook.
        `/go/bin/goimports -w #{endpoint.go_query_param_filename}`
      end
    end
  end
end
