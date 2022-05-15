module Setters
  private

  def get_body_setter(field)
    sig = {
      'bool' => 'SetBodyBool',
      'float64' => 'SetBodyFloat',
      'int32' => 'SetBodyInt32',
      'int' => 'SetBodyInt',
      'string' => 'SetBodyString',
      'time.Time' => 'SetBodyTime'
    }[field.go_type]

		sig = 'SetStringer' if sig.nil?

    adr = field.required && !field.go_slice? ? '&' : ''
    sig + "(\"#{field.identifier}\", #{adr}#{OPTIONS_ALIAS}.#{field.go_field_name})"
  end

  def get_query_param_setter(field)
    sig = {
      '[]string' => "\nSetQueryParamStrings",

      'bool' => 'SetQueryParamBool',
      'int32' => 'SetQueryParamInt32',
      'int' => 'SetQueryParamInt',
      'string' => 'SetQueryParamString',
      'time.Time' => 'SetQueryParamTime'
    }[field.go_type]

		sig = 'SetQueryParamStringer' if sig.nil?

    adr = field.required && !field.go_slice? ? '&' : ''
    sig + "(\"#{field.identifier}\", #{adr}#{OPTIONS_ALIAS}.#{field.go_field_name})"
  end

  public

  def option_setters(endpoint)
    return unless endpoint.params?

    endpoint.all_params.dup.map do |field|
      variable_name = field.go_variable_name
      variable_name = 't' if variable_name == 'type'
      sig_go_type = field.ptr_go_type.include?('[]*') ? "[]*#{field.go_type.dup.gsub('[]', '')}" : field.go_type
      struct = "#{endpoint.go_model_name}Options"

      r_ptr = "func (#{OPTIONS_ALIAS} *#{struct})"
      r_name = "Set#{field.go_protofield_name}"
      r_sig = "#{r_name}(#{variable_name} #{sig_go_type})"
      r_ret = "*#{struct}"
      comment = "#{r_name} sets the #{field.go_protofield_name} field on #{struct}."
			comment += "  #{field.description}" unless field.description.nil?

      logic = "#{OPTIONS_ALIAS}.#{field.go_protofield_name} = #{field.ptr_go_variable};"
      logic += "return #{OPTIONS_ALIAS}"

      { setter: "\n#{format_go_comment(comment)}\n#{r_ptr} #{r_sig} #{r_ret} {\n#{logic}\n}\n", name: struct }
    end
  end

  def option_body_setter(endpoint)
    return unless endpoint.body?

    name = 'SetBody'
    struct = "#{endpoint.go_model_name}Options"
    sig = "func (#{OPTIONS_ALIAS} *#{struct}) #{name}(req *#{CLIENT_PKG}.Request)"
    setters = endpoint.body.dup.map { |field| get_body_setter(field) }.compact.flatten
		logic = "if #{OPTIONS_ALIAS} == nil { return };"
    logic += "req.#{setters.join(".\n")}"
    { setter: "\n#{sig} {\n#{logic}\n}\n", name: struct }
  end

  def option_query_params_setter(endpoint)
    return unless endpoint.query_params?

    name = 'SetQueryParams'
    struct = "#{endpoint.go_model_name}Options"
    sig = "func (#{OPTIONS_ALIAS} *#{struct}) #{name}(req *#{CLIENT_PKG}.Request)"
    setters = endpoint.query_params.dup.map { |field| get_query_param_setter(field) }.compact.flatten
		logic = "if #{OPTIONS_ALIAS} == nil { return };"
    logic += "req.#{setters.join(".\n")}"
    { setter: "\n#{sig} {\n#{logic}\n}\n", name: struct }
  end
end
