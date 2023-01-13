package pkg

func Decode(descriptions []string, nounsToIdentiy Nouns) []DecodedDescription {

	var decodedDescs []DecodedDescription

	for _, description := range descriptions {
		decoded := false
		for _, noun := range nounsToIdentiy.Nouns {
			decodedDesc := findNoun(description, noun)
			if decodedDesc != nil {
				decodedDescs = append(decodedDescs, *decodedDesc)
				decoded = true
				break
			}
		}
		if !decoded {
			decodedDescs = append(decodedDescs, DecodedDescription{
				Description: description,
			})
		}
	}
	return decodedDescs
}

func findNoun(description string, noun Noun) *DecodedDescription {

	var decodedDesc DecodedDescription

	decodedDesc = DecodedDescription{
		Description: description,
		NounName:    noun.Name,
	}

	for _, attribute_name := range noun.Attributes {
		attribute := findAttribute(description, attribute_name)
		decodedDesc.Attributes = append(decodedDesc.Attributes, attribute)
	}

	return &decodedDesc
}

func findAttribute(description, attributeName string) Attribute {
	var optionName string
	switch attributeName {
	case "MaterialType":
		optionName = FindMaterialType(description)
	default:
		optionName = ""
	}
	return Attribute{
		Name:   attributeName,
		Option: optionName,
	}
}
