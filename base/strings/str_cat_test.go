package base

import (
	"math"
	"testing"

	. "github.com/dayfine/base/base/testing"
)

// https://github.com/abseil/abseil-cpp/blob/master/absl/strings/str_cat_test.cc
func TestStrCatInts(t *testing.T) {
	// All types but int32 (rune), uint8 (byte), which is handled differently.
	const (
		c     int8    = -1
		uc    int8    = 2
		s     int16   = -3
		us    uint16  = 4
		i     int     = -5
		ui    uint    = 6
		uiptr uintptr = 7
		ul    uint32  = 8
		ll    int64   = -9
		ull   uint64  = 10
	)
	ExpectEq(t, StrCat(c, uc), "-12")
	ExpectEq(t, StrCat(s, us), "-34")
	ExpectEq(t, StrCat(i, ui), "-56")
	ExpectEq(t, StrCat(uiptr, ul), "78")
	ExpectEq(t, StrCat(ll, ull), "-910")
}

func TestStrCatEnums(t *testing.T) {
	const (
		golang int = iota
		java
		cplusplus

		doraemon = -1
		assembly = 1000
	)

	const (
		airbus uint64 = iota
		boeing        = 1000
		canary        = 10000000000
	)

	ExpectEq(t, StrCat(cplusplus), "2")
	ExpectEq(t, StrCat(golang, java), "01")
	ExpectEq(t, StrCat(doraemon, assembly), "-11000")
	ExpectEq(t, StrCat(canary), "10000000000")

	const (
		TwoToTheZero        int32 = 1
		TwoToTheSixteenth         = 1 << 16
		TwoToTheThirtyFirst       = math.MinInt32
	)

	ExpectEq(t, StrCat(TwoToTheSixteenth), "65536")
	ExpectEq(t, StrCat(TwoToTheThirtyFirst), "-2147483648")
}

func TestStrCatBasic(t *testing.T) {
	ExpectEq(t, StrCat(), "")
	// Note this is different from C++, where the expectation would be 0123.
	ExpectEq(t, StrCat(false, true, 2, 3), "falsetrue23")
	ExpectEq(t, StrCat(-1), "-1")
	ExpectEq(t, StrCat(0.5), "0.5")

	strs := []string{"Hello", "Cruel", "World"}
	byteSlices := [][]byte{
		{'H', 'e', 'l', 'l', 'o'},
		{'C', 'r', 'u', 'e', 'l'},
		{'W', 'o', 'r', 'l', 'd'}}
	runes := []rune{'日', '本', '語'}

	ExpectEq(t, StrCat(strs[0], strs[1], strs[2]), "HelloCruelWorld")
	ExpectEq(t, StrCat(byteSlices[0], " ", byteSlices[1], byteSlices[2]),
		"Hello CruelWorld")
	ExpectEq(t, StrCat(runes[0], runes[1], runes[2]), "日本語")

	ExpectEq(t, StrCat(strs[1], byteSlices[2]), "CruelWorld")
	ExpectEq(t, StrCat(strs[0], ", ", byteSlices[2]), "Hello, World")
	ExpectEq(t, StrCat(strs[0], ", ", strs[1], " ", strs[2], "!"), "Hello, Cruel World!")
	ExpectEq(t, StrCat(byteSlices[0], ", ", byteSlices[1], " ", byteSlices[2]),
		"Hello, Cruel World")

	i32s := []int{'H', 'C', 'W'}
	ui64s := []uint64{0xBadFace, 10987654321}

	ExpectEq(t, StrCat("ASCII ", i32s[0], ", ", i32s[1], " ", i32s[2], "!"),
		"ASCII 72, 67 87!")
	ExpectEq(t, StrCat(ui64s[0], ", ", ui64s[1], "!"),
		"195951310, 10987654321!")

	var f float32 = 100000.5
	ExpectEq(t, StrCat("A hundred K and a half is ", f),
		"A hundred K and a half is 100000")

	f = 100001.5
	ExpectEq(t, StrCat("A hundred K and one and a half is ", f),
		"A hundred K and one and a half is 100002")

	var d float64 = 100000.5
	d *= d
	ExpectEq(t, StrCat("A hundred K and a half squared is ", d),
		"A hundred K and a half squared is 1.00001e+10")

	ExpectEq(t, StrCat(1, 2, 333, 4444, 55555, 666666, 7777777, 88888888, 999999999),
		"12333444455555666666777777788888888999999999")
}

func TestStrCatMaxArgs(t *testing.T) {
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a"), "123456789a")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b"), "123456789ab")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c"), "123456789abc")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d"), "123456789abcd")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e"), "123456789abcde")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f"), "123456789abcdef")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g"), "123456789abcdefg")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h"), "123456789abcdefgh")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i"), "123456789abcdefghi")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j"), "123456789abcdefghij")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k"), "123456789abcdefghijk")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l"), "123456789abcdefghijkl")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m"), "123456789abcdefghijklm")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n"), "123456789abcdefghijklmn")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o"), "123456789abcdefghijklmno")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p"), "123456789abcdefghijklmnop")
	ExpectEq(t, StrCat(1, 2, 3, 4, 5, 6, 7, 8, 9, "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"), "123456789abcdefghijklmnopq")
	ExpectEq(t, StrCat(
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, "a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w",
		"x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"),
		"12345678910abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
