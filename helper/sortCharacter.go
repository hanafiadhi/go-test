package helper

import (
	"sort"
	"strings"
)

func FindVowelAndConsonants(str string)(vocal,consonants string) {
	str = strings.ToLower(str)
	vList:= []string{}
	CtList:= []string{}

	for _, v := range str {
		if Vowel(v) {
			vList = append(vList, string(v))
		}else if Consonant(v) {
			CtList =append(CtList, string(v))
		}
	}

	sort.Strings(vList)

	vocal = strings.Join(vList, "")
	consonants = strings.Join(CtList, "")
	return vocal, consonants

}

func Vowel(character rune) bool {
	vowel := "aiueo"
	return strings.ContainsRune(vowel,character)
}

func Consonant(char rune) bool  {
	consonants := "bcdfghjklmnpqrstvwxyz"
	return strings.ContainsRune(consonants, char)
}
