# frozen_string_literal: true

require 'string_inflection'
using StringInflection

# Field holds state concerning endpoints given by the mete/schema json files
class Field
  attr_reader \
    :datetime_layout,
    :deserializer,
    :identifier,
    :go_type,
    :go_field_name,
    :go_field_tag,
    :description,
    :hash,
    :required

  GO_TYPES = %w(
    string
    bool
    time.Time
    int
    []string
    float64
  ).freeze

  def initialize(hash)
    return if hash.nil?

    @hash = hash
    @datetime_layout = hash[:datetimeLayout] || 'time.RFC3339Nano'
    @deserializer = hash[:unmarshaller]
    @identifier = hash[:identifier]
    @go_type = hash[:goType]
    @go_field_name = hash[:identifier].to_pascal
    @go_field_tag = "#{hash[:identifier]}_json_tag".to_camel
    @description = hash[:description] || ''
    @required = hash[:required]
  end

  def go_struct?
    return false if GO_TYPES.include?(hash[:goType])
    return false if go_type.include?('scalar')

    true
  end

  def go_protofield_name
    return go_field_name unless go_struct?

    "Proto#{go_field_name}"
  end
end
