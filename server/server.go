// Copyright (C) 2015 Baiju Muthukadan <baiju.m.mail@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package server

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/scrypt"
)

type Member struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TOC      bool   `json:"toc"`
	Verified bool   `json:"verified"`
	Active   bool   `json:"verified"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var member map[string]Member

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&member)
	u := member["member"]
	var salt, password []byte
	err := DB.QueryRow(`SELECT salt, password FROM "member"
	        WHERE (username = $1 OR email = $2) AND verified = true AND active = true`,
		u.Username, u.Email).Scan(&salt, &password)
	if err != nil {
		log.Println(err)
	}
	fs := append(salt, Pepper...)
	log.Printf("%v", fs)
	dk, _ := scrypt.Key([]byte(u.Password), fs, 16384, 8, 1, 32)
	if bytes.Equal(password, dk) {
		log.Printf("Correct password")
	} else {
		log.Printf("Wrong password")
	}
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	var member map[string]Member
	c := 64
	salt := make([]byte, c)
	_, err := rand.Read(salt)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&member)
	u := member["member"]
	var id int
	fs := append(salt, Pepper...)
	log.Printf("%v", fs)
	dk, _ := scrypt.Key([]byte(u.Password), fs, 16384, 8, 1, 32)
	err = DB.QueryRow(`INSERT INTO "member" (
	        fullname, username, email, salt, password) VALUES ($1, $2, $3, $4, $5)
                RETURNING id`, u.Fullname, u.Username, u.Email, salt, string(dk)).Scan(&id)
	if err != nil {
		log.Println(err)
	}

}

// Router for the URL endpoints
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/signin", signInHandler).Methods("POST")
	r.HandleFunc("/api/v1/signup", signUpHandler).Methods("POST")
	return r
}

// EmberClassic returns a new Negroni instance with the default
// middleware already in the stack.
//
// Recovery - Panic Recovery Middleware
// Logger - Request/Response Logging
// Static - Static File Serving
func EmberClassic(dir string) *negroni.Negroni {
	return negroni.New(negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(dir)))
}

var (
	DB     *sql.DB
	Pepper string
)

func OpenDB() {
	var err error
	u := os.Getenv("MDOCS_POSTGRES_USER")
	d := os.Getenv("MDOCS_POSTGRES_DBNAME")
	p := os.Getenv("MDOCS_POSTGRES_PASSWORD")
	cs := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", u, d, p)
	DB, err = sql.Open("postgres", cs)

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	Pepper = os.Getenv("MDOCS_RANDOM_SALT")
}
