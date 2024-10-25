package main

import "fmt"

func main() {
	// Değişken tanımlama işlemleri

	// İLK YÖNTEM
	var isim string
	var yas int
	var student bool

	isim = "ahmet"
	yas = 15
	student = true

	// İKİNCİ YÖNTEM
	var isim2 string = "Muhammed"
	var para float32 = 99.99

	// ÜÇÜNCÜ YÖNTEM
	isim3 := "Muhammed Talha"
	boy := 185

	// CONST LAR
	// Sabit değerler için geçerlidir, program boyunca değişmeyecektir.
	const pi float32 = 3.14

	// Değerleri yazdırma
	fmt.Println("İsim 1:", isim)
	fmt.Println("Yaş:", yas)
	fmt.Println("Öğrenci mi?", student)
	fmt.Println("İsim 2:", isim2)
	fmt.Println("Para:", para)
	fmt.Println("İsim 3:", isim3)
	fmt.Println("Boy:", boy)
	fmt.Println("Pi değeri:", pi)
}
