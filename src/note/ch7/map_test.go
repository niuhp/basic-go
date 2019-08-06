package ch7

import "testing"

func TestMapInit(t *testing.T) {

	m1 := map[int]int{1: 10, 2: 20, 3: 30}
	t.Log("len =", len(m1), ",m1 =", m1)

	m2 := map[int]string{}
	m2[8] = "haha"
	t.Log("len =", len(m2), ",m2 =", m2)

	m3 := make(map[string]int, 7)
	m3["a"] = 8
	t.Log("len =", len(m3), ",m3 =", m3, "m3[\"b\"] =", m3["b"])

	if v, ok := m3["b"]; ok {
		t.Log("exist,m3[\"b\"] =", v)
	} else {
		t.Log("not exist,m3[\"b\"] =", v)
	}

}

func TestMapTravel(t *testing.T) {
	m := map[string]string{"a": "aaa", "b": "bsd"}
	for k, v := range m {
		t.Log("k =", k, ",v =", v)
	}
}
