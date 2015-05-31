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
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type User struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TOC      bool   `json:"toc"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	var u map[string]User

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&u)
	log.Printf("%#v", u)
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
