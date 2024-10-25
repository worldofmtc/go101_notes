package main

import "fmt"

func main() {
	// // 1 DİZİNİN TANIMLANMASI
	// //  var diziadı = [uzunluğu]tipi{içini verilerlerle doldur}

	// var animals = [3]string
	// animals2 := [3]string{"koyun", "inek", "ceylan"}

	// //2 BİR DİZİYE DEĞER ATAMAK

	// //2.1 aynı satırda değer atamak
	// var animals = [3]string{"aslan", "kurt", "Kaplan"}

	// // 2.2 sonradan boş diziye değer atamak
	// ikininKatlari := [5]int{}
	// ikininKatlari[0] = 2
	// ikininKatlari[1] = 4
	// ikininKatlari[2] = 6
	// ikininKatlari[3] = 8
	// ikininKatlari[4] = 10

	// //DİZİLERİN UZUNLUĞU VE KAPASİTESİ
	ikininKatlari := [5]int{2}

	// len() bir dizinin  lenght - genişliği
	fmt.Println(len(ikininKatlari))

	// cap()  bir dizinin capacity - kapasitesi

	var dizi = [...]bool{true, false}
	fmt.Println(len(dizi))

}
