class Calculator
  NUMERIC = /^\d+$/
  SUM_OPERATOR = /^\+$/

  def initialize
    @main_stack = Stack.new
    @buffer_stack = Stack.new(1)
  end

  def call(string)
    items = string.split(',')
    items.each do |item|
      raise 'unexpected character' unless item.match(NUMERIC) || item.match(SUM_OPERATOR)
      @main_stack.push(item)
    end

    calculate(@main_stack.pop, @main_stack.pop, @main_stack.pop)
    @main_stack.pop
  end

  private

  def calculate(operator, first_param, second_param)
    raise "unexpected formation. Expected a '+'." unless operator.match(SUM_OPERATOR)
    raise "unexpected formation: two '+' in sequence" unless first_param.match(NUMERIC)

    if second_param.match(SUM_OPERATOR)
      begin
        @buffer_stack.push(second_param)
      rescue StackOverflow
        raise "unexpected formation: two '+' in sequence"
      end
        calculate(operator, first_param, @main_stack.pop)
    else
      result = send(operator, first_param.to_i, second_param.to_i)
      @main_stack.push(result)
      @main_stack.push(@buffer_stack.pop) if @buffer_stack.length > 0

      raise 'unexpected formation' if @main_stack.length === 2
      calculate(@main_stack.pop, @main_stack.pop, @main_stack.pop) if @main_stack.length > 2
    end
  end

  def +(number, other_number)
    (number + other_number).to_s
  end
end

c = Calculator.new
# c.call('5,8,+')
# c.call('9,4,+,6,+')
c.call('9,4,+,6,+,2,+,5,+')
