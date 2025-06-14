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

func TestDelete(t *testing.T) {
	testCases := []struct {
		name     string
		list     *linkedlist.LinkedList[int]
		input    int
		expected *linkedlist.LinkedList[int]
	}{
		{
			name:     "empty list",
			list:     linkedlist.New[int](),
			input:    2,
			expected: linkedlist.New[int](),
		},
		{
			name:     "delete head",
			list:     linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:    1,
			expected: linkedlist.NewFromSlice([]int{2, 3, 4}),
		},
		{
			name:     "delete tail",
			list:     linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:    4,
			expected: linkedlist.NewFromSlice([]int{1, 2, 3}),
		},
		{
			name:     "delete inner",
			list:     linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:    2,
			expected: linkedlist.NewFromSlice([]int{1, 3, 4}),
		},
		{
			name:     "delete multiple",
			list:     linkedlist.NewFromSlice([]int{1, 2, 1, 4, 1, 6, 8}),
			input:    1,
			expected: linkedlist.NewFromSlice([]int{2, 4, 6, 8}),
		},
		{
			name:     "delete all",
			list:     linkedlist.NewFromSlice([]int{1, 1, 1, 1}),
			input:    1,
			expected: linkedlist.New[int](),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := tc.list.Delete(tc.input)
			assertListsEqual(t, tc.expected, v)
		})
	}
}

func TestDeleteIndex(t *testing.T) {
	testCases := []struct {
		name           string
		list           *linkedlist.LinkedList[int]
		input          int
		expectedReturn bool
		expectedList   *linkedlist.LinkedList[int]
	}{
		{
			name:           "empty list",
			list:           linkedlist.New[int](),
			input:          2,
			expectedReturn: false,
			expectedList:   linkedlist.New[int](),
		},
		{
			name:           "invalid index",
			list:           linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:          25,
			expectedReturn: false,
			expectedList:   linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
		},
		{
			name:           "delete head",
			list:           linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:          0,
			expectedReturn: true,
			expectedList:   linkedlist.NewFromSlice([]int{2, 3, 4}),
		},
		{
			name:           "delete tail",
			list:           linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:          3,
			expectedReturn: true,
			expectedList:   linkedlist.NewFromSlice([]int{1, 2, 3}),
		},
		{
			name:           "delete inner",
			list:           linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:          2,
			expectedReturn: true,
			expectedList:   linkedlist.NewFromSlice([]int{1, 2, 4}),
		},
		{
			name:           "delete only node",
			list:           linkedlist.NewFromSlice([]int{5}),
			input:          0,
			expectedReturn: true,
			expectedList:   linkedlist.New[int](),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := tc.list.DeleteIndex(tc.input)
			if tc.expectedReturn != v {
				t.Errorf("expected %v but got %v", tc.expectedReturn, v)
			}
			assertListsEqual(t, tc.expectedList, tc.list)
		})
	}
}

func TestExists(t *testing.T) {
	l := linkedlist.NewFromSlice([]string{"a", "b", "c", "d", "d"})

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "exists",
			input: "b",
			want:  true,
		},
		{
			name:  "multiple",
			input: "d",
			want:  true,
		},
		{
			name:  "not exists",
			input: "z",
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := l.Exists(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name        string
		list        *linkedlist.LinkedList[int]
		input       int
		ok          bool
		resultCount int
		positions   []int
	}{
		{
			name:  "empty list",
			list:  linkedlist.New[int](),
			input: 2,
			ok:    false,
		},
		{
			name:        "one occurrence",
			list:        linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input:       3,
			ok:          true,
			resultCount: 1,
			positions:   []int{2},
		},
		{
			name:        "multiple occurrence",
			list:        linkedlist.NewFromSlice([]int{1, 2, 3, 4, 3, 3}),
			input:       3,
			ok:          true,
			resultCount: 3,
			positions:   []int{2, 4, 5},
		},
		{
			name:  "no occurrence",
			list:  linkedlist.NewFromSlice([]int{1, 2, 3, 4}),
			input: 99,
			ok:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, ok := tc.list.FindAll(tc.input)

			if tc.ok {
				if v == nil {
					t.Error("expected a return value but got none")
				}
				if len(v) != tc.resultCount {
					t.Errorf("expected result count: %v, got: %v", tc.resultCount, len(v))
				}

				for _, j := range tc.positions {
					_, err := tc.list.GetByIndex(j)
					if err != nil {
						t.Errorf("expected postion %v in result set, got none", j)
					}
				}
			} else {
				if ok {
					t.Error("expected ok return to be false but it is true")
				}
			}
		})
	}
}

func TestFindFirst(t *testing.T) {
	testCases := []struct {
		name             string
		list             *linkedlist.LinkedList[string]
		input            string
		ok               bool
		expectedPosition int
	}{
		{
			name:  "empty list",
			list:  linkedlist.New[string](),
			input: "a",
			ok:    false,
		},
		{
			name:             "found",
			list:             linkedlist.NewFromSlice([]string{"a", "b", "c", "d"}),
			input:            "c",
			ok:               true,
			expectedPosition: 2,
		},
		{
			name:             "finds first",
			list:             linkedlist.NewFromSlice([]string{"a", "b", "c", "d", "d"}),
			input:            "d",
			ok:               true,
			expectedPosition: 3,
		},
		{
			name:  "does not exist",
			list:  linkedlist.NewFromSlice([]string{"a", "b", "c", "d"}),
			input: "f",
			ok:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, ok := tc.list.FindFirst(tc.input)

			if tc.ok {
				if v == nil {
					t.Error("expected a NodeWithPosition value but got none")
				}
				if v.Position != tc.expectedPosition {
					t.Errorf("expected node to be at position: %v but it is at: %v", tc.expectedPosition, v.Position)
				}
			} else {
				if ok {
					t.Error("expected ok to be false but it is true")
				}
			}
		})
	}
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

func TestMiddle(t *testing.T) {
	testCases := []struct {
		name          string
		list          *linkedlist.LinkedList[int]
		expectError   bool
		expectedValue int
	}{
		{
			name:        "empty list",
			list:        linkedlist.New[int](),
			expectError: true,
		},
		{
			name:          "single node",
			list:          linkedlist.NewFromSlice([]int{1}),
			expectError:   false,
			expectedValue: 1,
		},
		{
			name:          "odd number nodes",
			list:          linkedlist.NewFromSlice([]int{1, 2, 3, 4, 5}),
			expectError:   false,
			expectedValue: 3,
		},
		{
			name:          "even number nodes",
			list:          linkedlist.NewFromSlice([]int{1, 2, 3, 4, 5, 6}),
			expectError:   false,
			expectedValue: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := tc.list.Middle()

			if tc.expectError && err == nil {
				t.Fatal("expected error but got none")
			}

			if err == nil && !tc.expectError {
				if r.Value() != tc.expectedValue {
					t.Errorf("expected %v got %v", tc.expectedValue, r.Value())
				}
			}
		})
	}
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

func assertListsEqual[T comparable](t *testing.T, a, b *linkedlist.LinkedList[T]) {
	t.Helper()

	if a.Length() != b.Length() {
		t.Error("lists do not have the same number of nodes")
	}

	if !reflect.DeepEqual(a.Values(), b.Values()) {
		t.Errorf("%v does not match %v", a.Values(), b.Values())
	}

	if a.Length() != 0 && b.Length() != 0 {
		if a.Head().Value() != b.Head().Value() {
			t.Error("head value mismatch")
		}

		if a.Tail().Value() != b.Tail().Value() {
			t.Error("tail value mismatch")
		}
	}
}
