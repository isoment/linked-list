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

// Get the first node in the list
func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

// The length of the list
func (l *LinkedList[T]) Length() int {
	return l.length
}

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

func (l *LinkedList[T]) Values() []T {
	var values []T

	current := l.head

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}
