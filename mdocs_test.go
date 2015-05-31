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

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/baijum/mdocs/server"
)

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	router := server.GetRouter()
	n := server.EmberClassic("web/dist")
	r, _ := http.NewRequest("GET", "/", nil)
	n.UseHandler(router)
	n.ServeHTTP(w, r)
	if !strings.Contains(w.Body.String(), "<title>Markdown Docs</title>") {
		t.Error("Home page has wrong title")
	}
}
