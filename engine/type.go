package engine

type ParserFunc func(s []byte) ParserResult
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParserResult struct {
	// agent
	Requests []Request
	// 用于存储有用信息
	Items [] Item
}

type Item struct {
	Url     string
	Name    string
	Id      string
	Type    string
	Payload interface{}
}
