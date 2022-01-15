package games

type GameOutput struct {
	Task      ElementsSet
	Result    Element
	ResultSet ElementsSet
}

type IElement interface{}

type Element struct {
	Type  string
	Value string
}

type ElementsSet struct {
	Type     string
	Elements []IElement
}
