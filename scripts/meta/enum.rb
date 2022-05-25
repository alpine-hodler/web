require_relative 'inflector'

class Enum
  attr_reader \
    :identifier,
    :values,
		:go_type_name,
		:go_var_name,
		:pluralize,
		:pluralize_var

	include Inflector

  def initialize(hash = {})
    @identifier = hash[:identifier]
		@pluralize = inflector.camelize_upper(hash[:pluralize].dup.gsub('.', '_'))
		@pluralize_var = inflector.camelize_lower(pluralize)
		@go_type_name = inflector.camelize_upper(hash[:identifier].dup.gsub('.', '_'))
		@go_var_name = inflector.camelize_lower(go_type_name)
    @values = hash[:values].map { |val| [val, go_type_name + inflector.camelize_upper(val.dup.gsub('.', '_'))] }
  end
end
