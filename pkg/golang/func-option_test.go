package golang

func Example_caller() {
	NewDuck(WithHeight(10),
		WithName("test")).Print()

	NewDuck().Print()
	// Output: test 10
	// default 1
}
