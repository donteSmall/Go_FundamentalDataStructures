package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	n := l.First()
	assert.Equal(t, 1, n.value, "Incorrect Value")
	n = n.Next()
	assert.Equal(t, 2, n.value, "Incorrect Next Value")
	n = n.Next()
	assert.Equal(t, 3, n.value, "Incorrect Next Value")
}
func TestPrev(t *testing.T) {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	n := l.First()
	assert.Equal(t, 1, n.value, "Incorrect Value")
	next := n.Next()
	assert.Equal(t, 2, next.value, "Incorrect Next Value")
	next = next.Prev()
	assert.Equal(t, 1, next.value, "Incorrect Next Value")
}

func TestLast(t *testing.T) {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	n := l.Last()
	assert.Equal(t, 4, n.value, "Retrieved incorrect last value :")

}
