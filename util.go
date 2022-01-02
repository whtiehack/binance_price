package main

// 从右向左，每隔 3 位，加一个单字符的逗号 ','
func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3]) + "," + comma(s[len(s)-3:])
}
