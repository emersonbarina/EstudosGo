package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
	generica(1.12345)

	// exemplo de função com interface
	fmt.Println(1, 2, "String", false, true, float64(12345))

	// desaconselhado
	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)

}
