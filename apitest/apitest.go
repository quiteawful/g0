// apitest project apitest.go
package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type IDTest struct {
	Page   string  `json:"page"`
	Images []Image `json:"images"`
}
type Image struct {
	Img   string `json:"img"`
	Thumb string `json:"thumb"`
}

func main() {
	fmt.Println("test")
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		EnableStatusService:      true,
		XPoweredBy:               "soda-api",
	}
	handler.SetRoutes(
		&rest.Route{"GET", "/api/:id", GetIDstuff},
		&rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(handler.GetStatus())
			},
		},
	)
	http.ListenAndServe(":31337", &handler)
}

func GetIDstuff(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	if id == "42" {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(
		&IDTest{
			Page:   id,
			Images: []Image{Image{"foo", "bar"}},
		})
}
