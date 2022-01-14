# frozen_string_literal: true

require 'fileutils'

# Protomodel is to extend scheme with the methods needed to generate the code
# the go protomodels
module Protomodel
  MSG = "\n// * This is a generated file, do not edit\n"
  PKG = 'protomodel'
  PKG_DEC = "package #{PKG}"
  RETURN_ERR	= 'if err != nil { return err }'
  SERIAL_DECLARATION	= "\ndata, err := serial.NewJSONTransform(d); #{RETURN_ERR}"

  def protomodel_comment
    "\n" + go_comment
  end

  def protomodel_struct
    literals = []
    fields.each do |field|
      literal = field.description.empty? ? '' : "\n// #{field.description}\n"
      literal += "#{field.go_protofield_name} #{field.go_type}"
      literal += "`json:\"#{field.identifier}\"`"
      literals << literal
    end
    literals = literals.sort.join("\n")
    "\ntype #{go_model_name} struct {#{literals}}"
  end

  def constantize_json_go_tags
    c = fields.dup.map { |field| "#{field.go_field_tag} = \"#{field.identifier}\"" }
    "const(#{c.join(';')})"
  end

  def struct_access_variable(field)
    "#{go_model_variable_name}.#{field.go_protofield_name}"
  end

  def unmarshal_fn_signature(field)
    "#{field.go_field_tag}, &#{struct_access_variable(field)}"
  end

  def custom_deserializer(field)
    return nil unless !field.deserializer.nil? && field.deserializer == "UnmarshalFloatFromString"

    "\ndata.UnmarshalFloatFromString(#{unmarshal_fn_signature(field)})"
  end

  def time_deserializer(field)
    sig = [field.datetime_layout, field.go_field_tag, "&#{struct_access_variable(field)}"]
    unm = ["err = data.UnmarshalTime(#{sig.join(',')})", RETURN_ERR]
    unm.join(';')
  end

  def scalar_deserializer(field, sig)
    "\ndata.Unmarshal#{field.go_type.dup.gsub('scalar.', '')}(#{sig})"
  end

  def slice_deserializer(field)
    accessor = struct_access_variable(field)
    data_accessor = "data.Value(#{field.go_field_tag})"
    type = field.go_type.dup.gsub!('[]', '').gsub!('*', '')
    marshal_logic = "bytes, _ := json.Marshal(item); obj := #{type}{};"
    unmarshal_logic = 'if err := json.Unmarshal(bytes, &obj); err != nil {return err};'
    logic = "#{marshal_logic} #{unmarshal_logic} #{accessor} = append(#{accessor}, &obj)"
    loop_logic = "for _, item := range #{data_accessor}.([]interface{}) {#{logic}}"

    "\nif v := #{data_accessor}; v != nil {#{loop_logic}}"
  end

  def struct_deserializer(field, sig)
    accessor = struct_access_variable(field)
    init = "#{accessor} = #{field.go_type}{}"

    "\n#{init}; if err := data.UnmarshalStruct(#{sig}); err != nil {return err}"
  end

  def get_deserializer(field, sig)
    {
      'string' => "\ndata.UnmarshalString(#{sig})",
      'bool' => "\ndata.UnmarshalBool(#{sig})",
      'time.Time' => "\n#{time_deserializer(field)}",
      'int' => "\ndata.UnmarshalInt(#{sig})",
      '[]string' => "\ndata.UnmarshalStringSlice(#{sig})",
      'float64' => "\ndata.UnmarshalFloat(#{sig})"
    }[field.go_type]
  end

  def generic_deserializer(field)
    sig = unmarshal_fn_signature(field)
    deserializer = get_deserializer(field, sig)
    return deserializer unless deserializer.nil?
    return scalar_deserializer(field, sig) if field.go_type.include?('scalar')
    return slice_deserializer(field) if field.go_type.include?('[]')

    struct_deserializer(field, sig)
  end

  def deserializers
    fields.dup.map do |field|
      custom_deserializer(field) || generic_deserializer(field)
    end.sort.join(';')
  end

  def deserializer_fn
    fn = [constantize_json_go_tags, SERIAL_DECLARATION, deserializers]

    comment = "UnmarshalJSON will deserialize bytes into a #{go_model_name} model"
    comment = format_go_comment(comment)

    reciever = "\nfunc (#{go_model_variable_name} *#{go_model_name}) "
    signature = "UnmarshalJSON(d []byte) error {#{fn.join('')}; return nil}"
    function = reciever + signature

    "\n" + [comment, function].join('')
  end

  def write_protomodel
    Dir.chdir(PARENT_DIR + '/protomodel') do
      File.open(go_model_filename, 'w') do |f|
        f.write(PKG_DEC)
        f.write(MSG)
        f.write(protomodel_comment)
        f.write(protomodel_struct)
        f.write(deserializer_fn)
      end

      # In addition to fixing imports, goimports also formats your code in the
      # same style as gofmt so it can be used as a replacement for your editor's
      # gofmt-on-save hook.
      `/go/bin/goimports -w #{go_model_filename}`
    end
  end
end
