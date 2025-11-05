package main

import (
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// ポインター型を宣言
	var p *Person

	// ポインター型に値を代入しようとして、コンパイルエラーとなる
	// p = Person{Name: "Alice", Age: 30}

	// ポインター型に生成した値のポインターを代入する
	p = &Person{Name: "Alice", Age: 30}

	log.Println("p: ", p)
}