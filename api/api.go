// apitest project apitest.go
package Api

import (
	"errors"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type IDTest struct {
	Page     string  `json:"page"`
	Count    string  `json:"count"`
	ImageSrc string  `json:"image-src"`
	ThumbSrc string  `json:"thumb-src"`
	Images   []Image `json:"images"`
}
type Image struct {
	ID    string `json:"id"`
	Img   string `json:"img"`
	Thumb string `json:"thumb"`
}

type Api struct {
	Addr string
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
		&rest.Route{"GET", "/api/:offset/:count", GetIDstuff},
		&rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(handler.GetStatus())
			},
		},
	)
	http.ListenAndServe(a.Addr, &handler)
	return nil
}
func GetIDstuff(w rest.ResponseWriter, r *rest.Request) {
	offset := r.PathParam("offset")
	if offset == "42" {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(
		&IDTest{
			Page:   offset,
			Images: []Image{Image{"1", "foo", "bar"}},
		})
}
