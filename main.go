package main

type myData struct {
	someField int
}

func (data *myData) someMethod() {
	// intentional no-op
}

func standaloneFunc(data myData) {
	// intentional no-op
}

func main() {
	// intentional no-op
}
