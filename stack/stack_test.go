package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		s := NewStack[int]()
		assert.NotNil(t, s)
		assert.True(t, s.Empty())
		assert.Equal(t, 0, s.Len())
	})

	t.Run("PushAndLen", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		assert.Equal(t, 1, s.Len())
		s.Push(2)
		assert.Equal(t, 2, s.Len())
	})

	t.Run("Pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)

		val, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 2, val)
		assert.Equal(t, 1, s.Len())

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, val)
		assert.Equal(t, 0, s.Len())

		_, ok = s.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, s.Len())
	})

	t.Run("Peek", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)

		val, ok := s.Peek()
		assert.True(t, ok)
		assert.Equal(t, 2, val)
		assert.Equal(t, 2, s.Len())
	})

	t.Run("Empty", func(t *testing.T) {
		s := NewStack[int]()
		assert.True(t, s.Empty())

		s.Push(1)
		assert.False(t, s.Empty())

		s.Pop()
		assert.True(t, s.Empty())
	})

	t.Run("Len", func(t *testing.T) {
		s := NewStack[int]()
		assert.Equal(t, 0, s.Len())

		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Len())

		s.Pop()
		assert.Equal(t, 1, s.Len())
	})

	t.Run("IndexOf", func(t *testing.T) {
		s := NewStack[int]()
		assert.Equal(t, -1, s.IndexOf(1))

		s.Push(1)
		s.Push(2)
		assert.Equal(t, 0, s.IndexOf(2))
		assert.Equal(t, 1, s.IndexOf(1))
		assert.Equal(t, -1, s.IndexOf(3))
	})

	t.Run("PopFromEmptyStack", func(t *testing.T) {
		s := NewStack[int]()
		val, ok := s.Pop()
		assert.False(t, ok)
		var zero int
		assert.Equal(t, zero, val)
		assert.Equal(t, 0, s.Len())
	})

	t.Run("PeekFromEmptyStack", func(t *testing.T) {
		s := NewStack[int]()
		val, ok := s.Peek()
		assert.False(t, ok)
		var zero int
		assert.Equal(t, zero, val)
		assert.Equal(t, 0, s.Len())
	})

	t.Run("PushAndPopWithPointers", func(t *testing.T) {
		s := NewStack[*int]()
		a := 1
		b := 2
		s.Push(&a)
		s.Push(&b)

		val, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, &b, val)
		assert.Equal(t, 1, s.Len())

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, &a, val)
		assert.Equal(t, 0, s.Len())
	})

	t.Run("PushAndPopWithNils", func(t *testing.T) {
		s := NewStack[*int]()
		s.Push(nil)
		s.Push(nil)

		val, ok := s.Pop()
		assert.True(t, ok)
		assert.Nil(t, val)
		assert.Equal(t, 1, s.Len())

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Nil(t, val)
		assert.Equal(t, 0, s.Len())

		val, ok = s.Pop()
		assert.False(t, ok)
		assert.Nil(t, val)
		assert.Equal(t, 0, s.Len())
	})

	t.Run("PushAndPopWithStrings", func(t *testing.T) {
		s := NewStack[string]()
		s.Push("hello")
		s.Push("world")

		val, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, "world", val)
		assert.Equal(t, 1, s.Len())

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, "hello", val)
		assert.Equal(t, 0, s.Len())
	})

	t.Run("PushAndPopWithStructs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		s := NewStack[Person]()
		s.Push(Person{Name: "A", Age: 30})
		s.Push(Person{Name: "B", Age: 25})

		val, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, Person{Name: "B", Age: 25}, val)
		assert.Equal(t, 1, s.Len())

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, Person{Name: "A", Age: 30}, val)
		assert.Equal(t, 0, s.Len())
	})
}
