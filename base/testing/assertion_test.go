package btest

import "testing"

func TestExpect(t *testing.T) {
	Expect(t, true, "True")
	Expect(t, !false, "Not false")
	Expect(t, 1+1 == 2, "1+1 == 2")

	dict := map[string]string{
		"exists": "yes",
	}

	var ok bool
	_, ok = dict["exists"]
	Expect(t, ok, "Key exists")
	_, ok = dict["doesNotExists"]
	Expect(t, !ok, "Key does not exist")
}

func TestExpectEq(t *testing.T) {
	ExpectEq(t, "a", "a")
	ExpectEq(t, 3+-3, 0)
	ExpectEq(t, 123.00, 123.0)
	ExpectEq(t, true, !false)

	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	ExpectEq(t, m1, m2)
}

func TestExpectEqStruct(t *testing.T) {
	type s struct {
		A, b int
		C    []int
	}
	a := s{1, 2, []int{1}}
	b := s{1, 2, []int{1}}
	ExpectEqStruct(t, a, b)
}

func TestAssert(t *testing.T) {
	Assert(t, true, "True")
	Assert(t, !false, "Not false")
	Assert(t, 1+1 == 2, "1+1 == 2")
}

func TestAssertOk(t *testing.T) {
	AssertOk(t, nil)
}
