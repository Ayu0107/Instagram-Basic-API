//MAIN FILE

package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/ayushi0107/INSTAGRAM-BACKEND-API/controllers"
	"github.com/ayushi0107/INSTAGRAM-BACKEND-API/routing"
)

func main() {

	const database, collection = "mongo-golang", "user"

	//CREATES A NEW INSTANT

	r := routing.New()

	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())
	upc := controllers.NewUserpostsController(getSession())

	//FOR USER -- CREATE USER AND GETUSER
	r.POST("/user", uc.CreateUser)
	r.GET("/user/:id", uc.GetUser)

	//FOR POSTS -- CREATEPOST AND GETPOST
	r.POST("/posts", pc.CreatePost)
	r.GET("/posts/:id", pc.GetPost)

	//TO GET ALL THE POSTS
	r.GET("/posts/users/:id", upc.GetAllPosts)

	http.ListenAndServe("localhost:8080", r)
}

//CREATING SESSIONS

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
