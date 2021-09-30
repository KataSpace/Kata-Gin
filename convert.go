package KataGin

import (
	"fmt"
	"regexp"
	"strings"
)

// CommentConvert get api uri from comment
// There are has three conditions
// 1.
// // Hello
// function Hello(){}
// will convert to GET /Hello
// 2.
// // Hello
// // GET /
// will convert to GET /Hello
// 3.
// // Hello
// // GET /v1
// will conver to GET /v1/Hello
func CommentConvert(name, comment string) string {
	path := strings.Split(comment, "\n")
	if len(path) == 1 {
		path = append(path, "GET /")
	} else {
		if !strings.Contains(path[1], "/") {
			path[1] += " /"
		}
	}

	if strings.HasSuffix(strings.TrimSpace(path[1]), "/") {
		return fmt.Sprintf("%s%s", path[1], name)
	}

	return fmt.Sprintf("%s/%s", path[1], name)
}

// CommentGetMethods Spilt method and path from url
// Method and uri path must use space isolation
// e.g.
// GET /v1/Hello ==> GET    /v1/Hello
func CommentGetMethods(s string) (method string, name string) {
	data := strings.Split(s, " ")
	return strings.TrimSpace(data[0]), strings.TrimSpace(data[1])
}

func DefaultConvert(name string) string {

	return name
}

func DefaultGetMethods(s string) (method string, name string) {

	var b strings.Builder
	b.Grow(len(s))

	b.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == '/' || 'A' <= c && c <= 'Z' {
			name = s[i:]
			break
		}
		b.WriteByte(c)
	}

	return b.String(), name
}

func SlashConvertWithOne(s string) string {
	return slashConvert(s, 1)
}

func SlashConvertWithTwo(s string) string {
	return slashConvert(s, 2)
}

func SlashConvertWithThree(s string) string {
	return slashConvert(s, 3)
}

func SlashConvertWithFive(s string) string {
	return slashConvert(s, 5)
}

func SlashConvertWithTen(s string) string {
	return slashConvert(s, 10)
}

// slashConvert  add slash before every upper char
// n  Capital letters with less than N consecutive characters are ignored
// e.g.
// n = 5 GetAPIAllName will get Get、All、Name. API is three consecutive capital letters, which is less than the requirement of 5, so it will be ignored
// n = 2 GetAPIAllName will get Get、AP、IA、Name.
// more examples please referee convert_test.go
func slashConvert(s string, n int) string {
	r := "[A-Z][^A-Z]+"
	if n >= 1 {
		r = fmt.Sprintf("[A-Z][a-z]+|([A-Z]|[0-9]){%d}", n)
	}

	reg := regexp.MustCompile(r)

	sub := reg.FindAllString(s, -1)

	return strings.Join(sub, "/")
}
