package main

import (
	"github.com/starxiang2/google-translate/language"
	"github.com/starxiang2/google-translate/translate"
	"testing"
)

func TestTrans(t *testing.T) {
	for i := 1; i < 30; i++ {
		t.Log(translate.Trans(language.CN, language.EN, []string{"大家好"}))
	}
}
