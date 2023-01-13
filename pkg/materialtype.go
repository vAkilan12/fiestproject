package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/hbollon/go-edlib"
)

func GetMaterialType(descriptionArray []string, attributeList []string, name string) {
	// Using for Loop
	for descriptionKey := 1; descriptionKey < len(descriptionArray); descriptionKey++ {
		// Spliting the description after "Comma"
		splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
		for _, str := range splitAftercomma {
			//Spliting the description after "Space"
			splitAfterspace := strings.Fields(str)
			for _, str1 := range splitAfterspace {
				// Converting description to uppercase
				convToupper := strings.ToUpper(str1)
				if edlib.JaroSimilarity(convToupper, name) >= 0.85 {
					splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
					for _, str2 := range splitAftercomma {
						splitAfterspace := strings.Fields(str2)
						for _, str3 := range splitAfterspace {
							convToupper := strings.ToUpper(strings.ReplaceAll(str3, " ", ""))
							for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
								if edlib.JaroSimilarity(convToupper, attributeList[attributeListkey]) > 0.76 {
									// to find similarity between description and AttributeList and returing the similar material type in AttributeList
									res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.3, edlib.Levenshtein)
									if err != nil {
										log.Println(err)
									}
									fmt.Println("Material Type:", convToupper, res[0])
								}

							}

						}
					}

				}
			}
		}
	}
}

func FindMaterial(descriptionArray, attributeList []string, descriptionKey int) string {
	var matchedMaterial string
	splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
	for _, str := range splitAftercomma {
		splitAfterspace := strings.Fields(str)
		for _, str1 := range splitAfterspace {
			convToupper := strings.ToUpper(strings.ReplaceAll(str1, " ", ""))
			for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
				if edlib.JaroSimilarity(convToupper, attributeList[attributeListkey]) > 0.76 {
					// to find similarity between description and AttributeList and returing the similar material type in AttributeList
					res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.3, edlib.Levenshtein)
					if err != nil {
						log.Println(err)
					}
					matchedMaterial = res[0]

					fmt.Println("Material Type:", convToupper, res[0])
				}

			}

		}
	}
	return matchedMaterial
}
