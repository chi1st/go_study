package parser

import (
	"learngo/crawler/engine"
	//"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {

	//x := `<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]</a>`
	//x := `<a href="http://www.zhenai.com/zhenghun/yuxi" data-v-1573aa7c>玉溪</a>`
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
		//fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
	}
	return result

}
