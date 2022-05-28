# frozen_string_literal: true

# Model will generate the go files with structs that inherit the protomodel
# fields and methods, but is open to modification.
module Model
  MSG = "\n// * This is a generated file, do not edit\n"
  MODEL_FILENAME = 'model.go'

  def self.apis(schema)
    tree = (proc { Hash.new { |hash, key| hash[key] = [] } }).call
    schema.each { |scheme| tree[scheme.api] << scheme }
    tree
  end

  def self.structs(schema)
    schema.dup.map { |scheme| scheme.model_struct }.join('')
  end

  def self.unmarshallers(schema)
    schema.dup.map { |scheme| scheme.unmarshaller }.join("\n")
  end

  def self.write(schema)
    apis(schema).each do |api, api_schema|
      path = Pathname.new(PARENT_DIR).join(api)
      Dir.chdir(path.to_s) do
        File.open('models.go', 'w') do |f|
          f.write("package #{api}")
          f.write("\nimport \"github.com/alpine-hodler/web/pkg/scalar\";")
          f.write("\nimport \"github.com/alpine-hodler/web/internal/serial\";")
          f.write(Model::MSG)
          f.write(structs(api_schema))
          f.write(unmarshallers(api_schema))
        end
        `/go/bin/goimports -w models.go`
      end
    end
  end
end
