package engine

import "crawler/fetcher"

func worker(r Request) (ParserResult, error) {
	bytes, err := fetcher.Fetch(r.Url)
	if err != nil{
		return ParserResult{},err
	}
	parserResult := r.ParserFunc(bytes, r.Url)
	return parserResult,nil
}
