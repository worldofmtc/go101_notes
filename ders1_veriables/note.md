# Go101 Ders 1 - Değişken Tanımlama

Bu derste Go dilinde değişken tanımlama ve `const` kullanımını göreceğiz.

## Go Program Yapısı

Go dilinde bir program `main` paketinde yazılır ve çalıştırılabilir bir dosyada `main` fonksiyonu bulunur. Basit bir Go programı aşağıdaki gibi başlar:

```go
package main

func main() {
      
}

```

Değişken Tanımlama Yöntemleri
Go dilinde değişkenler üç farklı şekilde tanımlanabilir:

1. Yöntem - Boş Değer Atama
Bu yöntemle değişken tipi belirtilir, ancak başlangıç değeri verilmez. Go, değişkene varsayılan bir değer atar:

int tipi için varsayılan değer: 0
bool tipi için varsayılan değer: false
string tipi için varsayılan değer: "" (boş string)
Örnek: