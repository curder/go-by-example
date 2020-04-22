package main

import (
	"github.com/curder/go-by-example/src/others/cases/crawler/engine"
	"github.com/curder/go-by-example/src/others/cases/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
