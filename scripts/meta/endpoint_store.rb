# frozen_string_literal: true

require_relative 'go_post_authority'

# EndpointStore holds endpoint data by api
class EndpointStore
  include GoPostAuthority

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
