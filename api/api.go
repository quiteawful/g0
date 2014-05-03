// apitest project apitest.go
package Api

import (
	"errors"
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

type Api struct {
	addr string
}

func NewApi(addr string) (*Api, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	return &Api{addr}, nil
}
func (a *Api) Run() (err error) {
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
	http.ListenAndServe(a.addr, &handler)
	return nil
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
