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
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestSignInHandler(t *testing.T) {
	cleanDB()
	w := httptest.NewRecorder()
	b := `{"member": {"fullname": "Baiju Muthukadan", "username": "baijum",
                "email": "baiju.m.mail@gmail.com", "password": "passwd"}}`
	r, _ := http.NewRequest("POST", "/api/v1/signup", strings.NewReader(b))
	GetRouter().ServeHTTP(w, r)

	w = httptest.NewRecorder()
	b = `{"member": {"username": "baijum",
                "password": "passwd"}}`
	r, _ = http.NewRequest("POST", "/api/v1/signin", strings.NewReader(b))
	GetRouter().ServeHTTP(w, r)
	closeDB()
}

func TestSignUpHandler(t *testing.T) {
	cleanDB()
	w := httptest.NewRecorder()
	b := `{"member": {"fullname": "Baiju Muthukadan", "username": "baijum",
                "email": "baiju.m.mail@gmail.com", "password": "passwd"}}`
	r, _ := http.NewRequest("POST", "/api/v1/signup", strings.NewReader(b))
	GetRouter().ServeHTTP(w, r)
	closeDB()
}

func TestGetRouter(t *testing.T) {
	r := GetRouter()
	router := mux.NewRouter()
	if reflect.TypeOf(r) != reflect.TypeOf(router) {
		t.Error("Router type is not matching *mux.Router")
	}
}

func TestOpenDB(t *testing.T) {
	cleanDB()
	if err := DB.Ping(); err != nil {
		t.Error("DB not connecting")
	}
	closeDB()
}

func cleanDB() {
	OpenDB()
	DB.Query(`DELETE FROM "doc"`)
	DB.Query(`DELETE FROM "member"`)
}

func closeDB() {
	DB.Close()
}
