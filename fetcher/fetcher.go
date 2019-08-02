package fetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("http get err:%s", err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("http get err:%s", err)
	}
	return bytes,nil
}
