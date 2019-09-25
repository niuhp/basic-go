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

func GetType2(p interface{}) string {
	switch v := p.(type) {
	case int:
		return "int:" + strconv.Itoa(v)
	case string:
		return "string:" + v
	}
	return "unknown"
}
