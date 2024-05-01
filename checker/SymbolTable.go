package checker

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type idType int

const (
	varId idType = 1 << iota
	typeId
	funId
	sliceId
	arrId
	strId
)

type Identifier struct {
	Token      antlr.Token
	Type       idType
	LitType    string
	Value      interface{}
	Level      int
	isConstant bool
	isSlice    bool
	isStruct   bool
	isArray    bool
	list       []*Identifier
}

type SymbolTable struct {
	Elements       []Identifier
	AvailableTypes []string
	CurrentLevel   int
}

func (s *SymbolTable) AddType(_type string) {
	//Check if the type is already defined
	for _, t := range s.AvailableTypes {
		if t == _type {
			return
		}
	}

	s.AvailableTypes = append(s.AvailableTypes, _type)
}

func (s *SymbolTable) AddPredefinedType() {
	s.AvailableTypes = append(s.AvailableTypes, "int")
	s.AvailableTypes = append(s.AvailableTypes, "float")
	s.AvailableTypes = append(s.AvailableTypes, "string")
	s.AvailableTypes = append(s.AvailableTypes, "bool")
	s.AvailableTypes = append(s.AvailableTypes, "char")
}

func (s *SymbolTable) isValidType(t string) bool {
	for _, _type := range s.AvailableTypes {
		if t == _type {
			return true
		}
	}
	return false

}

func (s *SymbolTable) Print() {

	fmt.Println("================ Symbol Table: ================")
	for i, element := range s.Elements {
		fmt.Printf("Element %d:\n", i+1)
		fmt.Printf("  Token: %v\n", element.Token)
		fmt.Printf("  Type: %d\n", element.Type)
		fmt.Printf("  Value: %v\n", element.Value)
		fmt.Printf("  Level: %d\n", element.Level)
		fmt.Printf("  LitType: %s\n", element.LitType)
		if element.isArray {
			fmt.Printf("  isArray: %v\n", element.isArray)
		}
		if element.isSlice {
			fmt.Printf("  isSlice: %v\n", element.isSlice)
		}
		if element.isStruct {
			fmt.Printf("  isStruct: %v\n", element.isStruct)
		}

		if len(element.list) > 0 {
			fmt.Println("  List of identifiers:")
			for j, subElement := range element.list {
				fmt.Printf("    Subelement %d:\n", j+1)
				fmt.Printf("      Token: %v\n", subElement.Token)
				fmt.Printf("      Type: %d\n", subElement.Type)
				fmt.Printf("      Value: %v\n", subElement.Value)
				fmt.Printf("      Level: %d\n", subElement.Level)
				fmt.Printf("      LitType: %s\n", subElement.LitType)

				// Y así sucesivamente para cada campo que desees imprimir
			}
		}
		fmt.Println() // Espacio entre elementos para mejor legibilidad
	}
}

func (s *SymbolTable) getCurrentFunction() *Identifier {
	for _, element := range s.Elements {
		if element.Type == funId {
			return &element
		}
	}
	return nil
}

func (s *SymbolTable) NewVar(token antlr.Token, value string, typ string) *Identifier {
	return &Identifier{
		Token:   token,
		Type:    varId,
		Value:   value,
		LitType: typ,
		Level:   s.CurrentLevel,
		list:    []*Identifier{},
	}
}

func (s *SymbolTable) NewType(token antlr.Token, value string, typ string, isSlice bool) *Identifier {
	return &Identifier{
		Token:   token,
		Type:    typeId,
		Value:   value,
		Level:   s.CurrentLevel,
		LitType: typ,
		list:    []*Identifier{},
		isSlice: isSlice,
	}
}

func (s *SymbolTable) NewAliasType(token antlr.Token, value string, typ string) *Identifier {
	return &Identifier{
		Token:   token,
		Type:    typeId,
		Value:   value,
		Level:   s.CurrentLevel,
		LitType: typ,
		list:    []*Identifier{},
	}
}

func (s *SymbolTable) NewFun(token antlr.Token, value string, typ string, list []*Identifier) *Identifier {
	return &Identifier{
		Token:   token,
		Type:    funId,
		Value:   value,
		Level:   s.CurrentLevel,
		LitType: typ,
		list:    list,
	}
}

func (s *SymbolTable) NewStruct(token antlr.Token, value string, optionalLists ...*Identifier) *Identifier {

	var list []*Identifier
	if len(optionalLists) > 0 {
		list = optionalLists
	} else {
		list = nil
	}

	return &Identifier{
		Token:    token,
		Type:     typeId,
		Value:    value,
		Level:    s.CurrentLevel,
		LitType:  "",
		list:     list,
		isStruct: true,
	}

}

func (s *SymbolTable) AddElement(element *Identifier) error {

	for _, e := range s.Elements {
		if e.Value == element.Value {
			if e.Type == funId && e.Level == element.Level {
				return fmt.Errorf("function '%s' already defined in the current scope", e.Value)
			}
			if e.Type == varId && e.Level == element.Level {
				return fmt.Errorf("variable '%s' already defined in the current scope", e.Value)
			}
			if e.Type == typeId && e.Level == element.Level {
				return fmt.Errorf("type '%s' already defined in the current scope", e.Value)
			}
		}
	}
	// Add the element to the list (add first)
	s.Elements = append([]Identifier{*element}, s.Elements...)
	s.Print()
	return nil

}

func (s *SymbolTable) search(token string) *Identifier {
	// search using find first
	for _, element := range s.Elements {
		if element.Value == token {
			return &element
		}
	}
	return nil
}

func (s *SymbolTable) openScope() {
	s.CurrentLevel = s.CurrentLevel + 1
}

func (s *SymbolTable) closeScope() {
	//Delete all elements in the current scope
	for i := 0; i < len(s.Elements); i++ {
		if s.Elements[i].Level == s.CurrentLevel {
			s.Elements = append(s.Elements[:i], s.Elements[i+1:]...)
			i = i - 1
		}
	}
	s.CurrentLevel = s.CurrentLevel - 1
}

func NewSymbolTable() *SymbolTable {
	symbolTable := &SymbolTable{
		Elements:       []Identifier{},
		AvailableTypes: []string{},
		CurrentLevel:   0,
	}
	symbolTable.AddPredefinedType()
	return symbolTable
}
