# frozen_string_literal: true

require_relative 'comment'
require 'string_inflection'
using StringInflection

# GoPostAuthority is responsible for methods that generate the endpoint.go code in
# sdk packages
module GoPostAuthority
  MSG = "\n// * This is a generated file, do not edit"

  def self.endpoint_consts(endpoints)
    consts = ["_ #{POST_AUTHORITY_TYPE_ALIAS} = iota"] | endpoints.dup.map { |ep| ep.go_const }.sort
    "const(#{consts.join(';')})"
  end

  def get_function(endpoints)
    mappings = endpoints.dup.map { |ep| "#{ep.go_const}: get#{ep.enum_root.to_pascal}PostAuthority," }
    comment = Comment.u_format_go_comment("Get takes an #{POST_AUTHORITY_TYPE_ALIAS} const and #{POST_AUTHORITY_TYPE_ALIAS} arguments to parse the URL #{POST_AUTHORITY_TYPE_ALIAS} path.")
    rec = "\n#{comment}\nfunc (#{POST_AUTHORITY_ALIAS} #{POST_AUTHORITY_TYPE_ALIAS})"
    sig = "PostAuthority(#{URI_BUILDER_ALIAS} #{TOOLS_PKG}.URIBuilder) string"

    map = "map[#{POST_AUTHORITY_TYPE_ALIAS}]func(#{URI_BUILDER_ALIAS} #{TOOLS_PKG}.URIBuilder) string"
    wrapper = "#{map}{\n#{mappings.join("\n")}\n}"

    "#{rec} #{sig} {return #{wrapper}[#{POST_AUTHORITY_ALIAS}](#{URI_BUILDER_ALIAS})};"
  end

  def self.pkg_name(api)
    "package #{api}"
  end

  def join_paths(endpoint)
    "path.Join(#{endpoint.path_parts.dup.map { |pp| pp.go_arg }.join(',')})"
  end

  def path_functions(endpoints)
    endpoints.dup.map do |endpoint|
      var = '_'
      var = URI_BUILDER_ALIAS if endpoint.path_parts? || endpoint.query_params?

      sig = "func get#{endpoint.enum_root.to_pascal}PostAuthority(#{var} #{TOOLS_PKG}.URIBuilder) "
      inner_logic = ''
      if endpoint.query_params?
        sig += '(p string)'
        inner_logic = 	"p = #{join_paths(endpoint)};"
        inner_logic += 'var sb strings.Builder;'
        inner_logic += 'sb.WriteString(p);'
        inner_logic += "sb.WriteString(#{URI_BUILDER_ALIAS}.QueryPath().String());"
        inner_logic += 'return sb.String();'
      else
        sig += 'string'
        inner_logic = "return #{join_paths(endpoint)}"
      end

      logic = "{\n#{inner_logic}}"
      desc = Comment.u_format_go_comment(endpoint.description)
      "\n#{desc}\n #{sig} #{logic}"
    end.join("\n\n")
  end

  def write_sdk
    all.each do |api, endpoints|
      Dir.chdir(PARENT_DIR + "/#{api}") do
        File.open(POST_AUTHORITY_FILENAME, 'w') do |f|
          f.write(GoPostAuthority.pkg_name(api))
          f.write("\nimport \"github.com/alpine-hodler/sdk/internal/client\";")
					f.write("\nimport \"github.com/alpine-hodler/sdk/tools\";")
          f.write(MSG)
          f.write("\ntype #{POST_AUTHORITY_TYPE_ALIAS}  uint8;")
          f.write(GoPostAuthority.endpoint_consts(endpoints))
          f.write(path_functions(endpoints))
          f.write(get_function(endpoints))
        end

        # In addition to fixing imports, goimports also formats your code in the
        # same style as gofmt so it can be used as a replacement for your editor's
        # gofmt-on-save hook.
        `/go/bin/goimports -w #{POST_AUTHORITY_FILENAME}`
      end
    end
  end
end
