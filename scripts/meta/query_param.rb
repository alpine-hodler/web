require 'string_inflection'
using StringInflection

class QueryParam
  attr_reader \
    :identifier,
    :go_type,
    :go_protofield_name,
    :go_variable_name,
    :required

  def initialize(hash)
    @identifier = hash[:identifier]
    @go_type = hash[:goType]
		@go_variable_name = @identifier.to_camel
		set_go_protofield_name(hash)
    @required = hash[:required] || false
  end

	private

	def set_go_protofield_name hash
		@go_protofield_name = hash[:identifier].to_pascal

		last_two = @go_protofield_name[-2..-1] || @go_protofield_name
		return unless last_two == "Id"

		@go_protofield_name[-1] = "D"
		@go_protofield_name[-2] = "I"
	end
end
