package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(s []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(s, -1)

	result := engine.ParserResult{}

	for _, m := range matches {

		//if i==5{
		//	break
		//}

		result.Items = append(result.Items, engine.Item{Url: string(m[1]), Name: string(m[2])})
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		fmt.Printf("%s,%s\n", m[1], m[2])
	}
	return result
}
