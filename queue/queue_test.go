package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("NewQueue", func(t *testing.T) {
		q := NewQueue[int]()
		assert.NotNil(t, q)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("PushAndLen", func(t *testing.T) {
		q := NewQueue[int]()
		q.Push(1)
		assert.Equal(t, 1, q.Len())
		q.Push(2)
		assert.Equal(t, 2, q.Len())
	})

	t.Run("Pop", func(t *testing.T) {
		q := NewQueue[int]()
		q.Push(1)
		q.Push(2)

		val, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, val)
		assert.Equal(t, 1, q.Len())

		val, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, 2, val)
		assert.Equal(t, 0, q.Len())

		_, ok = q.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("Peek", func(t *testing.T) {
		q := NewQueue[int]()
		q.Push(1)
		q.Push(2)

		val, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, 1, val)
		assert.Equal(t, 2, q.Len())
	})

	t.Run("Clear", func(t *testing.T) {
		q := NewQueue[int]()
		q.Push(1)
		q.Push(2)
		q.Clear()

		assert.Equal(t, 0, q.Len())
		_, ok := q.Peek()
		assert.False(t, ok)
		_, ok = q.Pop()
		assert.False(t, ok)
	})

	t.Run("PopFromEmptyQueue", func(t *testing.T) {
		q := NewQueue[int]()
		val, ok := q.Pop()
		assert.False(t, ok)
		var zero int
		assert.Equal(t, zero, val)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("PeekFromEmptyQueue", func(t *testing.T) {
		q := NewQueue[int]()
		val, ok := q.Peek()
		assert.False(t, ok)
		var zero int
		assert.Equal(t, zero, val)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("PushAndPopWithPointers", func(t *testing.T) {
		q := NewQueue[*int]()
		a := 1
		b := 2
		q.Push(&a)
		q.Push(&b)

		val, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, &a, val)
		assert.Equal(t, 1, q.Len())

		val, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, &b, val)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("PushAndPopWithNils", func(t *testing.T) {
		q := NewQueue[*int]()
		q.Push(nil)
		q.Push(nil)

		val, ok := q.Pop()
		assert.True(t, ok)
		assert.Nil(t, val)
		assert.Equal(t, 1, q.Len())

		val, ok = q.Pop()
		assert.True(t, ok)
		assert.Nil(t, val)
		assert.Equal(t, 0, q.Len())

		val, ok = q.Pop()
		assert.False(t, ok)
		assert.Nil(t, val)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("PushAndPopWithStrings", func(t *testing.T) {
		q := NewQueue[string]()
		q.Push("hello")
		q.Push("world")

		val, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "hello", val)
		assert.Equal(t, 1, q.Len())

		val, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "world", val)
		assert.Equal(t, 0, q.Len())
	})

	t.Run("PushAndPopWithStructs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		q := NewQueue[Person]()
		q.Push(Person{Name: "A", Age: 30})
		q.Push(Person{Name: "B", Age: 25})

		val, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, Person{Name: "A", Age: 30}, val)
		assert.Equal(t, 1, q.Len())

		val, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, Person{Name: "B", Age: 25}, val)
		assert.Equal(t, 0, q.Len())
	})
}
