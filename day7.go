package main

// import (
// "encoding/json"
// "net/http"
// "fmt"
// "io/ioutil"
// )


// type article struct{
// 	ID		int
// 	Title	string
// 	Content	string
// }

// type starsWarsPeople struct{
// 	Name string `json:"name"`
// 	Height string `json:"height"`
// 	Mass string `json:"Mass"`
// 	HairColor string `json:"hair_color"`
// 	SkinColor string `json:"skin_color"`
// 	EyeColor string `json:"eye_color"`
// 	BirthYear string `json:"birth_year"`
// 	Gender string `json:"gender"`
// 	Films []string `json:"films"`
// }


// func main(){
// 	response, _ := http.Get("http://swapi.co/api/people/1")
// 	responseData, _ := ioutil.ReadAll(response.Body)
// 	defer response.Body.Close()

// 	var People starsWarsPeople
// 	fmt.Println("dt", string(responseData))
// 	json.Unmarshal(responseData, &People)
// 	fmt.Println("--------Print Field--------")
// 	fmt.Println("Name: ", People.Name)
// 	fmt.Println("height: ",People.Height)
// 	fmt.Println("Mass: ",People.Mass)
// 	fmt.Println("Hair Color: ",People.HairColor)
// 	fmt.Println("Film: ",People.Films[0])
// }

// func articles(w http.ResponseWriter, r *http.Request){
// 	var data = []article{
// 		article{1, "Judul pertama", "isi pertama"},
// 		article{2, "Judul kedua", "isi kedua"},
// 	}
	
// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Method == "GET" {
// 		result, err := json.Marshal(data)
// 		if err != nil{
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(result)
// 		return
// 	}
// http.Error(w, "", http.StatusBadRequest)
// return
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		//ketika req sesuai lalu mendapat response, seharusnya http status http.StatusOK
// 		http.Error(w, "Hello", http.StatusBadRequest)
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// 	return
// }

// func main() {
// 	http.HandleFunc("/articles", articles)
// 	http.HandleFunc("/hello", hello)
// 	fmt.Println("Starting web server at http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }