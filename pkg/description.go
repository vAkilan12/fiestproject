package pkg

type Attribute struct {
	Name   string `json:"attribute_name"`
	Option string `json:"option_name"`
}

type DecodedDescription struct {
	Description string      `json:"description"`
	NounName    string      `json:"noun_name"`
	Attributes  []Attribute `json:"attributes"`
}

type Noun struct {
	Name       string   `json:"noun_name"`
	Attributes []string `json:"attributes"`
}

type Nouns struct {
	Nouns []Noun `json:"nouns"`
}
