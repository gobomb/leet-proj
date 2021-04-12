package golang

import (
	"fmt"
	"os"
	"testing"

	"justest/pkg/golang/flager"
)

func TestFlagParse(t *testing.T) {
	flagVar := "flagset"
	os.Args = append(os.Args, "-testflag", flagVar)
	fmt.Printf("%v\n", os.Args)
	if flager.Run() != flagVar {
		t.Fail()
	}
}

//  go test -v pkg/golang/flager_test.go -test.run TestFlagParseFail
func TestFlagParseFail(t *testing.T) {
	r := flager.Run()
	t.Logf("the returned flag is %v\n", r)
	if r != "none" {
		t.Fail()
	}
}
