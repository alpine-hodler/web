# frozen_string_literal: true

require 'json'
require 'json-schema'
require 'string_inflection'
using StringInflection

require_relative 'comment'
require_relative 'field'
require_relative 'endpoint'
require_relative 'protomodel'
require_relative 'model'
require_relative 'graph_schema'

# Scheme is the class encapsulation of a single json file in the meta/schema
# directory
class Scheme
  attr_reader \
    :api,
    :description,
    :go_comment,
    :go_model_filename,
    :go_model_name,
    :model,
    :model_only,
    :filename,
    :graphql_comment,
    :graphql_filename,
    :fields,
    :endpoints,
    :go_model_variable_name

  include Comment
  include Protomodel
  include Model
  include GraphSchema

  def initialize(filename)
    file = File.read(filename)
    hash = JSON.parse(file, symbolize_names: true)
    validate(hash)

    @api = hash[:api]
    @description = hash[:modelDescription]
    @filename = filename
    @model = hash[:model].to_s
    @model_only = hash[:modelOnly] || false

    @go_comment = format_go_comment(@description)
    @go_model_filename = "#{@model}.go"
    @go_model_name = @model.to_pascal
    @go_model_variable_name = @go_model_name.to_camel

    @graphql_comment = format_graphql_comment(@description)
    @graphql_filename = "#{@model}.graphqls"

    @fields = hash[:modelFields].map { |field| Field.new(field) }
    @endpoints = (hash[:endpoints] || []).map { |ep| Endpoint.new(api, ep) }
  end

  def validate(hash)
    schema = JSON.parse(File.read("#{File.dirname(__FILE__)}/schema/schema.json"))
    e = JSON::Validator.fully_validate(schema, hash)
    raise "Schema Error: #{e}" unless e.empty?
  end
end
