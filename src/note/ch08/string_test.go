package ch08

import (
	"testing"
	"strings"
)

func TestStringInit(t *testing.T) {
	s := "\xee\xee\xee"
	t.Log(s, "len =", len(s))
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for i, c := range s {
		t.Logf("i=%[1]d,v=%[2]c | %[2]x", i, c)
	}
}
func TestStringTool(t *testing.T) {
	s := "中,华人,民共,和国"
	arr := strings.Split(s, ",")
	for i, c := range arr {
		t.Logf("i=%d,v=%s", i, c)
	}
}
