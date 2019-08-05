package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// simple版
	//e := &engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	currentEngine := &engine.ConCurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		ChanCount: 10,
	}
	currentEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
