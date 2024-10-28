
# Go101 Ders 2 - Koşullu İfadeler: `if-else`, `else if`, `switch-case`

Bu derste Go dilinde koşullu ifadeleri kullanarak karar yapıları oluşturmayı öğreneceğiz. `if-else`, `else if` ve `switch-case` gibi yapılardan faydalanarak programlarımızda farklı durumlara göre çeşitli işlemler yapabiliriz.

---


## Koşullu İfadelerde Kullanılan Operatörler

1. Karşılaştırma Operatörleri
Bu operatörler iki değeri karşılaştırmak için kullanılır ve boolean (true veya false) bir sonuç döndürür.

  *  `==` : Eşitlik kontrolü
  *  `!=` : Eşit olmama durumu
  *  `>` : Büyüktür
  *  `<` : Küçüktür
  *  `>=` : Büyük veya eşittir
  *  `<=` : Küçük veya eşittir


````go

package main

import "fmt"

func main() {
    a := 10
    b := 20
                         // çıktılar: 
    fmt.Println(a == b)  // false
    fmt.Println(a != b)  // true
    fmt.Println(a > b)   // false
    fmt.Println(a < b)   // true
    fmt.Println(a >= b)  // false
    fmt.Println(a <= b)  // true
}

````

2. Mantıksal Operatörler
Bu operatörler birden fazla koşulu birleştirmek için kullanılır.

  *  `&&` : Ve (And) operatörü. Tüm koşullar true olduğunda true döner.
  *  `||` : Veya (Or) operatörü. En az bir koşul true olduğunda true döner.
  *  `!` : Değil (Not) operatörü. Sonucu tersine çevirir.



  ````go

  package main

    import "fmt"

    func main() {
        x := 15
        y := 30
                                       // çıktılar:
        fmt.Println(x > 10 && y > 20)  // true
        fmt.Println(x > 20 || y > 20)  // true
        fmt.Println(!(x > 10))         // false
    }

  ````



---

## 1. `if-else` Yapısı

Go dilinde `if-else` yapısı, bir koşulun doğru veya yanlış olmasına göre kod bloğunu çalıştırır. Koşul doğru (`true`) olduğunda `if` bloğundaki kodlar çalışır; aksi takdirde `else` bloğu çalışır.

**Örnek:**

```go
package main

import "fmt"

func main() {
    sayi := 10

    if sayi > 0 {
        fmt.Println("Pozitif bir sayı.")
    } else {
        fmt.Println("Negatif bir sayı veya sıfır.")
    }
}
```

## 2. `else if` Yapısı

`else if`, birden fazla koşul olduğunda kullanılır. Koşullar sırayla değerlendirilir; herhangi biri doğru olduğunda, o bloğun içindeki kod çalışır ve diğer `else if` veya `else` blokları atlanır.

**Örnek:**

```go
package main

import "fmt"

func main() {
    sayi := 0

    if sayi > 0 {
        fmt.Println("Pozitif bir sayı.")
    } else if sayi < 0 {
        fmt.Println("Negatif bir sayı.")
    } else {
        fmt.Println("Sayı sıfır.")
    }
}
```

## 3. `switch-case` Yapısı

`switch-case`, birden fazla durumu kontrol etmek için `else if` yerine kullanılabilir. `switch` ifadesi, farklı `case` koşullarını değerlendirir ve ilk eşleşen durumu çalıştırır.

**Temel Yapı:**

```go
package main

import "fmt"

func main() {
    gun := 3

    switch gun {
    case 1:
        fmt.Println("Pazartesi")
    case 2:
        fmt.Println("Salı")
    case 3:
        fmt.Println("Çarşamba")
    case 4:
        fmt.Println("Perşembe")
    case 5:
        fmt.Println("Cuma")
    default:
        fmt.Println("Hafta sonu")
    }
}
```

## 4. `switch-case` ile Birden Fazla Değer Eşleştirme

Bir `case` ifadesinde birden fazla değeri virgülle ayırarak aynı bloğu çalıştırmasını sağlayabilirsiniz.

**Örnek:**

```go
package main

import "fmt"

func main() {
    gun := "Cumartesi"

    switch gun {
    case "Cumartesi", "Pazar":
        fmt.Println("Hafta sonu.")
    default:
        fmt.Println("Hafta içi.")
    }
}
```

## 5. `switch` ile Koşul Kullanımı

`switch` ifadesinde yalnızca sabit değerlere değil, koşullara göre de kontrol yapılabilir.

**Örnek:**

```go
package main

import "fmt"

func main() {
    yas := 20

    switch {
    case yas < 18:
        fmt.Println("Çocuk")
    case yas >= 18 && yas < 65:
        fmt.Println("Yetişkin")
    default:
        fmt.Println("Yaşlı")
    }
}
```

---

Bu dosyada Go dilinde `if-else`, `else if` ve `switch-case` yapılarıyla koşullu ifadeleri kullanmayı öğrendik. Bu yapıları, programlarınızda karar vermek ve farklı koşullara göre işlemler yapmak için kullanabilirsiniz.
