package main

import (
// "encoding/json"
"net/http"
// "fmt"
// "io/ioutil"
// "strconv"
"github.com/labstack/echo"
)

// func main(){
// 	e := echo.New()

// 	e.GET("/", HelloController)
// 	e.Start(":8000")
// }

// func HelloController(c echo.Context) error{
// 	return c.String(http.StatusOK, "Hello World")
// }

//========2

// type User struct {
// 	Name string
// 	Email string
// }
// func GetUser(c echo.Context) error{
// 	user := User {Name: "Ainin", Email: "ainin@alterra.id"}
// 	return c.JSON (http.StatusOK, user)
// 	}

// func main(){
// 	e := echo.New()
// 	e.GET("/user", GetUser)
// 	e.Start(":8000")
// }

//======3 [Retrive Data]

// type User struct{
// 	Id		int
// 	Name	string
// 	Email	string
// }

// //User controller
// func GetUserController(c echo.Context) error{
// 	id, _:= strconv.Atoi(c.Param("id")) //ngambil param yg ada di -> id
// 	user := User{Id: id, Name: "Ainin", Email: "ainin@alterra.id"}
// 	//Render Data - JSON Response
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"user": user,
// 	})
// }

// func main(){
// 	e := echo.New()
// 	e.GET("/users/:id", GetUserController)
// 	//routing with routing parameter
// 	e.Start(":8000")
// }


//========4 query param

// func UserSeacrhController(c echo.Context) error{
// 	//get data 
// 	match := c.QueryParam("match")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"match": match,
// 		"result": []string{"ainin", "nur", "asiyah"},
// 	})
// }

// func main(){
// 	e := echo.New()
// 	e.GET("/users", UserSeacrhController)
// 	//routing with routing parameter
// 	e.Start(":8000")
// }

//=====Challenge=====


type User struct {
	Id			int	 `json:"id", form:"id"`
	Name		string `json:"name", form:"name"`
	Email		string `json:"email", form:"email"`
	Password	string `json:"password", form:"password"`
}

var users []User
//---------------controller-------------------
func GetUserController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": users,
	})
}


//------------------------------------------------------
// func GetUserController(c echo.Context) error{
// 	//your solution
// }

// func DeleteUserController(c echo.Context) error{
// 	//your solution
// }

// func UpdateUserController(c echo.Context) error{
// 	//your solution
// }


func CreateUserController(c echo.Context) error{
	// binding data

	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
users = append(users, user)
return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "success create user",
	"user": user,
	})
}

func main(){
	e := echo.New()
	//routing with query parameter
	e.GET("/users", GetUserController)
	e.POST("/users", CreateUserController)
}