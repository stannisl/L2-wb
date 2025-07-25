package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	res, _ := unpackStr("a4bc2d5e")
	res1, _ := unpackStr("qwe\\4\\5")

	fmt.Println(res, res == "aaaabccddddde")
	fmt.Println(res1, res1 == "qwe45")
}

func unpackStr(s string) (string, error) {
	var resstr strings.Builder

	var last rune
	var hasLast bool
	for _, i := range s {
		if unicode.IsDigit(i) && last != '\\' {
			if !hasLast {
				return "", errors.New("string cannot start from digit")
			}

			st := resstr.String()[:resstr.Len()-1]
			resstr.Reset()
			resstr.WriteString(st)

			amount, _ := strconv.Atoi(string(i))
			resstr.WriteString(strings.Repeat(string(last), amount))

			hasLast = false
		} else if i == '\\' {
			last = '\\'
		} else {
			resstr.WriteRune(i)
			last = i
			hasLast = true
		}
	}
	return resstr.String(), nil
}
