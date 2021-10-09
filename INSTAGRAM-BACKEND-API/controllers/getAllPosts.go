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
	return &UserpostsController{s3} //RETURNS ADDRESS OF USER CONTROLLER
}

func (upc UserpostsController) GetAllPosts(w3 http.ResponseWriter, r3 *http.Request, p3 routing.Params) {
	id := p3.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w3.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	po := models.Post{}
	u := models.User{}

	if err := upc.session.DB("mongo-golang").C("posts").FindId(oid).One(&po); err != nil {
		w3.WriteHeader(404)
		return
	}

	count, err := upc.session.DB("mongo-golang").C("posts").Count()

	if err != nil {
		fmt.Print(err)
	}

	// var items []string

	for i := 0; i < count; i++ {
		items := upc.session.DB("mongo-golang").C("posts").Find(po.UserId == u.Id)
		pj, err := json.Marshal(items)

		if err != nil {
			fmt.Print(err)
		}

		w3.Header().Set("Content-Type", "application/json")
		w3.WriteHeader(http.StatusOK)
		fmt.Fprintf(w3, "%s\n", pj)
	}

}

// func (upc UserpostsController) Pagination(page int) ([]map[string]interface{}) {

// 	const (
// 		itemsPerPage = 10
// 	)

// 	var data []map[string]interface{}

// 	count, err := upc.session.DB("mongo-golang").C("posts").Count()

// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	start := (page - 1) * itemsPerPage
// 	stop := start + itemsPerPage

// 	if start > count {
// 		return nil
// 	}

// 	if stop > len(data) {
// 		stop = len(data)
// 	}
// }
