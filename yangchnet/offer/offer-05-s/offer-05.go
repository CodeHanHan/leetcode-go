package offer_05_s

import "strings"

/**
@Author: lichang
@Date: 3/28/2021
@Version: 1.0
@Des: 替换空格
*/

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
