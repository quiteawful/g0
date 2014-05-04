// apitest project apitest.go
package Api

import (
	"errors"
	"fmt"
	"g0/db"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strconv"
)

type IDTest struct {
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
	var imgreturn []Image

	offset, _ := strconv.Atoi(r.PathParam("offset"))
	count, _ := strconv.Atoi(r.PathParam("count"))
	if offset == 42 {
		rest.NotFound(w, r)
		return
	}
	dbase, _ := db.NewDb("g0.db")
	dbarray, err := dbase.GetImages(offset, count)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, ele := range dbarray {
		var tmpImage Image
		tmpImage.ID = strconv.Itoa(ele.Id)
		tmpImage.Img = ele.Url
		tmpImage.Thumb = ele.Thumbnail
		imgreturn = append(imgreturn, tmpImage)
	}
	w.WriteJson(
		&IDTest{
			Images: imgreturn,
		})
}
