package src

import "fmt"

func ExecuteFuncWithMap(fn string) {
	a := func(name string) {
		fmt.Println("executou a ", name)
	}

	b := func(name string) {
		fmt.Println("executou b", name)
	}

	m := map[string]interface{}{
		"a": a,
		"b": b,
	}

	fx, ok := m[fn]

	if ok {
		fx.(func(string))("Tiago")
	} else {
		fmt.Println("função não encontrada")
	}
}
