# frozen_string_literal: true

require 'string_inflection'
require 'yaml'

require_relative 'model'

using StringInflection

# GraphSchema will generate the graphqls files with types and inputs that
# inherit defined in the schema JSON files.
module GraphSchema
  MSG = "# * This is a generated file, do not edit\n\n"

  def graphql_comment
    "\n" + graphql_comment
  end

  def graphql_type_map(_field)
    {
      'string' => 'String',
      'float64' => 'Float',
      'bool' => 'Boolean',
      'time.Time' => 'Time',
      'int' => 'Int',
      '[]string' => '[String]'
    }
  end

  def graphql_type(field)
    return field.go_type.dup.gsub('scalar.', '') if field.go_type.include?('scalar')
    return graphql_type_map(field)[field.go_type] unless graphql_type_map(field)[field.go_type].nil?
    return "[#{field.go_type.dup.gsub('[]', '').gsub('*', '')}]" if field.go_type.include?('[]')

    field.go_type
  end

  def fmt_graphql_type(field)
    type = graphql_type(field)
    type += '!' if field.required
    type
  end

  def graphql_struct
    return if fields.nil?

    pfields = fields.dup.map { |f| "  #{f.identifier.to_camel}: #{fmt_graphql_type(f)}" }
    "\ntype #{go_model_name} {\n#{pfields.join("\n")}\n}"
  end

  def update_gqlgen
    filename = PARENT_DIR + '/gqlgen.yml'
    configs = YAML.load_file(filename)
    models_key = 'models'
    configs[models_key] = {} if configs[models_key].nil?
    configs[models_key][go_model_name] = {
      'model' => "github.com/alpine-hodler/sdk/pkg/#{Model::PKG}.#{go_model_name}"
    }
    File.open(filename, 'w') { |f| f.write configs.to_yaml }
  end

  def graphql_input(endpoint)
    params = endpoint.query_params.dup.map do |qp|
      "  #{qp.go_protofield_name.to_camel}: #{fmt_graphql_type(qp)}"
    end
    "input #{endpoint.graphql_model_name} {\n#{params.join("\n")}\n}"
  end

  def write_graph_schema
    Dir.chdir(PARENT_DIR + '/graph/schema') do
      File.open(graphql_filename, 'w') do |f|
        f.write(MSG)
        f.write(graphql_comment)
        f.write(graphql_struct)
      end
    end
  end

  def write_graphql_inputs
    Dir.chdir(PARENT_DIR + '/graph/schema') do
      all.each do |_api, endpoints|
        endpoints.each do |endpoint|
          next unless endpoint.query_params?

          File.open(endpoint.graphql_query_param_filename, 'w') do |f|
            f.write(MSG)
            f.write(graphql_input(endpoint))
          end
        end
      end
    end
  end
end
