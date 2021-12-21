``` go
package main

import (
	"fmt"
	"github.com/starxiang2/google-translate/language"
	"github.com/starxiang2/google-translate/translate"
)

func main() {
	trans := translate.New()
	trans.SetProxy("http://127.0.0.1:7890")
	fmt.Println(trans.Translate(language.CN, "or", []string{"大家好"}))
}

```