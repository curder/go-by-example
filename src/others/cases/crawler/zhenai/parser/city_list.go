package parser

import (
	"github.com/curder/go-by-example/src/others/cases/crawler/engine"
	"regexp"
)

const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	var (
		re      *regexp.Regexp
		matches [][][]byte
		match   [][]byte
		result  engine.ParseResult
	)
	re = regexp.MustCompile(cityListRegex)
	matches = re.FindAllSubmatch(content, -1)

	for _, match = range matches {
		result.Items = append(result.Items, string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: engine.NilParser,
		})
		// fmt.Printf("City: %s Url: %s ", match[2], match[1])
		// fmt.Println()
	}
	return result
	// fmt.Printf("matchs found: %d\n", len(matches))
}
