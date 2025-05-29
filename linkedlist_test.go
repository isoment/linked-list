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

func TestNewFromSlice(t *testing.T) {
	t.Run("it creates a new list from a slice with the correct length, head and tail values", func(t *testing.T) {
		l := linkedlist.NewFromSlice([]int{1, 2, 3, 4})

		if l.Length() != 4 {
			t.Errorf("expected length 4 but got %v", l.Length())
		}

		if l.Head().Value() != 1 {
			t.Errorf("expected 1 to be the head but got %v", l.Head().Value())
		}

		if l.Tail().Value() != 4 {
			t.Errorf("expected 4 to be the head but got %v", l.Tail().Value())
		}
	})

	t.Run("it orders the elements correctly", func(t *testing.T) {
		l := linkedlist.NewFromSlice([]int{11, 22, 33, 44})

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

func TestInsert(t *testing.T) {
	t.Run("it returns an error if the index is less than 0", func(t *testing.T) {
		l := linkedlist.New[int]().
			Append(1).
			Append(2)

		_, err := l.Insert(-1, 99)
		if err == nil {
			t.Error("expected error got none")
		}
	})

	t.Run("it inserts to an empty list", func(t *testing.T) {
		l := linkedlist.New[int]()

		v, _ := l.Insert(0, 1)
		if v.Length() != 1 {
			t.Error("expected element to be added, list is empty")
		}
	})

	t.Run("value is inserted in the correct position", func(t *testing.T) {
		testCases := []struct {
			name  string
			index int
			value int
		}{
			{
				name:  "index 0",
				index: 0,
				value: 99,
			},
			{
				name:  "index 1",
				index: 1,
				value: 99,
			},
			{
				name:  "index 2",
				index: 2,
				value: 99,
			},
			{
				name:  "index 3",
				index: 3,
				value: 99,
			},
		}

		for _, tc := range testCases {
			l := linkedlist.New[int]().
				Append(1).
				Append(2).
				Append(3).
				Append(4)

			t.Run(tc.name, func(t *testing.T) {
				v, _ := l.GetByIndex(tc.index)
				if v.Value() == tc.value {
					t.Error("got unexpected insert value")
				}

				_, err := l.Insert(tc.index, tc.value)
				if err != nil {
					t.Error("got unexpected error")
				}

				w, _ := l.GetByIndex(tc.index)
				if w.Value() != tc.value {
					t.Errorf("expected %v, got %v", tc.value, w.Value())
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
