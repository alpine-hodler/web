require 'minitest/autorun'
require_relative '../comment'

describe Comment do
  include Comment

  describe 'when formatting comments' do
    it 'must have one line for the the exact character count' do
      list = split_go_comment_by_char('Loremoo', 10)
      assert_equal(['// Loremoo'], list)
    end

    it 'must have two line for {count} <= len <= {2 * count}' do
      list = split_go_comment_by_char('Lorem ipsum', 10)
      assert_equal(['// Lorem', '// ipsum'], list)
    end
  end
end
