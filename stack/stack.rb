class StackOverflow < StandardError
  def initialize(msg = 'stack overflow')
    super(msg)
  end
end

class Stack
  def initialize(size = 30)
    @size = size
    @top = -1
    @data = []
  end

  def push(item)
    raise StackOverflow if full?
    @top += 1
    @data[@top] = item
  end

  def pop
    raise 'empty stack' if empty?
    @top -= 1
    @data[@top + 1]
  end

  def peek
    raise 'empty stack' if empty?
    @data[@top]
  end

  def clear
    @top = -1
  end

  def length
    @top + 1
  end

  def to_s
    @data.take(@top + 1)
  end

  def full?
    @top == (@size - 1)
  end

  def empty?
    @top == -1
  end
end
