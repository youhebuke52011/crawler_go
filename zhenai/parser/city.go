package parser

import (
	"crawler/engine"
	"regexp"
)

var
(
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(s []byte, url string) engine.ParserResult {
	result := engine.ParserResult{}
	matches := cityRe.FindAllSubmatch(s,-1)

	for _,m := range matches{
		result.Items = append(result.Items,engine.Item{Name:string(m[2])})
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: ProfileParser,
		})
	}
	return result
}
