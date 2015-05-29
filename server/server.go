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

	"github.com/gorilla/mux"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
}

// Router for the URL endpoints
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	return r
}
