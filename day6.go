package main

// import (
// 	"fmt"
// 	"math"
// )

// type Person struct{
// 	FirstName string
// 	LastName string
// 	Age int
// } 

// type person2 struct{
// 	FirstName string
// 	LastName string
// 	Tangan []Tangan

// } 

// type Tangan struct{
// 	Warna string
// 	Panjang int
// }
//------

// type Person struct{
// 	Name string
// 	Age int
// }

// func (P Person) GetName() string{
// 	return P.Name + "amazing!" 
// }

// func (P *Person) IncreaseAge(){
// 	P.Age = P.Age+1
// 	}
//-------
// type Rect struct{
// 	width float64
// 	height float64
// 	area float64
// }

// type Circle struct{
// 	radius float64
// }

// func (r Rect) Area() float64{
// 	return r.width * r.height
// }

// func (r *Rect) CalculateArea(){
// 	r.area = r.width * r.width
// }

// func (c Circle) Area() float64{
// 	return math.Pi * c.radius * c.radius
// } 

//-------------------
// type square struct{
// 	side int
// }

// type calculate Interface {
// 	large() int
// 	doubleSide() int
// }

// func (s square) large() int{
// 	return s.side * s.side
// }

// func (s square) doubleSide() int{
// 	return s.side *2 
// }


//func main(){
	// //long declaration
	// var person0 Person
	// person0.FirstName = "Ainin"
	// person0.LastName = "Asiyah"
	// person0.Age= 22
	// fmt.Println(person0.FirstName, person0.LastName, person0.Age)

	// //long declaration with assign value
	// var person1 = Person{"Lek", "Nur", 25}
	// fmt.Println(person1)

	// //long declaration with assigned value each name fields
	// // var person2 = Person{
	// // 	FirstName: "Fazri",
	// // 	LastName: "Niken",
	// // 	Age: 25,
	// // }
	// // fmt.Println(person2)

	// //ila struct dalam struct
	// var person5 = person2{
	// 	FirstName: "Ila",
	// 	LastName: "Umam",
	// 	Tangan: []Tangan{
	// 		Tangan{
	// 		Warna: "coklat",
	// 		Panjang: 35,
	// 		},
	// 	Tangan{
	// 		Warna: "kuning langsat",
	// 		Panjang: 40,
	// 		},
	// 	},
	// }
	// fmt.Println(person5)
	


	// //sort declaration
	// person3 := Person{"Mastin" , "Good", 23}
	// fmt.Println(person3)

	// //short declaration with new keyword
	// person4 := new(Person)
	// person4.FirstName = "Ainin"
	// person4.LastName = "Asiyah"
	// person4.Age = 22
	// fmt.Println(*person4)

	//-----------------------
// PersonA := Person{"John", 50}
// fmt.Println("%v\n", PersonA)
// fmt.Println(PersonA.Age)
// }

//-----------

// fmt.Println("")
// rect := Rect{
// 	width: 5.0,
// 	height: 4.0,
// }

// cir := Circle{5.0}
// fmt.Printf("Area rectangle = %0.1f\n", rect.Area())
// fmt.Printf("Area Circle = %f\n", cir.Area())

// fmt.Printf("Rectangle before callculate pointer method = %0.1f\n", rect.Area())
// rect.CalculateArea()
// fmt.Printf("Rectangle after callculate pointer method = %v\n", rect)

// fmt.Println("")

//-----------------

// var dimResult calculate
// dimResult = square{10}
// fmt.Println("Large square: ", dimResult.large())
// fmt.Println("Large square: ", dimResult.doubleSide())
// }	