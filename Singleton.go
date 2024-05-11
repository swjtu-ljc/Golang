package main

type demo struct {
}

func newDemo() *demo {
	return &demo{}
}

var d *demo

func init() {
	d = newDemo()
}

func GetDemo() *demo {
	return d
}
