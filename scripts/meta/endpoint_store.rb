# frozen_string_literal: true

require_relative 'sdk_endpoints'
require_relative 'graph_schema'

# EndpointStore holds endpoint data by api
class EndpointStore
  include SDKEndpoints
	include GraphSchema

  def initialize
    tree = proc { Hash.new { |hash, key| hash[key] = [] } }
    @store = tree.call
  end

  def add(scheme)
    scheme.endpoints.each { |ep| @store[scheme.api] << ep }
  end

  def all
    @store
  end
end
