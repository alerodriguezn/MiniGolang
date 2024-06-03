package llvm

import (
	"github.com/llir/llvm/ir/value"
)

type moduleStorage struct {
	currentScope int
	Elements     []moduleElement
}

func newModuleStorage() *moduleStorage {
	return &moduleStorage{
		currentScope: 0,
		Elements:     make([]moduleElement, 0),
	}
}

// Open Scope
func (m *moduleStorage) openScope() {
	m.currentScope++
}

func (m *moduleStorage) PutBack(name string, symbol value.Value) {
	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i].Name == name {
			m.Elements[i].value = symbol
		}
	}
}

// Close Scope and remove all elements in the current scope
func (m *moduleStorage) closeScope() {
	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i].Scope == m.currentScope {
			m.Elements = append(m.Elements[:i], m.Elements[i+1:]...)
			i = i - 1
		}
	}
	m.currentScope--
}

// Add an element to the storage
func (m *moduleStorage) addElement(name string, val value.Value) {
	m.Elements = append(m.Elements, moduleElement{
		Name:  name,
		Scope: m.currentScope,
		value: val,
	})
}

// Get an element from the storage using find first
func (m *moduleStorage) getElement(name string) (*moduleElement, bool) {
	for _, element := range m.Elements {
		if element.Name == name {
			return &element, true
		}
	}
	return nil, false
}

type moduleElement struct {
	Name  string
	Scope int
	value value.Value
}
