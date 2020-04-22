package engine

import (
	"github.com/curder/go-by-example/src/others/cases/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var (
		requests     []Request
		request      Request
		parserResult ParseResult
		body         []byte
		err          error
		item         interface{}
	)
	for _, request = range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		request = requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", request.Url)
		if body, err = fetcher.Fetch(request.Url); err != nil {
			log.Printf("Fetcher error: fetching url %s: %v", request.Url, err.Error())
			continue
		}
		parserResult = request.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)
		for _, item = range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
