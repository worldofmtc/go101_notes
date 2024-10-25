package main

import "fmt"

func main() {

	// ŞART BLOKLARI

	// if - else - else if

	isim := "Talha"
	yas := 25
	var input int

	fmt.Println("lütfen yaşınızı yazınız")

	// Scan işlemleri
	fmt.Scan(&input)

	if input == yas {
		fmt.Printf("Merhaba %s yaşınız doğru", isim)
	}

}
