package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_isNumber(t *testing.T) {
	trueTests := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	for _, v := range trueTests {
		t.Run("true tests", func(t *testing.T) {
			if got := isNumber(v); got != true {
				t.Errorf("isNumber() = %v, want %v", got, true)
			}
		})
	}
	falseTests := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}
	for _, v := range falseTests {
		t.Run("false tests", func(t *testing.T) {
			if got := isNumber(v); got != false {
				t.Errorf("isNumber() = %v, want %v", got, false)
			}
		})
	}
}
