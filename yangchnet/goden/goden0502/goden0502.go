package goden0502

import "bytes"

func printBin(num float64) string {
	buf := bytes.Buffer{}
	buf.WriteString("0.")
	var s, c float64 = num * 2, num
	for i := 0; i < 32; i++ {
		if s == 1 {
			buf.WriteString("1")
			return buf.String()
		}

		if s > 1 {
			c = s - 1
			buf.WriteString("1")
		} else {
			c = s
			buf.WriteString("0")
		}

		s = c * 2
	}

	return "ERROR"
}
