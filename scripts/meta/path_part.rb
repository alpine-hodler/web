require 'string_inflection'
using StringInflection

class PathPart
  attr_reader \
    :name,
    :go_var,
    :go_arg

  def initialize(name)
    @name = name
    @go_var = self.name.to_camel

    set_go_arg
  end

	def param_go_var
		return '' unless path_param?

		name.dup.gsub('}', '').gsub('{', '').to_camel
	end

	def param_name
		name.dup.gsub('}', '').gsub('{', '')
	end

  def path_param?
    return true if name.include?('{') && name.include?('}')

    false
  end

  private

  def set_go_arg
    @go_arg = if path_param?
								"params[\"#{name.dup.gsub!('{', '').gsub('}', '')}\"]"
              else
                "\"#{name}\""
              end
  end
end
