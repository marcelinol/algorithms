class Calculator
  NUMERIC = /^\d+$/
  OPERATOR = /^(\+|\*)$/

  def initialize
    @main_stack = Stack.new
    @buffer_stack = Stack.new
  end

  def call(string)
    items = string.split(',')
    items.each do |item|
      raise 'unexpected character' unless item.match(NUMERIC) || item.match(OPERATOR)
      @main_stack.push(item)
    end

    @buffer_stack.push(@main_stack.pop) until @main_stack.empty?

    until @buffer_stack.length < 3
      number_one = @buffer_stack.pop
      number_two = @buffer_stack.pop
      operator = @buffer_stack.pop

      raise 'unexpected char' unless number_one.match(NUMERIC) && number_two.match(NUMERIC) && operator.match(OPERATOR)

      result = send(operator, number_one.to_i, number_two.to_i)
      puts "result of #{number_one} #{operator} #{number_two} = #{result}"
      @buffer_stack.push(result)
    end

    raise 'bad input' if @buffer_stack.length == 2

    @buffer_stack.pop
  end

  private

  def +(number, other_number)
    (number + other_number).to_s
  end

  def *(number_one, number_two)
    (number_one * number_two).to_s
  end
end

c = Calculator.new
# c.call('5,8,+')
# c.call('9,4,+,6,+')
# c.call('9,4,+,6,+,2,+,5,+')
c.call('2,4,+,6,+,2,*,2,+,5,+,100,*') # (((2+4+6) * 2)+2+5) * 100) = 3100
