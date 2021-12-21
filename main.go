package main

import (
	"fmt"
	"github.com/starxiang2/google-translate/language"
	"github.com/starxiang2/google-translate/translate"
)

func main() {
	fmt.Println(translate.Trans(language.CN, "or", []string{"大家好"}))
	fmt.Println(translate.Trans(language.EN, language.CN, []string{"Hello everyone"}))
}
