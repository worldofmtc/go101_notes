# Go101 Ders 1 - Değişken Tanımlama ve Yazdırma

Bu derste Go dilinde değişken tanımlama  `const` ve `var` kullanımını göreceğiz.

## Go Program Yapısı

Go dilinde bir program `main` paketinde yazılır ve çalıştırılabilir bir dosyada `main` fonksiyonu bulunur. Basit bir Go programı aşağıdaki gibi başlar:

```go
package main

func main() {
      
}

```

###Değişken Tanımlama Yöntemleri
1. Yöntem - Boş Değer Atama
Bu yöntemle değişken tipi belirtilir, ancak başlangıç değeri verilmez. Go, değişkene varsayılan bir değer atar:

`int` tipi için varsayılan değer: 0
`float64` tipi için varsayılan değer: 0.0
`bool` tipi için varsayılan değer: false
`string` tipi için varsayılan değer: "" (boş string)


```go

package main

import "fmt"

func main() {
    var a int
    var b bool
    var c string
    var d float64
    
    fmt.Println(a) // 0
    fmt.Println(b) // false
    fmt.Println(c) // ""
    fmt.Println(d) // 0.0
}
```

2. Yöntem - Değişkeni Tanımlayıp Başlangıç Değeri Atama
Bu yöntemde, `var` anahtar kelimesi kullanılarak değişken tanımlanır ve aynı anda başlangıç değeri atanır. Go, değişkenin türünü otomatik olarak belirler.

```go
package main

import "fmt"

func main() {
    var d int = 10
    var e bool = true
    var f string = "Merhaba"
    
    fmt.Println(d) // 10
    fmt.Println(e) // true
    fmt.Println(f) // Merhaba
}
```

3. Yöntem - := Kısa Tanımlama
Kısa tanımlama yöntemi ile var anahtar kelimesi kullanılmaz; değişken adı ve değeri doğrudan := ile atanır. Bu yöntem sadece fonksiyon içlerinde kullanılabilir.

```go
package main

import "fmt"

func main() {
    g := 20
    h := false
    i := "Dünya"
    
    fmt.Println(g) // 20
    fmt.Println(h) // false
    fmt.Println(i) // Dünya
}
```

4. Yöntem - Çoklu Tanımlama
Go dilinde çoklu değişken tanımlama, aynı veya farklı türden değişkenleri tek bir satırda tanımlamamıza olanak sağlar. Bu yöntemde değişkenlere başlangıç değerleri verilebilir veya boş değerler atanabilir.

```go


// Örnek 1: Farklı Türde Çoklu Tanımlama
package main

import "fmt"

func main() {
    var x, y, z = 10, "Merhaba", true

    fmt.Println(x) // 10
    fmt.Println(y) // Merhaba
    fmt.Println(z) // true
}
```


```go
//Örnek 2: Aynı Türde Çoklu Tanımlama
//Aynı türden birden fazla değişken tanımlarken türü bir kez belirterek var anahtar kelimesi ile tanımlayabiliriz.
package main

import "fmt"

func main() {
    var a, b, c int = 1, 2, 3

    fmt.Println(a) // 1
    fmt.Println(b) // 2
    fmt.Println(c) // 3
}
```


```go

//Örnek 3: Kısa Tanımlama ile Çoklu Tanımlama
//Kısa tanımlama (:=) ile de çoklu değişken tanımlaması yapılabilir.
package main

import "fmt"

func main() {
    p, q, r := 5, "Ders", false

    fmt.Println(p) // 5
    fmt.Println(q) // Ders
    fmt.Println(r) // false
}
```


# Yazdırma Seçenekleri
Go dilinde, `fmt` paketini kullanarak ekrana bilgi yazdırmak için çeşitli fonksiyonlar mevcuttur. En yaygın kullanılan fonksiyonlar:


fmt.Print(): Değişkenleri ve metinleri yan yana yazar, yeni satıra geçmez.

fmt.Println(): Değişkenleri ve metinleri yan yana yazar, sonuna otomatik olarak yeni bir satır ekler.

```go

fmt.Print("Merhaba, ")
fmt.Print("Dünya!")
// Çıktı: Merhaba, Dünya!


fmt.Println("Merhaba, Dünya!")
fmt.Println("Golang öğreniyorum.")
// Çıktı:
// Merhaba, Dünya!
// Golang öğreniyorum.


```


fmt.Printf(): Biçimlendirilmiş yazdırma yapmak için kullanılır. % ile belirtilen format belirteçlerini kullanarak çıktıyı özelleştirebiliriz.

### Temel Format Belirteçleri:
    * %d: Tamsayı (int)
    * %f: Ondalıklı sayı (float64)
    * %s: Dize (string)
    * %t: Boolean (bool)

```go
    var isim string = "Ahmet"
    var yas int = 25
    var maas float64 = 4000.50
    var aktif bool = true

    fmt.Printf("İsim: %s, Yaş: %d, Maaş: %.2f, Aktif: %t\n", isim, yas, maas, aktif)
    // Çıktı: İsim: Ahmet, Yaş: 25, Maaş: 4000.50, Aktif: true
```

Format Belirteçleri ile İleri Seviye Biçimlendirme

   * %v: Değişkenin varsayılan değerini yazdırır.
   * %T: Değişkenin türünü yazdırır.
   * %%: % karakterini yazdırmak için kullanılır.

```go

fmt.Printf("Yaş değişkeni: %v, Türü: %T\n", yas, yas)
// Çıktı: Yaş değişkeni: 25, Türü: int

```

