class Node
  attr_reader :data, :next_node
  attr_writer :data, :next_node

  def initialize(data, next_node)
    @data = data
    @next_node = next_node
  end
end

class LinkedList
  def initialize(value)
    @head = Node.new(value, nil)
    @size = 0
  end

  def add_first(value)
    raise 'empty list' if empty?

    new_node = if empty?
                 Node.new(value, nil)
               else
                 Node.new(value, @head)
               end
    @head = new_node
    @size += 1
  end

  def remove_first
    raise 'empty list' if empty?

    node = @head.data
    @head = @head.next_node
    @size -= 1

    node
  end

  def add_in_position(index, value)
    raise 'impossible to reach position' if index > @size + 1
    return add_first(value) if index.zero?
    return add_last(value) if index == @size + 1

    current_node = @head
    (index - 1).times do
      current_node = current_node.next_node
    end

    new_node = Node.new(value, current_node.next_node)
    current_node.next_node = new_node
    @size += 1
  end

  def remove_from_position(index)
    raise 'empty list' if empty?
    raise 'invalid position' if index > @size || index < 0
    return remove_first if index.zero?
    return remove_last if index == @size

    current_node = @head
    previous_node = nil
    index.times do
      previous_node = current_node
      current_node = current_node.next_node
    end

    previous_node.next_node = current_node.next_node
    @size -= 1
    current_node
  end

  def add_last(value)
    new_node = Node.new(value, nil)
    current_node = @head

    current_node = current_node.next_node while current_node.next_node
    current_node.next_node = new_node
    @size += 1
  end

  def remove_last
    raise 'empty list' if empty?
    return remove_first if @size.zero?

    current_node = @head
    previous_node = nil
    while current_node.next_node
      previous_node = current_node
      current_node = current_node.next_node
    end

    previous_node.next_node = nil
    @size -= 1
    current_node
  end

  def length
    @size + 1
  end

  def empty?
    @size == -1
  end

  ## debugging method
  def to_s
    raise 'empty list' if empty?
    elements = []
    current_node = @head
    while current_node.next_node
      elements << current_node.data
      current_node = current_node.next_node
    end
    elements << current_node.data

    p elements
  end
end
