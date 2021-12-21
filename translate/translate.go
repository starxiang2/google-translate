package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var translate *Translate

func init() {
	translate = &Translate{}
}

func Trans(sourceLang string, targetLang string, text []string) ([]string, error) {
	return translate.Translate(sourceLang, targetLang, text)
}



type Translate struct {
}

func (this *Translate) Translate(sourceLang string, targetLang string, text []string) ([]string, error) {
	client := &http.Client{
		Timeout: 20*time.Second,
		Transport: &http.Transport{
			//Proxy: func(_ *http.Request) (*url.URL, error) {
			//	return url.Parse("socket5://127.0.0.1:7890")
			//},
		},
	}
	queryString := ""
	for _, v := range text {
		queryString += fmt.Sprintf("&q=%s", url.QueryEscape(v))
	}
	var urls []string = []string{
		"https://translate.googleapis.com",
		"https://translate.google.cn",
		"https://translate.google.com",
		"https://translate.google.jp",
		"https://translate.google.ca",
		"https://translate.google.it",
		"https://translate.google.us",
		"https://translate.google.com.hk",
		"https://translate.google.com.tw",
		"https://translate.google.com.ru",
		"https://translate.google.ae",
		"https://translate.google.co.in",
		"https://translate.google.de",
		"https://translate.google.com.sa",
		"https://translate.google.com.np",
		"https://translate.google.fr",
		"https://translate.google.co.uk",
		"https://translate.google.gr",
		"https://translate.google.pt",
		"https://translate.google.es",
		"https://translate.google.co.il",
		"https://translate.google.se",
		"https://translate.google.nl",
		"https://translate.google.be",
		"https://translate.google.at",
		"https://translate.google.pl",
		"https://translate.google.pl",
		"https://translate.google.co.th",
		"https://translate.google.com.sg",
		"https://translate.google.com.my",
		"https://translate.google.com.ru",
		"https://translate.google.es",
		"https://translate.google.pt",
		"https://translate.google.fi",
		"https://translate.google.ro",
		"https://translate.google.dk",
		"https://translate.google.no",
		"https://translate.google.com.au",
		"https://translate.google.co.nz",
		"https://translate.google.ca",
		"https://translate.google.com.br",
		"https://translate.google.com.ar",
		"https://translate.google.cl",
		"https://translate.google.com.pe",
		"https://translate.google.com.eg",
		"https://translate.google.com.pa",
		"https://translate.google.lt",
		"https://translate.google.bi",
		"https://translate.google.pn",
		"https://translate.google.li",
		"https://translate.google.com.nf",
		"https://translate.google.vg",
		"https://translate.google.mw",
		"https://translate.google.fm",
		"https://translate.google.sh",
		"https://translate.google.cd",
		"https://translate.google.ms",
		"https://translate.google.co.cr",
		"https://translate.google.lv",
		"https://translate.google.ie",
		"https://translate.google.co.je",
		"https://translate.google.com.py",
		"https://translate.google.gm",
		"https://translate.google.td",
		"https://translate.google.com.ua",
		"https://translate.google.co.ve",
		"https://translate.google.com.tr",
		"https://translate.google.com.mt",
		"https://translate.google.com.uy",
		"https://translate.google.hn",
		"https://translate.google.com.ni",
		"https://translate.google.gl",
		"https://translate.google.kz",
		"https://translate.google.sm",
		"https://translate.google.uz",
		"https://translate.google.rw",
		"https://translate.google.com.tj",
	}
	var resp *http.Response
	var err error
	for _, apiUrl := range urls {
		apiUrl = fmt.Sprintf("%s/translate_a/t?anno=3&client=te_lib&format=html&v=1.0&sl=%s&tl=%s&tk=%s&mode=1", apiUrl, sourceLang, targetLang, this.getTk(text))
		request, err := http.NewRequest("POST", apiUrl, strings.NewReader(strings.Trim(queryString, "&")))
		if err != nil {
			fmt.Println("err", err.Error())
			continue
		}
		request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36")
		resp, err = client.Do(request)
		if err == nil && resp.StatusCode == 200 {
			break
		}
	}
	if resp == nil {
		return nil, errors.New("请求接口失败")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var slice []string
	if len(text) == 1 {
		slice = append(slice, strings.Trim(string(body), `"`))
	} else {
		if err = json.Unmarshal(body, &slice); err != nil {
			return nil, errors.New("解析数据失败")
		}
	}

	res, _ := regexp.Compile("(?U)<i>.*</i>")
	for k, v := range slice {
		slice[k] = res.ReplaceAllString(v, "")
		slice[k] = strings.ReplaceAll(slice[k], "<b>", "")
		slice[k] = strings.ReplaceAll(slice[k], "</b>", "")
	}
	return slice, nil
}

func (this *Translate) getTk(text []string) string {
	str := strings.Join(text, "")
	salt := "406448.272554134"
	e := strings.Split(salt, ".")
	e0, _ := strconv.Atoi(e[0])
	value := e0
	for i := 0; i < len(str); i++ {
		value += int(str[i])
		value = this.hq(value, "+-a^+6")
	}

	value = this.hq(value, "+-3^+b+-f")
	e1, _ := strconv.Atoi(e[1])
	value ^= e1
	if 0 > value {
		value = (value & 2147483647) + 2147483648
	}
	x := value % 1e6
	return fmt.Sprintf("%s.%s", strconv.Itoa(x), strconv.Itoa((x ^ e0)))
}

func (this *Translate) bitwiseZFRS(char int, b int) int {
	if b == 0 {
		return char
	}
	return (char >> b) & (^(-9223372036854775808) >> (b - 1)) //-9223372036854775808
}

func (this *Translate) hq(char int, chunk string) int {
	for offset := 0; offset < (len(chunk) - 2); offset += 3 {
		b := int(chunk[offset+2])
		stringInt, err := strconv.Atoi(string(b))
		if err == nil {
			b = stringInt
		}

		if b >= int("a"[0]) {
			b = b - 87
		}

		if chunk[offset+1] == "+"[0] {
			b = this.bitwiseZFRS(char, b)

		} else {
			b = char << b
		}

		if chunk[offset] == "+"[0] {
			char = (char + b) & 4294967295
		} else {
			char = char ^ b
		}
	}
	return char
}
