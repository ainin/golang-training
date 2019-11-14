package main

// // import(
// 	// "fmt"
// 	// "errors"
// 	// "day2/aritmatika"
// 	//"reflect"
// 	// "strings"
// // )

// //------4
// // func newCounter() func() int {
// // 	count := 0
// // 	return func() int {
// // 		count += 1
// // 		return count
// // 	}
// // }


// // func main(){
// 	// fmt.Println("Hello world!")
// 	// //cara 1
// 	// var name string
// 	// //cara 2
// 	// var nama string = "Ainin"
// 	// //cara 3
// 	// var nami, gender string
// 	// //cara 4
// 	// var namo, gendery string = "Ainin", "P"

// 	//=======================
// 	// //boolean
// 	// isLogin := false

// 	// //numeric
// 	// var a int = 15
// 	// b:= 20

// 	// //numeric float (c) & integer
// 	// c, d := 2.1, 30

// 	// //string
// 	// s := "Hello"

// 	// const pi= 3.14
// 	// fmt. Println(isLogin, a, b, c, d, s, pi)
// 	//============================
// 	//Number operation
// 	// a:=10
// 	// t:=15
// 	// L:=(a*t)/2
// 	// fmt.Println(L)
// 	// //String Operation
// 	// helloworld:="hello"+" "+"world"
// 	// fmt.Println(helloworld)
// 	//===========================
// 	//CHALLENGE DAY2 NO.4

// 	// var inputan1, inputan2 float64
	
// 	// fmt.Print("inputan1 : ")
// 	// fmt.Scanf("%v", &inputan1)
// 	// fmt.Print("inputan2: ")
// 	// fmt.Scanf("%v", &inputan2)
// 	// L := (inputan1*inputan2)/2
// 	// fmt.Println(L)


// 	//gapake inputan
// 	// a:=20
// 	// t:=25
// 	// L:=(a*t)/2
// 	// fmt.Println(L)


// 	//CHALLENGE DAY2 NO.5
// 	//WITHOUT INPUT
// 	// const pi = 3.14
// 	// t:=20.0
// 	// r:=4.0
// 	// L:=2*pi*r*(r+t)
// 	// fmt.Println(L)

// 	//WITH INPUTAN
// 	// var t, r float32
// 	// const pi = 3.14
// 	// fmt.Print("t : ")
// 	// fmt.Scanf("%v", &t)
// 	// fmt.Print("r: ")
// 	// fmt.Scanf("%v", &r)
// 	// Luas_permukaan := 2*pi*r*(r+t)
// 	// fmt.Println(Luas_permukaan)
// //==================
// //DAY 3
// //Bilangan Faktor
// // var input int
// // 	fmt.Print("masukan faktor bilangan : ")
// // 	fmt.Scanf("%d", &input)
// // for i :=1; i<=input; i++{
// // 	if input%i<=0{
// // 		fmt.Println(i)
// // 	}
// // 	}
// //==================
// // Bilangan Prima
// // var input int
// // 	fmt.Println("Bilangan Prima")
// // 	x := 0
// // 	fmt.Print("masukan angka: ")
// // 	fmt.Scanf("%d", &input)
// // for i :=1; i<=input; i++{
// // 	if input <= 1{
// // 		fmt.Println("Bukan bil prima")
// // 	} else if input % i == 0{
// // 		x++
// // 	}
// // }
// // 	if x == 2{
// // 		fmt.Println("Bilangan Prima")
// // 	} else{
// // 		fmt.Println("Bukan bilangan prima")
// // 	}


// //===========
// // PALINDROME
// 	// var a, b string
// 	// fmt.Println("Input: ")
// 	// fmt.Scanf("%s", &a)

// 	// for x:= len(a); x>0; x-- {
// 	// 	fmt.Println(string(a[x-1]))
// 	// 	b+= string(a[x-1])
// 	// }
// 	// if a == b{
// 	// 	fmt.Println("PALINDROME")
// 	// } else {
// 	// 	fmt.Println("BUKAN PALINDROME")
// 	// }

 
// //===========================================================================================================
// //ASTERIK
// // var a, i, j, k int
// // fmt.Println("input tinggi segitiga: ")
// // fmt.Scanf("%v", &a)
// // for i=1; i<=a; i++ {
// // for j=i; j<=a; j++ {
// // 	fmt.Printf(" ")
// // }
// // for k=1; k<=i; k++{
// // 	fmt.Printf("%s", "* ")
// // }
// // fmt.Print("\n")
// // }
// //===============
// //=====DAY 4===== 
// //Array
// // var countries[2] string
// // countries [0] = "India"
// // countries [1] = "Indonesia"
// // fmt.Println()

// //split
// // sentence:= "hello, evv"
// // var potongk= string(sentence[5])
// // fmt.Println(potongk)

// // var splitKarakter = strings.Split(sentence, "e")
// // fmt.Println(splitKarakter)


// //=======DAY 5=========

// // var name string = "John"
// // var nameAddress *string = &name
// // var Aku *string = nameAddress

// // fmt.Println("name (value)", name)
// // fmt.Println("name (address)", &name)
// // fmt.Println("nameAddress (value)", *nameAddress)
// // fmt.Println("nameAddress (address)", nameAddress)

// // name = "Doe"
// // fmt.Println("name (value)", name)
// // fmt.Println("name (address)", &name)
// // fmt.Println("nameAddress (value)", *nameAddress)
// // fmt.Println("nameAddress (address)", nameAddress)

// // name = "Ai"
// // fmt.Println("name (value)", name)
// // fmt.Println("name (address)", &name)
// // fmt.Println("Aku (value)", *Aku)
// // fmt.Println("Aku (address)", Aku)

// //-----
// // numberA := 25
// // var numberB *int
// // if numberB == nil {
// // 	fmt.Println("NumberB: ",numberB)
// // 	numberB = &numberA
// // 	fmt.Println("NumberB: ", *numberB)

// // }

// //------

// // var size = new(int)
// // fmt.Printf("Size (value) is %d \n", *size)
// // fmt.Printf("Size (Tye) is %T \n", size)
// // fmt.Printf("Size (Address) is %v \n", size)
// // *size = 11
// // fmt.Println("New size value is", *size)
// // var em *int = size
// // fmt.Println("New EM is", *em)
// // fmt.Println("New EM is", em)

// //--------1
// // hour:= 11
// // greeting(hour)

// //----2
// // p, l := 10, 5
// // luas, ket := calculate(p,l)
// // fmt.Println(luas, ket)

// //----3
// // hasil := sum(1,2,3,4,5)
// // fmt.Println(hasil)

// //----4
// // func(){
// // 	fmt.Println("funct pertama")
// // 	}()

// // 	value := func(){
// // 		fmt.Println("funct kedua")
// // 	}
// // 	value()

// // 	func(sentence string){
// // 		fmt.Println(sentence)
// // 	}("func ketiga")
// //------4
// // counter := newCounter()
// // fmt.Println(counter())
// // fmt.Println(counter())


// //-----5 [huda]
// // hasilTambah := aritmatika.Tambah(1,2)
// // fmt.Println("Hasil Tambah: ", hasilTambah)

// // hasilKurang := aritmatika.Kurang(5,2)
// // fmt.Println("Hasil Kurang: ", hasilKurang)

// //-----HUDA 2
// //DEFER
// defer catch()

// name := "" //kalo di isi akan ke line 277
// _, err:= validate(name)

// if err ==nil{
// 	fmt.Print("hello, ", name)
// }else{
// 	// fmt.Println("Error: ", err)
// 	panic(err.Error())
// } 
// fmt.Printf("%s \n", "end")

// } //end of main

// //-----HUDA 2

// func validate(name string) (bool, error){
// 	if name != ""{
// 		return true, nil
// 	}
// 	return false, errors.New("Nama tidak boleh kosong")
// }
// //-------HUDA 3
// func catch(){
// 	if r:= recover(); r!=nil{
// 		fmt.Println("Error: ", r)
// 	}
// }

// //------1
// // func greeting(hour int){
// // 	if hour < 12 {
// // 		fmt.Println("Selamat jam lebih dari 12")
// // 	} else {
// // 		fmt.Println("selamat")
// // 	}
// // }

// //----2
// // func calculate(panjang, lebar int) (luas int, ket string){ // luas, ket adalah yg di return
// // 	luas = panjang * lebar
// // 	//fmt.Println(luas, "1")
// // 	ket = "keterangan"
// // 	return luas, ket
// // }
// //
// //----3
// // func sum(numb ...int) int{
// // 	var total int = 0
// // 	banyakNumb := len(numb)
// // 	for _, number := range numb{
// // 		total += number
// // 	}
// // 	return total / banyakNumb
// // }
