package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

// url = <a href="http://album.zhenai.com/u/1961184609" target="_blank">夏末烟雨</a>
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`

func ParserCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
	}
	return result
}
