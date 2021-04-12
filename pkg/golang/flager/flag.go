package flager

import (
	"flag"
	"fmt"
)

var Ftestflag = flag.String("testflag", "none", "just test flag")

func Run() string {
	flag.Parse()
	fmt.Printf("the flag testflag is %v\n", *Ftestflag)
	return *Ftestflag
}
