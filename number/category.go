package number

import "strings"

type Category string

func (c Category) String() string {
	return string(c)
}

const (
	CatDefault Category = "DEFAULT"
	CatFC2     Category = "FC2"
)

func IsFc2(number string) bool {
	number = strings.ToUpper(number)
	return strings.HasPrefix(number, "FC2")
}

func DecodeFc2ValID(n string) (string, bool) {
	if !IsFc2(n) {
		return "", false
	}
	idx := strings.LastIndex(n, "-")
	if idx < 0 {
		return "", false
	}
	return n[idx+1:], true
}
