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
