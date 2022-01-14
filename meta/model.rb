# frozen_string_literal: true

require_relative 'protomodel'

# Model will generate the go files with structs that inherit the protomodel
# fields and methods, but is open to modification.
module Model
  PKG = 'model'
  PKG_DEC = "package #{PKG}"
  MSG = "\n// * This file was initialized by the meta-program, but is open to modification\n"

  def model_comment
    "\n" + go_comment
  end

  def model_struct
    "\ntype #{go_model_name} struct {#{Protomodel::PKG}.#{go_model_name}}"
  end

  def write_model
    Dir.chdir(PARENT_DIR + '/model') do
      File.open(go_model_filename, 'w') do |f|
        f.write(PKG_DEC)
        f.write(MSG)
        f.write('import "github.com/alpine-hodler/sdk/protomodel"')
        f.write(model_comment)
        f.write(model_struct)
      end

      # In addition to fixing imports, goimports also formats your code in the
      # same style as gofmt so it can be used as a replacement for your editor's
      # gofmt-on-save hook.
      `/go/bin/goimports -w #{go_model_filename}`
    end
  end

  def self.protomodel_accesor_comment(scheme, field)
    scheme.format_go_comment("#{field.go_field_name} converts the native protomodel"\
      " #{field.go_protofield_name} to a local #{field.go_type}.\n")
  end

  def self.protomodel_accessor(scheme)
    reciever = "m *#{scheme.go_model_name}"
    scheme.fields.dup.map do |field|
      next unless field.go_struct?

      return_type = ''
      logic = ''
      if field.go_type.include?('[]')
        return_type = "(slice #{field.go_type})"
        base_model = field.go_type.dup.gsub('[]', '').gsub('*', '')
        append_slice = "slice = append(slice, &#{base_model}{#{base_model}: *v})"
        logic += "for _, v := range m.#{field.go_protofield_name} {#{append_slice}};"
        logic += 'return'
      else
        return_type = "*#{field.go_type}"
        logic = "return &#{field.go_type}{#{field.go_type}: m.#{field.go_protofield_name}}"
      end

      comment = Model.protomodel_accesor_comment(scheme, field)
      "\n#{comment}\nfunc (#{reciever}) #{field.go_field_name}() #{return_type} {#{logic}};"
    end.join('')
  end

  def self.protomodel_accessors(schema)
    schema.dup.map { |scheme| Model.protomodel_accessor(scheme) }.join("\n\n")
  end

  def self.write_protomodel_accessors(schema)
    Dir.chdir(PARENT_DIR + '/model') do
      File.open('protomodel_accessors.go', 'w') do |f|
        f.write(PKG_DEC)
        f.write(Protomodel::MSG)
        f.write(Model.protomodel_accessors(schema))
      end

      # In addition to fixing imports, goimports also formats your code in the
      # same style as gofmt so it can be used as a replacement for your editor's
      # gofmt-on-save hook.
      `/go/bin/goimports -w protomodel_accessors.go`
    end
  end
end
