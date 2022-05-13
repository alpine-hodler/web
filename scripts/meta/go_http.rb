# frozen_string_literal: true

require 'string_inflection'
using StringInflection

# HTTP will build the HTTP client methods for a scheme.
module GoHTTP
  private

  def get_http_method(endpoint)
    {
      'GET' => 'Get',
      'POST' => 'Post',
      'DELETE' => 'Delete'
    }[endpoint.http_method]
  end

  def path_param_args(endpoint)
    return nil unless endpoint.path_parts?

    endpoint.path_params.dup.map { |part| "#{part.param_go_var.to_camel} string" }
  end

  def option_arg(endpoint)
    return nil unless endpoint.params?

    "opts *#{endpoint.go_model_name}Options"
  end

  def sig_args(endpoint)
    args = [path_param_args(endpoint), option_arg(endpoint)].flatten.compact
    return "(#{args[0]})" if args.length == 1

    "(#{args.join(',')})"
  end

  def return_args(endpoint)
    val = if endpoint.return_type.nil?
            endpoint.slice ? "[]*#{go_model_name}" : "*#{go_model_name}"
          else
            endpoint.return_type
          end

    "(#{RETURL_ALIAS} #{val}, _ error)"
  end

  def fetch_call(_endpoint)
    "Fetch().Assign(&#{RETURL_ALIAS}).JoinMessages()"
  end

  def http_call(endpoint)
    "#{CLIENT_ALIAS}.#{get_http_method(endpoint)}(#{endpoint.go_const})"
  end

  def path_call endpoint
    # PathParam("account_id", accountId).
    fns = endpoint.path_params.dup.map { |part| "PathParam(\"#{part.param_name}\", #{part.param_go_var.to_camel})" }
    return nil if fns.empty?
    return fns[0] if fns.length == 1

    fns.join('.')
  end

  def query_params_call endpoint
    return nil unless endpoint.query_params?

    "SetQueryParams(#{OPTIONS_ALIAS})"
  end

  def body_call endpoint
    return nil unless endpoint.body?

    "SetBody(client.BodyTypeJSON, #{OPTIONS_ALIAS})"
  end

  def return_stmt endpoint
    calls = [
      http_call(endpoint),
      path_call(endpoint),
      query_params_call(endpoint),
      body_call(endpoint),
      fetch_call(endpoint)
    ].flatten.compact
    "return #{RETURL_ALIAS}, #{calls.join(".\n")}"
  end

  public

  # client_receivers will generate all of the Go `cleint` struct receivers from a scheme's endpoints.
  def client_receivers
    return nil if endpoints.empty?

    receiver_alias = "(#{CLIENT_ALIAS} *#{CLIENT_STRUCT_NAME})"
    endpoints.dup.map do |endpoint|
      r = "\n" + format_go_comment(endpoint.description)
      r += "\nfunc"
      r += receiver_alias
      r += ' ' + endpoint.enum_root.to_pascal
      r += sig_args(endpoint)
      r += ' ' + return_args(endpoint)
      r += ' ' + "{\n #{return_stmt(endpoint)} \n}"

      { receiver: r, name: endpoint.enum_root }
    end
  end
end
