``` go
package main

import (
    "fmt"
    "github.com/starxiang2/google-translate/translate"
    "github.com/starxiang2/google-translate/language"
)

func main() {
    fmt.Println(translate.Trans(language.CN,"or",[]string{ "大家好"}))
    fmt.Println(translate.Trans(language.EN,language.CN,[]string{ "Hello everyone"}))
}

```