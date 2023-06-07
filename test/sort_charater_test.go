package test

import (
	"Go-test/helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingCharacter(t *testing.T) {
	character :=  "Sample Case"
	vokal,konsonan := helper.FindVowelAndConsonants(character)
	var numFamilies int
	fmt.Print("Input the number of families: ")
	fmt.Scanln(&numFamilies)
	assert.Equal(t,"aaee",vokal,"Result must be 'aaee'")
	assert.Equal(t,"smplcs",konsonan,"Result must be 'smplcs'")
	fmt.Println("Testing Done")
}
