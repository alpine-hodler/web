require 'minitest/autorun'
require_relative '../scheme'

describe Scheme do
  before do
    @filename = "#{File.dirname(__FILE__)}/model/test_model_1.json"
    @scheme = Scheme.new(@filename)
  end

  describe 'when reading the filename' do
    it 'must respond with initialized filename' do
      _(@scheme.filename).must_equal(@filename)
    end
  end

  describe 'when parsing json from the filename' do
    it 'must respond with a hash of data' do
      _(@scheme.api).must_equal('test_api')

      _(@scheme.description).must_equal('Lorem ipsum dolor sit amet, ' \
        'consectetuer adipiscing elit. Aenean commodo ligula eget dolor. ' \
        'Aenean massa. Cum sociis natoque penatibus et magnis dis parturient ' \
        'montes, nascetur ridiculus mus')

      _(@scheme.go_comment).must_equal('// Lorem ipsum dolor sit amet, ' \
        'consectetuer adipiscing elit. Aenean commodo// ligula eget dolor. ' \
        'Aenean massa. Cum sociis natoque penatibus et magnis dis// ' \
        'parturient montes, nascetur ridiculus mus')

      _(@scheme.go_model_filename).must_equal('test_model.go')
      _(@scheme.go_model_name).must_equal('TestModel')
      _(@scheme.graphql_filename).must_equal('test_model.graphqls')
    end
  end
end
