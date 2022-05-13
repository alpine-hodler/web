# frozen_string_literal: true

require 'json'
require 'json-schema'
require 'string_inflection'
using StringInflection

require_relative 'comment'
require_relative 'field'
require_relative 'endpoint'
require_relative 'unmarshaller'
require_relative 'go_struct'
require_relative 'option'
require_relative 'setters'
require_relative 'go_http'

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
    :fields,
    :endpoints,
    :go_model_variable_name,
		:non_struct,
		:package

  include Comment
  include Unmarshaller
	include Option
	include GoStruct
	include Setters
	include GoHTTP

  def initialize(filename)
    file = File.read(filename)
    hash = JSON.parse(file, symbolize_names: true)
    validate(hash)

    @api = hash[:api]
		@package = "pacakge #{api}"
    @description = hash[:modelDescription]
    @filename = filename
    @model = hash[:model].to_s
    @model_only = hash[:modelOnly] || false
		@non_struct = hash[:nonStruct]

    @go_comment = format_go_comment(@description)
    @go_model_filename = "#{@model}.go"
    @go_model_name = @model.to_pascal
    @go_model_variable_name = @go_model_name.to_camel

    @fields = hash[:modelFields].map { |field| Field.new(field) }
    @endpoints = (hash[:endpoints] || []).map { |ep| Endpoint.new(api, ep) }
  end

  def validate(hash)
    schema = JSON.parse(File.read("#{File.dirname(__FILE__)}/schema/schema.json"))
    e = JSON::Validator.fully_validate(schema, hash)
    raise "Schema Error: #{e}" unless e.empty?
  end
end
