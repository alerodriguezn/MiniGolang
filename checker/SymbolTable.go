package checker

import "github.com/antlr4-go/antlr/v4"

const (
	Integer = "Integer"
	Float   = "Float"
	String  = "String"
	Boolean = "Boolean"
	Char    = "Char"
)

type SymbolTable struct {
	Elements     []Identifier
	CurrentLevel int
}

type Identifier struct {
	Token antlr.Token
	Type  string
	Value interface{}
	Level int
}

func (s *SymbolTable) addElement(token antlr.Token, t string, v interface{}) {
	s.Elements = append([]Identifier{{Token: token, Type: t, Value: v}}, s.Elements...)
}

func (s *SymbolTable) search(token string) *Identifier {
	for _, element := range s.Elements {
		if element.Token.GetText() == token {
			return &element
		}
	}
	return nil
}

func (s *SymbolTable) searchOnCurrentScope(token string) *Identifier {
	for _, element := range s.Elements {
		if element.Token.GetText() == token && s.CurrentLevel == element.Level {
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

func (s *SymbolTable) printTable() {
	for _, element := range s.Elements {
		println(element.Token.GetText(), element.Type, element.Value, element.Level)
	}
}
