//CONTAINS FUNCTION TO GET ALL POSTS OF A USER

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayushi0107/INSTAGRAM-BACKEND-API/models"
	"github.com/ayushi0107/INSTAGRAM-BACKEND-API/routing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserpostsController struct {
	session *mgo.Session
}

func NewUserpostsController(s3 *mgo.Session) *UserpostsController {
	return &UserpostsController{s3}
}

//FUNCTION TO POST ALL POSTS OF A USER -- NOT GETTING COMPLETE OUTPUT

func (upc UserpostsController) GetAllPosts(w http.ResponseWriter, r *http.Request, p routing.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
	}

	po := models.Post{}
	u := models.User{}

	count, err := upc.session.DB("mongo-golang").C("posts").Count()

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < count; i++ {
		if upc.session.DB("mongo-golang").C("posts").Find(po.UserId == u.Id) != nil {

			pj, err := json.Marshal(po)

			if err != nil {
				fmt.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			fmt.Fprintf(w, "%s\n", pj)
		}
	}
}
