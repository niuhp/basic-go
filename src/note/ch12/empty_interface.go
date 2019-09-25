package ch12

import "strconv"

func GetType(p interface{}) string {
	if i, ok := p.(int); ok {
		return "int:" + strconv.Itoa(i)
	}
	if s, ok := p.(string); ok {
		return "string:" + s
	}
	return "unknown"
}
