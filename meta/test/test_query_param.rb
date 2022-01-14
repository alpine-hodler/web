require 'minitest/autorun'
require_relative '../query_param'

describe QueryParam do
  before do
    @query_param1 = QueryParam.new({
                                     identifier: 'query_param',
                                     goType: 'string'
                                   })

    @query_param2 = QueryParam.new({
                                     identifier: 'query_param',
                                     goType: 'string',
                                     required: true
                                   })
  end

  describe 'when parsing data from the list' do
    it 'must initialize accessors' do
      _(@query_param1.identifier).must_equal('query_param')
      _(@query_param1.go_type).must_equal('string')
      _(@query_param1.go_protofield_name).must_equal('QueryParam')
      _(@query_param1.required).must_equal(false)

      _(@query_param2.required).must_equal(true)
    end
  end
end
