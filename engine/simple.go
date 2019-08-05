package engine

import (
	"fmt"
)

type SimpleEngine struct {
}

func (s *SimpleEngine) Run(sends ...Request) {
	var requests []Request
	for _, row := range sends {
		requests = append(requests, row)
	}

	for len(requests) > 0 {
		row := requests[0]
		requests = requests[1:]

		result, err := worker(row)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)

		fmt.Printf("%s:\n %v\n\n", row.Url,result.Items)
	}
}
