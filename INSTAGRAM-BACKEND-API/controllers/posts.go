//CONTAINS FUNCTIONS FOR POSTS -- CREATE POST AND GET POST

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

type PostController struct {
	session *mgo.Session
}

func NewPostController(s2 *mgo.Session) *PostController {
	return &PostController{s2} //RETURNS ADDRESS OF Post CONTROLLER
}

//FUNCTION TO CREATE POST

func (pc PostController) CreatePost(w2 http.ResponseWriter, r2 *http.Request, _ routing.Params) {
	po := models.Post{}

	json.NewDecoder(r2.Body).Decode(&po)

	po.Id = bson.NewObjectId()

	pc.session.DB("mongo-golang").C("posts").Insert(po)

	pj, err := json.Marshal(po)

	if err != nil {
		fmt.Println(err)
	}

	w2.Header().Set("Content-Type", "application/json")
	w2.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w2, "%s\n", pj)
}

//FUNCTION TO GET POST

func (pc PostController) GetPost(w2 http.ResponseWriter, r2 *http.Request, p2 routing.Params) {
	id := p2.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w2.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	po := models.Post{}

	if err := pc.session.DB("mongo-golang").C("posts").FindId(oid).One(&po); err != nil {
		w2.WriteHeader(404)
		return
	}

	pj, err := json.Marshal(po)

	if err != nil {
		fmt.Print(err)
	}

	w2.Header().Set("Content-Type", "application/json")
	w2.WriteHeader(http.StatusOK)
	fmt.Fprintf(w2, "%s\n", pj)
}
