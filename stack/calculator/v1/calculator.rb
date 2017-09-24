# Public Internal Deprecated: description.
#
# this class still fails because it processes ('1,2,+,+,3,3')
#
#   example
class Calculator
  NUMERIC = /^\d+$/
  SUM_OPERATOR = /^\+$/

  def initialize
    @numbers_stack = Stack.new
    @operators_stack = Stack.new
  end

  def call(string)
    collect_input(string)
    operate while @operators_stack.length > 0
    @numbers_stack.pop
  end

  private

  def collect_input(string)
    items = string.split(',')
    items.each do |item|
      if item.match(NUMERIC)
        @numbers_stack.push(item)
      elsif item.match(SUM_OPERATOR)
        @operators_stack.push(item)
      else
        raise 'unexpected character'
      end
    end
    raise 'too many operators' if @operators_stack.length > @numbers_stack.length + 1
  end

  def operate
    number = @numbers_stack.pop.to_i
    other_number = @numbers_stack.pop.to_i
    operator = @operators_stack.pop

    result = send(operator, number, other_number)
    @numbers_stack.push(result)
  end

  def +(number, other_number)
    number + other_number
  end
end


c = Calculator.new
c.call('5,8,+')
c.call('9,4,+,6,+')
c.call('9,4,+,6,+,2,+,5,+')
