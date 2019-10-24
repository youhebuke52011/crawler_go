package persist

import (
	"crawler/engine"
	es "gopkg.in/olivere/elastic.v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// 返回回去的channel
	out := make(chan engine.Item)

	client, err := es.NewClient(es.SetSniff(false))
	if err != nil {

	}

	// 初始化es

	go func() {
		//item := <-out
	}()

	return out, nil
}
