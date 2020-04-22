package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	var (
		resp    *http.Response
		err     error
		content []byte
	)
	if resp, err = http.Get("http://www.zhenai.com/zhenghun"); err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	printCityList(content)
}

func printCityList(content []byte) {
	var (
		re      *regexp.Regexp
		matches [][][]byte
		match [][]byte
	)
	re = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches = re.FindAllSubmatch(content, -1)

	for _, match = range matches {
		fmt.Printf("City: %s Url: %s ", match[2], match[1])
		fmt.Println()
	}
	fmt.Printf("matchs found: %d\n", len(matches))
}
