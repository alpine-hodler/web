require 'string_inflection'
using StringInflection

class QueryParam
  attr_reader \
    :identifier,
    :go_type,
    :go_protofield_name,
    :required

  def initialize(hash)
    @identifier = hash[:identifier]
    @go_type = hash[:goType]
    @go_protofield_name = @identifier.to_pascal
    @required = hash[:required] || false
  end
end
