// 标准库的 strings 包提供了很多有用的字符串相关的函数
package main

import (
	"fmt"
	"strings"
)

// 给 `fmt.Println` 一个短名字的别名，随后将会经常用到
var p = fmt.Println

func main() {

	// 这是一些 `strings` 中的函数例子。注意他们都是包中的函数，不是字符串对象自身的方法，这意味着我们需要考虑在调用时传递字符作为第一个参数进行传递
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p()

	// 可以在 [`strings`](http://golang.org/pkg/strings/) 包文档中找到更多的函数
	// 虽然不是 `strings` 的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
