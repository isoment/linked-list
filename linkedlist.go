package linkedlist

type Node[T any] struct {
	value T
	next  *Node[T]
}

func (n *Node[T]) Next() T {
	return n.value
}

func (n *Node[T]) Value() T {
	return n.value
}

type LinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func New[T any]() *LinkedList[T] {
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

// Get the Node at the given 0 based index
func (l *LinkedList[T]) GetByIndex(index int) (*Node[T], error) {
	if index < 0 {
		return nil, ErrorInvalidIndex
	}

	if l.length == 0 {
		return nil, ErrorEmptyList
	}

	count := 0
	current := l.head

	for current != nil {
		if count == index {
			return current, nil
		} else {
			count++
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

	count := 0
	current := l.head

	for current != nil {
		if count+1 == index {
			new := &Node[T]{value: value}

			next := current.next
			current.next = new
			new.next = next
			break
		} else {
			count++
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
