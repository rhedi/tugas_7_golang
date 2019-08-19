package main

import "fmt"
import "reflect"
import "runtime"

func main()  {
  var number = 20
  var huruf = "Ini Siapa Coba"
  var reflectnumber = reflect.ValueOf(number)
  var reflecthuruf = reflect.ValueOf(huruf)

  fmt.Println("Tipe Variabel :", reflectnumber.Type())
  fmt.Println("Tipe Variabel :", reflecthuruf.Type())

  if reflectnumber.Kind() == reflect.Int{
    fmt.Println("Nilai Variabel :", reflectnumber.Int())
  }
  if reflecthuruf.Kind() == reflect.String{
    fmt.Println("Nilai Variabel :", reflecthuruf.String())
  }

  runtime.GOMAXPROCS(2)

  go bacatipe(5, "Saya Kirim")
  bacatipe(5, "Saya Terima")

  var input string
  fmt.Scanln(&input)
}

func bacatipe(x int, pesan string)  {
  for i:=0;i<x;i++{
    fmt.Println(i+1, pesan)
  }
}
