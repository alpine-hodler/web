#!/usr/local/bin/ruby -w
# frozen_string_literal: true

require_relative 'scheme'
require_relative 'protomodel'
require_relative 'model'
require_relative 'endpoint_store'
require_relative 'option'

PARENT_DIR = File.expand_path('..', Dir.pwd)

def generate_models
  schema = []
  endpoint_store = EndpointStore.new
  Dir.glob("#{File.dirname(__FILE__)}/schema/*.json").map do |filename|
    next if filename.include?('/schema.json')

    scheme = Scheme.new(filename)
    endpoint_store.add(scheme) unless scheme.model_only
    scheme.write_protomodel
    scheme.write_model
    scheme.write_option

    # TODO: fully deprecate support for graphql in the sdk.  We keep the functionality around for now in case we want to
    # TODO use it to generate files in a graphql aux repo.
    # scheme.write_graph_schema
    # scheme.update_gqlgen
    schema << scheme
  end

  Model.write_protomodel_accessors(schema)

  endpoint_store.write_sdk
  # endpoint_store.write_graphql_inputs
end

generate_models
