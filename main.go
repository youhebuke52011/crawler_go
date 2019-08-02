package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	// simple版
	e := &engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
