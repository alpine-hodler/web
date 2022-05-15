#!/usr/local/bin/ruby -w
# frozen_string_literal: true

require_relative 'scheme'
require_relative 'model'
require_relative 'post_authority'
require_relative 'option'
require_relative 'go_client'
require_relative 'path_part'

URI_BUILDER_ALIAS = 'builder'
CLIENT_ALIAS = 'c'
CLIENT_FILENAME = 'http.go'
CLIENT_PKG = 'client'
CLIENT_STRUCT_NAME = 'C'
GEN_MSG = "\n// * This is a generated file, do not edit\n"
OPTIONS_ALIAS = 'opts'
OPTIONS_FILENAME = 'options.go'
PARENT_DIR = File.expand_path('..', Dir.pwd)
POST_AUTHORITY_ALIAS = 'pa'
POST_AUTHORITY_FILENAME = 'post_authority.go'
POST_AUTHORITY_TYPE_ALIAS = 'postAuthority'
RETURL_ALIAS = 'm'
TOOLS_PKG = 'tools'

def generate_models
  schema = []
  post_authority = PostAuthority.new
  Dir.glob("#{File.dirname(__FILE__)}/schema/*").map do |dir|
		Dir.glob("#{dir}/*.json").map do |filename|
			scheme = Scheme.new(filename)
			post_authority.add(scheme) unless scheme.model_only
			schema << scheme
		end
  end

  Model.write(schema)
  Option.write(schema)
  GoClient.write(schema)

  post_authority.write_sdk
end

generate_models
