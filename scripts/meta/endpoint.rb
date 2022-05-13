# frozen_string_literal: true

require 'string_inflection'
using StringInflection

require_relative 'path_part'
require_relative 'field'

# Endpoint holds state concerning endpoints given by the mete/schema json files
class Endpoint
  attr_reader \
    :hash,
    :path,
    :enum_root,
    :go_const,
    :path_parts,
    :description,
    :query_params,
    :graphql_model_name,
    :go_model_name,
    :graphql_query_param_filename,
    :go_query_param_filename,
    :slice,
    :body,
    :return_type,
    :http_method

  def initialize(api, hash)
    return if hash.nil?

    @hash = hash
    @path = hash[:path]
    @enum_root = hash[:enumRoot].to_camel
    @go_const = "#{enum_root.to_camel}Endpoint"
    @description = hash[:description] || ''
    @slice = hash[:slice]
    @return_type = hash[:returnType]
    @http_method = hash[:httpMethod]

    gql_base = "#{api}_#{enum_root.to_snake}_options"
    @graphql_model_name = gql_base.to_pascal
    @go_model_name = enum_root.to_snake.to_s.to_pascal
    @graphql_query_param_filename = "#{gql_base}.graphqls"
    @go_query_param_filename = "#{api}_#{enum_root.to_snake}.go"

    set_path_parts
    set_query_params
    set_body
  end

  def all_params
    @query_params + @body
  end

  def body?
    !@body.empty?
  end

  def path_parts?
    !@path_parts.empty?
  end

  def params?
    body? || query_params?
  end

  def path_params
    return [] if path_parts.empty?

    path_parts.map { |part| part.path_param? ? part : nil }.compact
  end

  def query_params?
    !@query_params.empty?
  end

  private

  def set_path_parts
    @path_parts = []
    first_part_set = false
    path.split('/').each do |part|
      next if part == ''

      part = '/' + part unless first_part_set
      first_part_set = true
      @path_parts << PathPart.new(part)
    end
  end

  def set_body
    @body = []
    (hash[:body] || []).each do |subhash|
      @body << Field.new(subhash)
    end
  end

  def set_query_params
    @query_params = []
    (hash[:queryParams] || []).each do |subhash|
      @query_params << Field.new(subhash)
    end
  end
end
