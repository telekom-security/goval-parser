package oval

import "encoding/xml"

// VariableTest : >tests>variable_test
type VariableTest struct {
	XMLName       xml.Name `xml:"variable_test"`
	ID            string   `xml:"id,attr"`
	StateOperator string   `xml:"state_operator,attr"`
	Comment       string   `xml:"comment,attr"`
	testRef
}

var _ Test = (*VariableTest)(nil)

// VariableObject : >objects>variable_object
type VariableObject struct {
	XMLName xml.Name `xml:"variable_object"`
	ID      string   `xml:"id,attr"`
	VarRef  string   `xml:"var_ref"`
}

// VariableState : >states>variable_state
type VariableState struct {
	XMLName xml.Name `xml:"variable_state"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value"`
}
