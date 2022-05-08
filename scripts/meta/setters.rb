module Setters
  def option_setters(endpoint)
    return unless endpoint.query_params?

    endpoint.query_params.dup.map do |field|
      variable_name = field.go_variable_name
      variable_name = 't' if variable_name == 'type'
      sig_go_type = field.ptr_go_type.include?('[]*') ? "[]*#{field.go_type.dup.gsub('[]', '')}" : field.go_type
			struct = "#{endpoint.go_model_name}Options"

      r_ptr = "func (opts *#{struct})"
      r_name = "Set#{field.go_protofield_name}"
      r_sig = "#{r_name}(#{variable_name} #{sig_go_type})"
      r_ret = "*#{struct}"
      comment = "#{r_name} sets the #{field.go_protofield_name} field on #{struct}."

      logic = "opts.#{field.go_protofield_name} = #{field.ptr_go_variable};"
      logic += 'return opts'

      "\n#{format_go_comment(comment)}\n#{r_ptr} #{r_sig} #{r_ret} {\n#{logic}\n}\n"
    end.join('')
  end
end
