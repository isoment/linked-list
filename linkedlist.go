package linkedlist

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func (n *Node[T]) Next() T {
	return n.value
}

func (n *Node[T]) Value() T {
	return n.value
}

type LinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

type NodeWithPosition[T comparable] struct {
	Node     *Node[T]
	Position int
}

func NewFromSlice[T comparable](input []T) *LinkedList[T] {
	new := New[T]()

	if len(input) == 0 {
		return new
	}

	for _, v := range input {
		new.Append(v)
	}

	return new
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add a node to the end of the list
func (l *LinkedList[T]) Append(value T) *LinkedList[T] {
	new := &Node[T]{value: value}
	l.length++

	if l.head == nil {
		l.head = new
		l.tail = new
		return l
	}

	l.tail.next = new
	l.tail = new

	return l
}

// Delete all occurrences of a value in the list
func (l *LinkedList[T]) Delete(value T) *LinkedList[T] {
	if l.length == 0 {
		return l
	}

	current := l.head
	var prev *Node[T] = nil

	for current != nil {
		if current.value == value {
			l.length--
			// Need to handle the cases when we delete a head node, a tail node, or a middle node
			if current == l.head {
				l.head = current.next
				if current == l.tail {
					l.tail = nil
				}
				current = l.head
			} else if current == l.tail {
				l.tail = prev
				prev.next = nil
				current = nil
			} else {
				prev.next = current.next
				current = current.next
			}
		} else {
			prev = current
			current = current.next
		}
	}

	return l
}

// Check if a node with the given value exists in the list
func (l *LinkedList[T]) Exists(value T) bool {
	current := l.head

	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}

	return false
}

// Get the Node at the given 0 based index
func (l *LinkedList[T]) GetByIndex(index int) (*Node[T], error) {
	if index < 0 {
		return nil, ErrorInvalidIndex
	}

	if l.length == 0 {
		return nil, ErrorEmptyList
	}

	i := 0
	current := l.head

	for current != nil {
		if i == index {
			return current, nil
		} else {
			i++
			current = current.next
		}
	}

	return nil, ErrorInvalidIndex
}

// Get the first node in the list
func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

// Insert a value at the given 0 based index in the list
func (l *LinkedList[T]) Insert(index int, value T) (*LinkedList[T], error) {
	if index < 0 {
		return nil, ErrorInvalidIndex
	}

	// If the list is empty or the index is outside the list size append
	if l.length == 0 || index >= l.length {
		l.Append(value)
		return l, nil
	}

	// Handle the case where we are inserting at the beginning
	if index == 0 {
		l.Prepend(value)
		return l, nil
	}

	i := 0
	current := l.head

	for current != nil {
		if i+1 == index {
			new := &Node[T]{value: value}

			next := current.next
			current.next = new
			new.next = next
			break
		} else {
			i++
			current = current.next
		}
	}

	return l, nil
}

// The length of the list
func (l *LinkedList[T]) Length() int {
	return l.length
}

// Add a node to the beginning of the list
func (l *LinkedList[T]) Prepend(value T) *LinkedList[T] {
	new := &Node[T]{value: value}
	l.length++

	if l.head == nil {
		l.head = new
		l.tail = new
		return l
	}

	new.next = l.head
	l.head = new

	return l
}

// Find all occurrences of a value in the list
func (l *LinkedList[T]) FindAll(value T) ([]NodeWithPosition[T], bool) {
	var result []NodeWithPosition[T]

	i := 0
	current := l.head

	for current != nil {
		if current.value == value {
			n := NodeWithPosition[T]{
				Node:     current,
				Position: i,
			}
			result = append(result, n)
		}
		i++
		current = current.next
	}

	if len(result) > 0 {
		return result, true
	} else {
		return nil, false
	}
}

// Find the first occurrence of a value in the list and its 0 based index
func (l *LinkedList[T]) FindFirst(value T) (result *NodeWithPosition[T], ok bool) {
	i := 0
	current := l.head

	for current != nil {
		if current.value == value {
			return &NodeWithPosition[T]{
				Node:     current,
				Position: i,
			}, true
		} else {
			i++
			current = current.next
		}
	}

	return nil, false
}

// Get the last node in the list
func (l *LinkedList[T]) Tail() *Node[T] {
	return l.tail
}

// Get the values from the list as a slice
func (l *LinkedList[T]) Values() []T {
	var values []T

	current := l.head

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}
