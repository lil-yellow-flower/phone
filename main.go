package main

import (
	// "bytes"
	"regexp"
	"strconv"
)


func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")

	res := re.ReplaceAllString(phone, "")

	i, err := strconv.Atoi(res[:2])
	if err == nil {
		if i == 46 {
			len := len([]rune(res))
			res = "0" + res[2:len]
		}
	}

	return res
}



// func normalize(phone string) string {
// 	var buf bytes.Buffer

// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}

// 	return buf.String()
// }

