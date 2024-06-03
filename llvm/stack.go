package llvm

import (
	"fmt"
	"github.com/llir/llvm/ir"
)

// Stack representa una pila genérica que contiene elementos de cualquier tipo
type Stack[T any] struct {
	elements []T
}

// NewStack crea y devuelve una nueva pila vacía
func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{elements: []T{}}
}

// Push agrega un elemento a la pila
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop elimina y devuelve el elemento superior de la pila
func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("pop from empty stack")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

// Peek devuelve el elemento superior de la pila sin eliminarlo
func (s *Stack[T]) Peek() (T, error) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("peek from empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

type Func struct {
	*ir.Func
	body *ir.Block
}
