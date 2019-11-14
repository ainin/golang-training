package main

// import (
// "encoding/json"
// "net/http"
// "fmt"
// "io/ioutil"
// )

// //================NO 2========================
// type suku struct{
// 	ID		int
// 	Title	string
// 	Content	string
// }

// // func Suku(w http.ResponseWriter, r *http.Request){
// // 	var data = []suku{
// // 		suku{1, "Jawa", "Jawa Timur"},
// // 		suku{2, "Borneo ", "Kalimantan"},
// // 	}
	
// // 	w.Header().Set("Content-Type", "application/json")
// // 	if r.Method == "GET" {
// // 		result, err := json.Marshal(data)
// // 		if err != nil{
// // 			http.Error(w, err.Error(), http.StatusInternalServerError)
// // 			return
// // 		}
// // 		w.Write(result)
// // 		return
// // 	}
// // http.Error(w, "", http.StatusBadRequest) 
// // return
// // }

// // func hello(w http.ResponseWriter, r *http.Request) {
// // 	if r.Method == "GET" {
// // 		//ketika req sesuai lalu mendapat response, seharusnya http status http.StatusOK
// // 		http.Error(w, "Hello", http.StatusBadRequest)
// // 		return
// // 	}
// // 	http.Error(w, "", http.StatusBadRequest)
// // 	return
// // }

// // func main() {
// // 	http.HandleFunc("/suku", Suku)
// // 	// http.HandleFunc("/hello", hello)
// // 	fmt.Println("Starting web server at http://localhost:8080")
// // 	http.ListenAndServe(":8080", nil)
// // }

// //==============NO 3 =======================
// type Plannet struct{
// 	name string `json:"name"`
// 	rotation_period string `json:"height"`
// 	orbital_period string `json:"Mass"`
// 	diameter string `json:"hair_color"`
// 	climate string `json:"skin_color"`
// 	gravity string `json:"eye_color"`
// 	terrain string `json:"birth_year"`
// 	surface_water string `json:"gender"`
// 	population []string `json:"films"`
// }


// func main(){
// 	response, _ := http.Get("http://swapi.co/api/people/1")
// 	responseData, _ := ioutil.ReadAll(response.Body)
// 	defer response.Body.Close()

// 	var People Plannet
// 	fmt.Println("dt", string(responseData))
// 	json.Unmarshal(responseData, &People)
// 	fmt.Println("--------Print Field--------")
// 	fmt.Println("Name: ", People.Name)
// 	fmt.Println("height: ",People.Height)
// 	fmt.Println("Mass: ",People.Mass)
// 	fmt.Println("Hair Color: ",People.HairColor)
// 	fmt.Println("Film: ",People.Films[0])
// }