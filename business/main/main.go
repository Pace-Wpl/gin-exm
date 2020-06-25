package main

func main() {
	err := initAll()
	if err != nil {
		panic(err.Error())
	}

	err = startAll()
	if err != nil {
		panic(err.Error())
	}

	defer close()

	for {
	}
}
