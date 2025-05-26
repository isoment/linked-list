package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/isoment/linked-list"
)

func TestNew(t *testing.T) {
	t.Run("it creates a new empty list", func(t *testing.T) {
		l := linkedlist.New[int]()
		if l.Length() != 0 {
			t.Error("expected an empty list")
		}
	})
}

func TestAppend(t *testing.T) {
	t.Run("it appends a new value to the list", func(t *testing.T) {
		l := linkedlist.New[int]()
		l.Append(1)
		l.Append(2)
		if l.Length() != 2 {
			t.Error("expected a list of 2 items")
		}
	})

	t.Run("it has the elements in correct order", func(t *testing.T) {
		l := linkedlist.New[int]()
		l.Append(1)
		l.Append(2)
		l.Append(3)
		l.Append(50)

		got := l.Values()
		want := []int{1, 2, 3, 50}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, got %v", want, got)
		}

		if l.Head().Value() != 1 || l.Tail().Value() != 50 {
			t.Errorf("error asserting head and tail values")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("it returns an error if the list is empty", func(t *testing.T) {
		l := linkedlist.New[int]()
		_, err := l.GetByIndex(6)
		if err == nil {
			t.Error("expected error got none")
		}
	})

	t.Run("it returns an error if the index param is invalid", func(t *testing.T) {
		l := linkedlist.New[int]().Append(1)
		_, err := l.GetByIndex(-1)
		if err == nil {
			t.Error("expected error got none")
		}
	})

	t.Run("it returns an error if the index is out of bounds", func(t *testing.T) {
		l := linkedlist.New[int]().
			Append(1).
			Append(2)

		_, err := l.GetByIndex(6)
		if err == nil {
			t.Error("expected error got none")
		}
	})

	t.Run("it returns the node at nth index", func(t *testing.T) {
		l := linkedlist.New[int]().
			Append(11).
			Append(22).
			Append(33).
			Append(44)

		testCases := []struct {
			name  string
			index int
			value int
		}{
			{
				name:  "first element",
				index: 0,
				value: 11,
			},
			{
				name:  "second element",
				index: 1,
				value: 22,
			},
			{
				name:  "third element",
				index: 2,
				value: 33,
			},
			{
				name:  "fourth element",
				index: 3,
				value: 44,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				v, _ := l.GetByIndex(tc.index)
				if v.Value() != tc.value {
					t.Errorf("expected %v, got %v", tc.value, v.Value())
				}
			})
		}
	})
}

func TestPrepend(t *testing.T) {
	t.Run("it prepends a value to the list", func(t *testing.T) {
		l := linkedlist.New[int]()
		l.Append(1)
		l.Append(2)
		l.Prepend(23)

		if l.Length() != 3 {
			t.Error("expected a list of 3 items")
		}
	})
	t.Run("it has the elements in correct order", func(t *testing.T) {
		l := linkedlist.New[int]()
		l.Append(1)
		l.Append(2)
		l.Append(3)
		l.Prepend(50)

		got := l.Values()
		want := []int{50, 1, 2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, got %v", want, got)
		}

		if l.Head().Value() != 50 || l.Tail().Value() != 3 {
			t.Errorf("error asserting head and tail values")
		}
	})
}
