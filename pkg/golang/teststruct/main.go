package main

type A struct {
	b      *B
	target string
}

func (a *A) callA() {
	println(a.target)
}

type B struct {
	a *A
}

func (b *B) callB() {
	b.a.callA()
	println("b")
}

func main() {
	aa := &A{
		target: "only print by A",
	}
	bb := &B{
		a: aa,
	}
	aa.b = bb
	bb.callB()

	// refract

	a2 := &A2{
		target: "only print by A",
	}
	b2 := &B2{}
	a2.b = b2
	b2.callB(a2.getCallA())
}

func (a2 *A2) getCallA() func() {
	return a2.callA
}

type A2 struct {
	b      *B2
	target string
}

func (a *A2) callA() {
	println(a.target)
}

type B2 struct {
}

func (b *B2) callB(callA func()) {
	callA()
	println("b")
}
