package core_use

type human interface {
}

func main() {
	var emptyF interface{}
	println(emptyF)
	var peopleF human
	println(peopleF)
}
