# frozen_string_literal: true

# TypesWriter will write the types file.
module TypesWriter
  def self.apis(types)
    tree = (proc { Hash.new { |hash, key| hash[key] = [] } }).call
    types.each { |t| tree[t.api] << t }
    tree
  end

  def self.stringer(api_types)
    api_types.dup.map do |t|
      t.enums.dup.map do |enum|
        logic = "\nfunc (#{enum.go_var_name} *#{enum.go_type_name}) String() string {;"
        logic += "if #{enum.go_var_name} != nil { return string(*#{enum.go_var_name}) };"
        logic += "return \"\"; };\n"
        unless enum.pluralize.nil?
          logic += "\nfunc (#{enum.pluralize_var} *#{enum.pluralize}) String() string {;"
          logic += 'var str string;'
          logic += "if #{enum.pluralize_var} != nil {;"
          logic += 'slice := []string{};'
          logic += "for _, val := range *#{enum.pluralize_var} {;"
          logic += 'slice = append(slice, val.String());'
          logic += '};'
          logic += 'str = strings.Join(slice, ",");'
          logic += '};'
          logic += "return str;};\n"
        end
        logic
      end
    end.flatten.sort.join('')
  end

  def self.enum_structures api_types
    api_types.dup.map do |t|
      t.enums.dup.map do |enum|
        const_insides = enum.values.dup.map { |val| "#{val[1]} #{enum.go_type_name} = \"#{val[0]}\"" }

        structure = "type #{enum.go_type_name} string;"
        structure += "type #{enum.pluralize} []#{enum.go_type_name};"
        structure += "const (#{const_insides.join("\n")});"
        structure
      end
    end.flatten.sort.join("\n")
  end

  def self.write(types)
    apis(types).each do |api, api_types|
      path = Pathname.new(PARENT_DIR).join(api)
      Dir.chdir(path.to_s) do
        File.open('types.go', 'w') do |f|
          f.write("package #{api}")
          f.write(GEN_MSG)
          f.write(enum_structures(api_types))
          f.write(stringer(api_types))
        end
        `/go/bin/goimports -w types.go`
      end
    end
  end
end
