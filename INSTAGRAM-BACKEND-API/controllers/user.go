package controllers

//routing -- USED FOR PARAMS AND IN MAIN FUNCTION WITH NEW() ATTRIBUTE

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
	return &UserController{s1} //RETURNS ADDRESS OF USER CONTROLLER
}

//SAME STRUCT METHOD BUT EMPTY FOR PARAMS AS NOT USED HERE -- w is the request that is sent back to user and r takes request and params not needed as post method

func (uc UserController) CreateUser(w1 http.ResponseWriter, r1 *http.Request, _ routing.Params) {
	u := models.User{} //Taking user struct created on our own

	json.NewDecoder(r1.Body).Decode(&u) //decoding body coz it was in json -- so we get data in form of function

	u.Id = bson.NewObjectId() //creating id -- random object id

	pass := u.Password
	u.Password = string(encrypt([]byte(pass), "password"))

	uc.session.DB("mongo-golang").C("users").Insert(u) //taking session.db with help of mongodb and db is called mongo-goland and inserting user into it.

	uj, err := json.Marshal(u) //marshalling it so that we can send it back -- converting back to json

	if err != nil { // if marshalling returns error
		fmt.Println(err)
	}

	w1.Header().Set("Content-Type", "application/json") //if no error -- it returns uj we send it to frontend or postman
	w1.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w1, "%s\n", uj)

}

//GETUSER -- STRUCT METHOD WHICH TAKES 3 THINGS

func (uc UserController) GetUser(w1 http.ResponseWriter, r1 *http.Request, p1 routing.Params) {
	id := p1.ByName("id")

	if !bson.IsObjectIdHex(id) { //IF THAT ID IS HEX OR NOT CHECK
		w1.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id) //IF HEX THEN WE TAKE THIS OID AND USE IT IN MONGODB TO FIND DATA

	u := models.User{} //STRUCT OF TYPE MODELS.USER IN WHICH THE VALUES THAT WE ARE FINDING WILL COME

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w1.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u) //MARSHALLING - CONVERTING TO JSON

	if err != nil {
		fmt.Print(err)
	}

	w1.Header().Set("Content-Type", "application/json") //IF EVERYTHING FINE -- SET HEADER
	w1.WriteHeader(http.StatusOK)
	fmt.Fprintf(w1, "%s\n", uj)
}

// here we need params coz we need id to delete that particular data with that id

func (uc UserController) DeleteUser(w1 http.ResponseWriter, r1 *http.Request, p1 routing.Params) {
	id := p1.ByName("userId")

	if !bson.IsObjectIdHex(id) {
		w1.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w1.WriteHeader(404)
	}

	w1.WriteHeader(http.StatusOK)
	fmt.Fprint(w1, "Deleted user", oid, "\n")
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

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

// func decrypt(data []byte, passphrase string) []byte {
// 	key := []byte(createHash(passphrase))
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonceSize := gcm.NonceSize()
// 	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return plaintext
// }
