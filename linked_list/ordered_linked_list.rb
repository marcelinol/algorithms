class Node
  attr_reader :data, :next_node
  attr_writer :data, :next_node

  def initialize(data, next_node)
    @data = data
    @next_node = next_node
  end
end

class OrderedLinkedList
  def initialize(value)
    @head = Node.new(value, nil)
    @size = 0
  end

  def add(value)
    raise 'empty list' if empty?
    if empty?
      new_node = Node.new(value, nil)
      @head = new_node
      @size += 1
    elsif value <= @head.data
      @head = Node.new(value, @head)
      @size += 1
    else
      add_in_order(value)
    end
  end

  def add_in_order(value)
    current_node = @head.next_node
    previous_node = @head

    while current_node && current_node.data < value
      previous_node = current_node
      current_node = current_node.next_node
    end

    new_node = Node.new(value, current_node)
    previous_node.next_node = new_node
    @size += 1
  end

  def empty?
    @size == -1
  end

  def remove_first
    raise 'empty list' if empty?

    node = @head.data
    @head = @head.next_node
    @size -= 1

    node
  end

  def remove_from_position(index)
    raise 'empty list' if empty?
    raise 'invalid position' if index > @size || index < 0
    return remove_first if index.zero?

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

l = OrderedLinkedList.new(1)
