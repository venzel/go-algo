/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve o problema de chamar uma função com um map de funções
 */

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
