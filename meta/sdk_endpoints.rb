# frozen_string_literal: true

require_relative 'comment'

# SDKEndpoints is responsible for methods that generate the endpoint.go code in
# sdk packages
module SDKEndpoints
  ENDPOINT_TYPE = "\ntype Endpoint int;"
  MSG = "\n// * This is a generated file, do not edit"
  GET_MSG = 'Get takes an endpoint const and endpoint arguments to parse the URL endpoint path.'

  def self.endpoint_consts(endpoints)
    consts = ['_ Endpoint = iota'] | endpoints.dup.map { |ep| ep.go_const }.sort
    "const(#{consts.join(';')})"
  end

  def get_function(endpoints)
    mappings = endpoints.dup.map { |ep| "#{ep.go_const}: #{ep.enum_root}Path," }
    comment = Comment.u_format_go_comment(GET_MSG)
    rec = "\n#{comment}\nfunc (endpoint Endpoint)"
    sig = 'Path(args client.EndpointArgs) string'

    map = 'map[Endpoint]func(args client.EndpointArgs) string'
    wrapper = "#{map}{\n#{mappings.join("\n")}\n}"

    "#{rec} #{sig} {return #{wrapper}[endpoint](args)};"
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
      var = 'args' if endpoint.part_paths? || endpoint.query_params?

      sig = "func #{endpoint.enum_root}Path(#{var} client.EndpointArgs) "
      inner_logic = ''
      if endpoint.query_params?
        sig += '(p string)'
        inner_logic = 	"p = #{join_paths(endpoint)};"
        inner_logic += 'var sb strings.Builder;'
        inner_logic += 'sb.WriteString(p);'
        inner_logic += 'sb.WriteString(args.QueryPath().String());'
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
        File.open('endpoint.go', 'w') do |f|
          f.write(SDKEndpoints.pkg_name(api))
          f.write(MSG)
          f.write(ENDPOINT_TYPE)
          f.write(SDKEndpoints.endpoint_consts(endpoints))
          f.write(path_functions(endpoints))
          f.write(get_function(endpoints))
        end

        # In addition to fixing imports, goimports also formats your code in the
        # same style as gofmt so it can be used as a replacement for your editor's
        # gofmt-on-save hook.
        `/go/bin/goimports -w endpoint.go`
      end
    end
  end
end
