//CONTAINS FUNCTION FOR USERS -- CREATEUSER AND GETUSER AND ENCRYPTPASSWORD

package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ayushi0107/INSTAGRAM-BACKEND-API/models"
	"github.com/ayushi0107/INSTAGRAM-BACKEND-API/routing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s1 *mgo.Session) *UserController {
	return &UserController{s1}
}

//CREATE USER FUNCTION

func (uc UserController) CreateUser(w1 http.ResponseWriter, r1 *http.Request, _ routing.Params) {
	u := models.User{}

	json.NewDecoder(r1.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	//ENCRYPTING PASSWORD

	pass := u.Password
	u.Password = string(encrypt([]byte(pass), "password"))

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w1.Header().Set("Content-Type", "application/json")
	w1.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w1, "%s\n", uj)

}

//GET USER FUNCTION

func (uc UserController) GetUser(w1 http.ResponseWriter, r1 *http.Request, p1 routing.Params) {
	id := p1.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w1.WriteHeader(404)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w1.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Print(err)
	}

	w1.Header().Set("Content-Type", "application/json")
	w1.WriteHeader(201)
	fmt.Fprintf(w1, "%s\n", uj)
}

//TO ENCRYPT PASSWORD -- CREATING HASH

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//FUNCTION TO ENCRYPT PASSWORD

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
