package parser

import (
	"crawler/engine"
	"regexp"
)

var
(
	cityRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(s []byte) engine.ParserResult {
	result := engine.ParserResult{}

	// 个人信息
	profileMatches := cityRe.FindAllSubmatch(s, -1)
	for _, m := range profileMatches {
		result.Items = append(result.Items, engine.Item{Name: string(m[2])})
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(s []byte) engine.ParserResult {
				return ProfileParser(s, url, name)
			},
		})
	}

	// 城市
	cityMatches := cityUrlRe.FindAllSubmatch(s, -1)
	for _, m := range cityMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[0]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
