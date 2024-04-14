package checker

import "github.com/antlr4-go/antlr/v4"

const (
	Integer = "Integer"
	Float   = "Float"
	String  = "String"
	Boolean = "Boolean"
	Char    = "Char"
)

type Identifier interface {
	GetToken() antlr.Token
	GetType() string
	GetValue() interface{}
	GetLevel() int
	IsConstant() bool
}

type VariableIdentifier struct {
	Token      antlr.Token
	Type       string
	Value      interface{}
	Level      int
	isConstant bool
}

func (v *VariableIdentifier) GetToken() antlr.Token {
	return v.Token
}

func (v *VariableIdentifier) GetType() string {
	return v.Type
}

func (v *VariableIdentifier) GetValue() interface{} {
	return v.Value
}

func (v *VariableIdentifier) GetLevel() int {
	return v.Level
}

func (v *VariableIdentifier) IsConstant() bool {
	return v.isConstant
}

type MethodIdentifier struct {
	Token  antlr.Token
	Type   string
	Value  interface{}
	Level  int
	Params []Param
}

type Param struct {
	Type string
	Name string
}

func (m *MethodIdentifier) GetToken() antlr.Token {
	return m.Token
}

func (m *MethodIdentifier) GetType() string {
	return m.Type
}

func (m *MethodIdentifier) GetValue() interface{} {
	return m.Value
}

func (m *MethodIdentifier) GetLevel() int {
	return m.Level
}

func (m *MethodIdentifier) IsConstant() bool {
	// Métodos no son constantes
	return false
}

type SymbolTable struct {
	Elements     []Identifier
	CurrentLevel int
}

//type SymbolTableVar struct {
//	Elements     []Identifier
//	CurrentLevel int
//}
//
//type VariableIdentifier struct {
//	Token      antlr.Token
//	Type       string
//	Value      interface{}
//	Level      int
//	isConstant bool
//}
//
//type MethodIdentifier struct {
//	Token  antlr.Token
//	Type   string
//	Value  interface{}
//	Level  int
//	Params []Param
//}
//
//type Param struct {
//	Type string
//	Name string
//}

//func (s *SymbolTable) addElement(token antlr.Token, t string, v interface{}) {
//	s.Elements = append([]Identifier{{Token: token, Type: t, Value: v}}, s.Elements...)
//}

func (s *SymbolTable) AddVariable(token antlr.Token, t string, v interface{}, isConstant bool) {
	identifier := &VariableIdentifier{
		Token:      token,
		Type:       t,
		Value:      v,
		Level:      s.CurrentLevel,
		isConstant: isConstant,
	}
	s.Elements = append(s.Elements, identifier)
}

func (s *SymbolTable) AddMethod(token antlr.Token, t string, v interface{}, params []Param) {
	identifier := &MethodIdentifier{
		Token:  token,
		Type:   t,
		Value:  v,
		Level:  s.CurrentLevel,
		Params: params,
	}
	s.Elements = append(s.Elements, identifier)
}

func (s *SymbolTable) search(token string) *Identifier {
	for _, element := range s.Elements {
		if element.GetToken().GetText() == token {
			return &element
		}
	}
	return nil
}

func (s *SymbolTable) searchOnCurrentScope(token string) *Identifier {
	for _, element := range s.Elements {
		if element.GetToken().GetText() == token && s.CurrentLevel == element.GetLevel() {
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
		if s.Elements[i].GetLevel() == s.CurrentLevel {
			s.Elements = append(s.Elements[:i], s.Elements[i+1:]...)
			i = i - 1
		}
	}
	s.CurrentLevel = s.CurrentLevel - 1
}

func (s *SymbolTable) printTable() {
	for _, element := range s.Elements {
		println(element.GetToken().GetText(), element.GetType(), element.GetValue(), element.GetValue())
	}
}
