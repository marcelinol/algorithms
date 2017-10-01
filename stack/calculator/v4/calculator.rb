class Calculator
  NUMERIC = /^\d+$/
  SUM_OPERATOR = /^\+$/

  def initialize
    @operators = Stack.new
    @numbers = Stack.new
  end

  def call(string)
    load_operators(string)

    until @operators.length.zero?
      op = @operators.pop
      number_one = @numbers.pop
      number_two = @numbers.pop

      @numbers.push(send(op, number_one, number_two))
    end

    @numbers.pop
  end

  def load_operators(string)
    items = string.split(',')
    items.each do |item|
      @operators.push(item) if item.match(SUM_OPERATOR)
      @numbers.push(item) if item.match(NUMERIC)
    end
  end

  def +(number_one, number_two)
    (number_one.to_i + number_two.to_i).to_s
  end
end

c = Calculator.new
# c.call('5,8,+')
# c.call('9,4,+,6,+')
c.call('9,4,+,6,+,2,+,5,+')
