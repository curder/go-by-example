package parser

import (
	"github.com/curder/go-by-example/src/others/cases/crawler/engine"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	var (
		content      []byte
		err          error
		result       engine.ParseResult
		resultSize   = 470
		expectedUrls = []string{
			"http://www.zhenai.com/zhenghun/aba",
			"http://www.zhenai.com/zhenghun/akesu",
			"http://www.zhenai.com/zhenghun/alashanmeng",
		}
		expectedCities = []string{
			"阿坝", "阿克苏", "阿拉善盟",
		}
		url  string
		city string
		i    int
	)

	// content,err = fetcher.Fetch("http://www.zhenai.com/zhenghun")
	content, err = ioutil.ReadFile("city_list_test_data.html")

	if err != nil {
		panic(err)
	}

	result = ParseCityList(content)

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; got %d", resultSize, len(result.Items))
	}

	for i, url = range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("exected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	for i, city = range expectedCities {
		if result.Items[i] != city {
			t.Errorf("exected city #%d: %s; but was %s", i, city, result.Items[i])
		}
	}
}
