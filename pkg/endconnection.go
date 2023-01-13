package pkg

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/hbollon/go-edlib"
)

func GetEndConnection(descriptionArray []string, attributeList []string, name string) {
	// Using for Loop
	for descriptionKey := 1; descriptionKey < len(descriptionArray); descriptionKey++ {
		// Spliting the description after "Comma"
		splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
		for _, str := range splitAftercomma {
			//Spliting the description after "Space"
			splitAfterspace := strings.Fields(str)
			for _, str1 := range splitAfterspace {
				// removing all the spaces
				// Converting description to uppercase
				convToupper := strings.ToUpper(strings.ReplaceAll(str1, " ", ""))
				if edlib.JaroSimilarity(convToupper, name) >= 0.85 {
					for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {

						// SXT
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w+)(X)(?P<size2>\w+)`, `(?m)(?P<size1>\w+)(X)`,
							`(?m)(X)(?P<size2>\w+)`, 0.3, 1)

						//1-1/2INCHSOCXTHD
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\w+)`, `(?m)(?P<size1>\d\W\d\W\w+)(X)`,
							`(?m)(X)(?P<size2>\w+)`, 0.1, 0.1)

						//3/4INSOCX3/4INSOC
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\w+)(X)(?P<size2>\d\W\d\w+)`, `(?m)(?P<size1>\d\W\d\w+)(X)`,
							`(?m)(X)(?P<size2>\d\W\d\w+)`, 0.1, 0.1)

						////M20MALEXM25FEMALE
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w\d+\w+)(X)(?P<size2>\w\d+\w+)`, `(?m)(?P<size1>\w\d+\w+)(X)`,
							`(?m)(X)(?P<size2>\w\d+\w+)`, 0.1, 0.1)

						//1.5"HoseXTHD
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\w+)`, `(?m)(?P<size1>\d\W\d\W\w+)(X)`,
							`(?m)(X)(?P<size2>\w+)`, 0.3, 0.1)

						//3/8"MTX3/8"TB
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\d\W\d\W\w+)`, `(?m)(?P<size1>\d\W\d\W\w+)(X)`,
							`(X)(?P<size2>\d\W\d\W\w+)`, 0.1, 0.1)

						//M16X1/2"
						// Extracting the above End type by using the regexp expression
						FindEndAtrributes(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w\d+)(X)(?P<size2>\d\W\d\W)`, `(?m)(?P<size1>\w\d+)(X)`,
							`(X)(?P<size2>\d\W\d\W)`, 0.1, 0.1)

						//MALE
						// Extracting the above End type by using the regexp expression
						if edlib.JaroSimilarity(convToupper, attributeList[attributeListkey]) > 0.8 {
							// to find similarity between description andattributeList and returing the similar end type inattributeList
							res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.4, edlib.Levenshtein)
							if err != nil {
								log.Println(err)
							}
							fmt.Println("End Conn2:", convToupper, res[0])
						}
					}
				}
			}

		}
	}
}

func FindEndAtrributes(descriptionArray, attributeList []string, descriptionKey, attributeListkey int,
	regExp, regExp1, regExp2 string, searchThreshold1, searchThreshold2 float32) (string, string) {
	var matchedEnd1, matchedEnd2 string
	var re = regexp.MustCompile(regExp)
	convToupper := strings.ToUpper(strings.ReplaceAll(descriptionArray[descriptionKey], " ", ""))
	for _, match := range re.FindAllString(convToupper, -1) {
		var re1 = regexp.MustCompile(regExp1)
		var re2 = regexp.MustCompile(regExp2)
		for _, match1 := range re1.FindAllString(match, -1) {
			if edlib.JaroSimilarity(match1, attributeList[attributeListkey]) > 0.4 {
				// to find similarity between description andattributeList and returing the similar Size inattributeList
				res1, err := edlib.FuzzySearchSetThreshold(match1, attributeList, 1, searchThreshold1, edlib.Levenshtein)
				if err != nil {
					log.Println(err)
				}
				matchedEnd1 = res1[0]
				fmt.Println("End Conn1:", match, res1[0])
			}
			for _, match2 := range re2.FindAllString(match, -1) {
				if edlib.JaroSimilarity(match2, attributeList[attributeListkey]) > 0.4 {
					Aftertrim := strings.Trim(match2, "X")
					res2, err := edlib.FuzzySearchSetThreshold(Aftertrim, attributeList, 1, searchThreshold2, edlib.Levenshtein)
					if err != nil {
						log.Println(err)
					}
					matchedEnd2 = res2[0]
					fmt.Println("End Conn2:", match, res2[0])

				}
			}
		}
	}
	return matchedEnd1, matchedEnd2
}

var endExp = []string{
	`(?m)(?P<size1>\w+)(X)(?P<size2>\w+)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\w+)`,
	`(?m)(?P<size1>\d\W\d\w+)(X)(?P<size2>\d\W\d\w+)`,
	`(?m)(?P<size1>\w\d+\w+)(X)(?P<size2>\w\d+\w+)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\w+)`,
	`(?m)(?P<size1>\w\d+)(X)(?P<size2>\d\W\d\W)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\d\W\d\W\w+)`,
}

var endExp1 = []string{
	`(?m)(?P<size1>\w+)(X)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)`,
	`(?m)(?P<size1>\d\W\d\w+)(X)`,
	`(?m)(?P<size1>\w\d+\w+)(X)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)`,
	`(?m)(?P<size1>\w\d+)(X)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)`,
}

var endExp2 = []string{
	`(?m)(X)(?P<size2>\w+)`,
	`(?m)(X)(?P<size2>\w+)`,
	`(?m)(X)(?P<size2>\d\W\d\w+)`,
	`(?m)(X)(?P<size2>\w\d+\w+)`,
	`(?m)(X)(?P<size2>\w+)`,
	`(X)(?P<size2>\d\W\d\W)`,
	`(X)(?P<size2>\d\W\d\W\w+)`,
}

var endSimilarity1 = []float32{
	0.3, 0.1, 0.1, 0.1, 0.3, 0.1, 0.1, 0.1,
}

var endSimilarity2 = []float32{
	0.5, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1,
}

func FindEnd(descriptionArray, attributeList []string, descriptionKey int) (string, string) {
	var matchedEnd1, matchedEnd2 string
	splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
	for _, str := range splitAftercomma {
		splitAfterspace := strings.Fields(str)
		for _, str1 := range splitAfterspace {
			convToupper := strings.ToUpper(strings.ReplaceAll(str1, " ", ""))
			for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
				for endExplistKey := 0; endExplistKey < len(endExp); endExplistKey++ {
					string1, string2 := FindEndAtrributes(descriptionArray, attributeList, descriptionKey,
						attributeListkey, endExp[endExplistKey], endExp1[endExplistKey], endExp2[endExplistKey],
						endSimilarity1[endExplistKey], endSimilarity2[endExplistKey])
					if len(string1) > 0 {
						matchedEnd1 = string1
						matchedEnd2 = string2
					}
				}
				for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
					//MALE
					// Extracting the above End type by using the regexp expression
					if edlib.JaroSimilarity(convToupper, attributeList[attributeListkey]) > 0.8 {
						// to find similarity between description andattributeList and returing the similar end type inattributeList
						res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.7, edlib.Levenshtein)
						if err != nil {
							log.Println(err)
						}
						matchedEnd2 = res[0]
						fmt.Println("End Conn2:", convToupper, res[0])
					}
				}

			}
		}
	}
	return matchedEnd1, matchedEnd2
}
