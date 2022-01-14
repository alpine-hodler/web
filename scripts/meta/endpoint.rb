# frozen_string_literal: true

require 'string_inflection'
using StringInflection

require_relative 'path_part'
require_relative 'query_param'

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
    :graphql_query_param_filename

  def initialize(api, hash)
    return if hash.nil?

    @hash = hash
    @path = hash[:path]
    @enum_root = hash[:enumRoot]
    @go_const = "#{enum_root.to_pascal}Endpoint"
    @description = hash[:description] || ''

    gql_base = "#{api}_#{enum_root.to_snake}_options"
    @graphql_model_name = gql_base.to_pascal
    @graphql_query_param_filename = "#{gql_base}.graphqls"

    set_path_parts
    set_query_params
  end

  def part_paths?
    !@path_parts.empty?
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

  def set_query_params
    @query_params = []
    (hash[:queryParams] || []).each do |subhash|
      @query_params << QueryParam.new(subhash)
    end
  end
end
