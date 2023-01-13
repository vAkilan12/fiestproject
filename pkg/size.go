package pkg

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/hbollon/go-edlib"
)

func GetSize(descriptionArray []string, attributeList []string, name string) {
	// Using for Loop
	for descriptionKey := 1; descriptionKey < len(descriptionArray); descriptionKey++ {
		// Spliting the description after "Comma"
		splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
		for _, str := range splitAftercomma {
			//Spliting the description after "Space"
			splitAfterspace := strings.Fields(str)
			for _, str2 := range splitAfterspace {
				// removing all the spaces
				// Converting description to uppercase
				convToupper := strings.ToUpper(strings.ReplaceAll(str2, " ", ""))
				if edlib.JaroSimilarity(convToupper, name) >= 0.85 {
					for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
						//6"
						// Extracting the above Size by using the regexp expression
						FindSingleSize(descriptionArray, attributeList, descriptionKey, attributeListkey,
							`(?m)(?P<size2>\d\W)`, 0.5, 0.9)

						//1/2"
						// Extracting the above Size by using the regexp expression
						FindSingleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d+)(?P<UoM1>\W)`, 0.5, 0.9)

						//3/4INCH
						// Extracting the above Size by using the regexp expression
						FindSingleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d+)(?P<UoM1>\w+)`, 0.1, 0.1)

						//1/4MTx3inch
						// Extracting the above Size by using the regexp expression
						FindSingleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\w+)`, 0.5, 0.9)

						//8"X1/4"
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d)(?P<UoM1>\w+|\W)(X)(?P<size2>\d\W\d)(?P<Uom2>\w+|\W)`, `(?m)(?P<size1>\d)(?P<UoM1>\w+|\W)(X)`,
							`(?m)(X)(?P<size2>\d\W\d)(?P<Uom2>\w+|\W)`, 0.5, 0.5)

						//2INCHMIPTX2INCH
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\w+)(X)(?P<size2>\d\w+)`, `(?m)(?P<size1>\d\w+)(X)`,
							`(?m)(X)(?P<size2>\d\w+)`, 0.1, 0.1)

						//1"NPTX1"T
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d)(?P<UoM1>\W\w+)(X)(?P<size2>\d)(?P<Uom2>\W\w+)`,
							`(?m)(?P<size1>\d)(?P<UoM1>\W\w+)(X)`, `(?m)(X)(?P<size2>\d)(?P<Uom2>\W\w+)`, 0.1, 0.1)

						//M16X1/2"inch
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w\d+)(X)(?P<size2>\d\W\d\W\w+)`, `(?m)(?P<size1>\w\d+)(X)`,
							`(?m)(X)(?P<size2>\d\W\d\W\w+)`, 0.1, 0.1)

						//3/4INCHFNPTX3/4INCH
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\w+)(X)(?P<size2>\d\W\w+)`, `(?m)(?P<size1>\d\W\d\w+)(X)`,
							`(?m)(X)(?P<size2>\d\W\w+)`, 0.1, 0.1)

						//M16X1/2"
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w\d+)(X)(?P<size2>\d\W\d\W)`, `(?m)(?P<size1>\w\d+)(X)`,
							`(?m)(X)(?P<size2>\d\W\d\W)`, 0.1, 0.1)

						//1/2MNPTXM18
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\w+)(X)(?P<size2>\w\d+)`, `(?m)(?P<size1>\d\W\d\w+)(X)`,
							`(?m)(X)(?P<size2>\w\d+)`, 0.3, 0.1)

						//2SX1S
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\w)(X)(?P<size2>\d\w)`, `(?m)(?P<size1>\d\w)(X)`,
							`(?m)(X)(?P<size2>\d\w)`, 0.3, 0.1)

						//M16XM20
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w\d+)(X)(?P<size2>\w\d+)`, `(?m)(?P<size1>\w\d+)(X)`,
							`(?m)(X)(?P<size2>\w\d+)`, 0.3, 0.1)

						//3/4INCHFNPTX3Inch
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\W\d)(?P<UoM1>\w+|\W)(X)(?P<size2>\w+)`, `(?m)(?P<size1>\d\W\d\W\d)(?P<UoM1>\w+|\W)(X)`,
							`(?m)(X)(?P<size2>\w+)`, 0.3, 0.1)

						//1/4"MTX3/8"Barb
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d)(?P<UoM1>\W\w+|\W)(X)(?P<size2>\d\W\d)(?P<Uom2>\W\w+|\W)`, `(?m)(?P<size1>\d\W\d)(?P<UoM1>\W\w+|\W)(X)`,
							`(?m)(X)(?P<size2>\d\W\d)(?P<Uom2>\W\w+|\W)`, 0.3, 0.1)

						//3/4"MALEJICX3/4"MNP
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\d\W\d\W\w+)`, `(?m)(?P<size1>\d\W\d\W\w+)(X)`,
							`(?m)(X)(?P<size2>\d\W\d\W\w+)`, 0.2, 0.1)

						//M20MALEXM25FEMALE
						// Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\w\d+\w+)(X)(?P<size2>\w\d+\w+)`, `(?m)(?P<size1>\w\d+\w+)(X)`,
							`(?m)(X)(?P<size2>\w\d+\w+)`, 0.2, 0.1)

						//1/2"X1/4"
						//Extracting the above Size by using the regexp expression
						FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, `(?m)(?P<size1>\d\W\d\W)(X)(?P<size2>\d\W\d\W)`, `(?m)(?P<size1>\d\W\d\W)(X)`,
							`(?m)(X)(?P<size2>\d\W\d\W)`, 0.5, 0.5)

					}

				}

			}
		}
	}
}

func FindDoubleSize(descriptionArray, attributeList []string, descriptionKey, attributeListkey int,
	regExp, regExp1, regExp2 string, searchThreshold1, searchThreshold2 float32) (string, string) {
	var matchedSize1, matchedSize2 string
	var re = regexp.MustCompile(regExp)
	convToupper := strings.ToUpper(strings.ReplaceAll(descriptionArray[descriptionKey], " ", ""))
	for _, match := range re.FindAllString(convToupper, -1) {
		var re1 = regexp.MustCompile(regExp1)
		var re2 = regexp.MustCompile(regExp2)
		for _, match1 := range re1.FindAllString(match, -1) {
			if edlib.JaroSimilarity(match1, attributeList[attributeListkey]) > 0.4 {
				// to find similarity between description and attributeList and returing the similar Size in attributeList
				res1, err := edlib.FuzzySearchSetThreshold(match1, attributeList, 1, searchThreshold1, edlib.Levenshtein)
				if err != nil {
					log.Println(err)
				}
				matchedSize1 = res1[0]
				fmt.Println("End Size1:", match, res1[0])
			}
		}
		for _, match2 := range re2.FindAllString(match, -1) {
			if edlib.JaroSimilarity(match2, attributeList[attributeListkey]) > 0.4 {
				res2, err := edlib.FuzzySearchSetThreshold(match2, attributeList, 1, searchThreshold2, edlib.Levenshtein)
				if err != nil {
					log.Println(err)
				}
				matchedSize2 = res2[0]
				fmt.Println("End Size2:", match, res2[0])
			}
		}
	}
	return matchedSize1, matchedSize2
}

func FindSingleSize(descriptionArray, attributeList []string, descriptionKey, attributeListkey int,
	regExp string, searchThreshold, jaroThreshold float32) string {
	var matchedSize string
	var re = regexp.MustCompile(regExp)
	str := strings.ToUpper(strings.ReplaceAll(descriptionArray[descriptionKey], " ", ""))
	for _, match := range re.FindAllString(str, -1) {
		if edlib.JaroSimilarity(match, attributeList[attributeListkey]) > jaroThreshold {
			// to find similarity between description and attributeList and returing the similar Size in attributeList
			res, err := edlib.FuzzySearchSetThreshold(match, attributeList, 1, searchThreshold, edlib.Levenshtein)
			if err != nil {
				log.Println(err)
			}
			matchedSize = res[0]
			fmt.Println("size1:", match, res[0])
		}
	}
	return matchedSize
}

var singleSizeexp = []string{
	`(?m)(?P<size2>\d\W)`,
	`(?m)(?P<size1>\d\W\d+)(?P<UoM1>\W)`,
	`(?m)(?P<size1>\d\W\d+)(?P<UoM1>\w+)`,
	`(?m)(?P<size1>\d\W\d\w+)`,
}

var similarityThreshold1 = []float32{
	0.5, 0.5, 0.1, 0.5,
}

var similarityThreshold2 = []float32{
	0.9, 0.9, 0.1, 0.9,
}

var regExp = []string{
	`(?m)(?P<size1>\d)(?P<UoM1>\w+|\W)(X)(?P<size2>\d\W\d)(?P<Uom2>\w+|\W)`,
	`(?m)(?P<size1>\d\w+)(X)(?P<size2>\d\w+)`,
	`(?m)(?P<size1>\d)(?P<UoM1>\W\w+)(X)(?P<size2>\d)(?P<Uom2>\W\w+)`,
	`(?m)(?P<size1>\w\d+)(X)(?P<size2>\d\W\d\W\w+)`,
	`(?m)(?P<size1>\d\W\d\w+)(X)(?P<size2>\d\W\w+)`,
	`(?m)(?P<size1>\w\d+)(X)(?P<size2>\d\W\d\W)`,
	`(?m)(?P<size1>\d\W\d\w+)(X)(?P<size2>\w\d+)`,
	`(?m)(?P<size1>\d\w)(X)(?P<size2>\d\w)`,
	`(?m)(?P<size1>\w\d+)(X)(?P<size2>\w\d+)`,
	`(?m)(?P<size1>\d\W\d\W\d)(?P<UoM1>\w+|\W)(X)(?P<size2>\w+)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)(?P<size2>\d\W\d\W\w+)`,
	`(?m)(?P<size1>\w\d+\w+)(X)(?P<size2>\w\d+\w+)`,
	`(?m)(?P<size1>\d\W\d\W)(X)(?P<size2>\d\W\d\W)`,
	`(?m)(?P<size1>\d\W\d)(?P<UoM1>\W\w+|\W)(X)(?P<size2>\d\W\d)(?P<Uom2>\W\w+|\W)`,
}

var regExp1 = []string{
	`(?m)(?P<size1>\d)(?P<UoM1>\w+|\W)(X)`,
	`(?m)(?P<size1>\d\w+)(X)`,
	`(?m)(?P<size1>\d)(?P<UoM1>\W\w+)(X)`,
	`(?m)(?P<size1>\w\d+)(X)`,
	`(?m)(?P<size1>\d\W\d\w+)(X)`,
	`(?m)(?P<size1>\w\d+)(X)`,
	`(?m)(?P<size1>\d\W\d\w+)(X)`,
	`(?m)(?P<size1>\d\w)(X)`,
	`(?m)(?P<size1>\w\d+)(X)`,
	`(?m)(?P<size1>\d\W\d\W\d)(?P<UoM1>\w+|\W)(X)`,
	`(?m)(?P<size1>\d\W\d\W\w+)(X)`,
	`(?m)(?P<size1>\w\d+\w+)(X)`,
	`(?m)(?P<size1>\d\W\d\W)(X)`,
	`(?m)(?P<size1>\d\W\d)(?P<UoM1>\W\w+|\W)(X)`,
}

var regExp2 = []string{
	`(?m)(X)(?P<size2>\d\W\d)(?P<Uom2>\w+|\W)`,
	`(?m)(X)(?P<size2>\d\w+)`,
	`(?m)(X)(?P<size2>\d)(?P<Uom2>\W\w+)`,
	`(?m)(X)(?P<size2>\d\W\d\W\w+)`,
	`(?m)(X)(?P<size2>\d\W\w+)`,
	`(?m)(X)(?P<size2>\d\W\d\W)`,
	`(?m)(X)(?P<size2>\w\d+)`,
	`(?m)(X)(?P<size2>\d\w)`,
	`(?m)(X)(?P<size2>\w\d+)`,
	`(?m)(X)(?P<size2>\w+)`,
	`(?m)(X)(?P<size2>\d\W\d\W\w+)`,
	`(?m)(X)(?P<size2>\d\W\d\W\w+)`,
	`(?m)(X)(?P<size2>\w\d+\w+)`,
	`(?m)(X)(?P<size2>\d\W\d\W)`,
	`(?m)(X)(?P<size2>\d\W\d)(?P<Uom2>\W\w+|\W)`,
}

var similarity1 = []float32{
	0.1, 0.1, 0.1, 0.1, 0.1, 0.5, 0.3, 0.3, 0.3, 0.3, 0.3, 0.2, 0.2, 0.5,
}

var similarity2 = []float32{
	0.5, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.5,
}

func FindSize(descriptionArray, attributeList []string, descriptionKey int) (string, string) {
	var matchedSize1, matchedSize2 string
	for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
		for sizeExpkey := 0; sizeExpkey < len(singleSizeexp); sizeExpkey++ {
			//6"
			// Extracting the above Size by using the regexp expression
			string1 := FindSingleSize(descriptionArray, attributeList, descriptionKey, attributeListkey, singleSizeexp[sizeExpkey],
				similarityThreshold1[sizeExpkey], similarityThreshold2[sizeExpkey])
			if len(string1) > 0 {
				matchedSize1 = string1
			}
		}
		for regExplistKey := 0; regExplistKey < len(regExp); regExplistKey++ {
			string1, string2 := FindDoubleSize(descriptionArray, attributeList, descriptionKey, attributeListkey,
				regExp[regExplistKey], regExp1[regExplistKey], regExp2[regExplistKey], similarity1[regExplistKey], similarity2[regExplistKey])
			if len(string1) > 0 {
				matchedSize1 = string1
				matchedSize2 = string2
			}
		}
	}
	return matchedSize1, matchedSize2
}
