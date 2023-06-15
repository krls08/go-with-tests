package main

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	//cases := []struct {
	//	Description string
	//	Arabic      int
	//	Want        string
	//}{
	//	{"convert 1 to I", 1, "I"},
	//	{"convert 2 to II", 2, "II"},
	//	{"convert 3 to III", 3, "III"},
	//	{"convert 4 to IV", 4, "IV"},
	//	{"convert 5 to V", 5, "V"},
	//	{"convert 6 to VI", 6, "VI"},
	//	{"convert 7 to VII", 7, "VII"},
	//	{"convert 8 to VIII", 8, "VIII"},
	//	{"convert 9 to IX", 9, "IX"},
	//	{"convert 10 to X", 10, "X"},
	//	{"convert 11 to XI", 11, "XI"},
	//	{"14 - XIV", 14, "XIV"},
	//	{"18 - XVIII", 18, "XVIII"},
	//	{"20 - XX", 20, "XX"},
	//	{"39 - XXXIX", 39, "XXXIX"},
	//	{"40 - XL", 40, "XL"},
	//	{"47 - XLVII", 47, "XLVII"},
	//	{"49 - XLIX", 49, "XLIX"},
	//	{"50 - L", 50, "L"},
	//}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertEqualsString(t, got, test.Roman)
		})
	}

}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func assertEqualsString(t testing.TB, got, want string) {

	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
