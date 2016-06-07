package main

import (
    "fmt"
    "net/http"
    "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)


type Account struct {
	email string
	password string
}

func CreateAccount(account *Account) acc1 {
	// Dial up a mongoDB session
	session, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Opens the "data" database, "accounts" collection
	c := session.DB("data").C("accounts")
	result := Account{}

	// Search the "library" databases, "accounts" collection

	err = c.Find(bson.M{"email": account.email}).One(&result)

	if result.email != "" {
		// return true because account is present in the database
		// and we can say, "it's been added" without causing errors
		return true
	}

	err = c.Insert(*account)

	if err != nil {
		return false
	}

	// Close session to save resources
	session.close()
	return true
}




func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func data_base(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let me check if its working"))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/db", data_base)
    http.ListenAndServe(":8080", nil)
}