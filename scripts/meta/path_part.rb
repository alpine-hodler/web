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

  def path_param?
    return true if name.include?('{') && name.include?('}')

    false
  end

  private

  def set_go_arg
    @go_arg = if path_param?
                "*args[\"#{name.dup.gsub!('{', '').gsub('}', '')}\"].PathParam"
              else
                "\"#{name}\""
              end
  end
end
