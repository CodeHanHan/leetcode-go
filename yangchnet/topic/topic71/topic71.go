package topic71

import "strings"

func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	s := make([]string, 0)
	for _, v := range list {
		switch v {
		case ".", "": // 当前目录不变
			continue
		case "..": // 返回上一级目录
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		default:
			s = append(s, v)
		}
	}

	return "/" + strings.Join(s, "/")
}
