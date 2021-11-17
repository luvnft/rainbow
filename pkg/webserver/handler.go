// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teal-finance/garcon/reserr"
	"github.com/teal-finance/garcon/webserver"
)

// Handler to server the static web files of the front-end.
func Handler(resErr reserr.ResErr, wwwDir string) http.Handler {
	r := chi.NewRouter()

	// Static website files
	fs := webserver.WebServer{Dir: wwwDir, ResErr: resErr}
	r.NotFound(fs.ServeFile("index.html", "text/html; charset=utf-8")) // catch index.html and other Vue sub-folders
	r.Get("/favicon.ico", fs.ServeFile("favicon.ico", "image/x-icon"))
	r.Get("/favicon.png", fs.ServeFile("favicon.png", "image/png"))
	r.Get("/preview.jpg", fs.ServeFile("preview.jpg", "image/jpeg"))
	r.Get("/assets/js/*", fs.ServeDir("text/javascript; charset=utf-8"))
	r.Get("/assets/css/*", fs.ServeDir("text/css; charset=utf-8"))

	return r
}
