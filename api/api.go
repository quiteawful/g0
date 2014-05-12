// apitest project apitest.go
package Api

import (
	"errors"
	"github.com/aimless/g0/conf"
	"github.com/aimless/g0/db"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
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
	Date  int64  `json:"date"`
	Nick  string `json:"user"`
	Chan  string `json:"channel"`
	Link  string `json:"source"`
}

type Api struct {
	Addr string
}

var (
	_api *Api = nil // singleton api holder
)

func init() {
	if _api == nil {
		_api = new(Api)
	}
	// conf foo here
	tmpApi := new(Api)
	conf.Fill(tmpApi)
	_api.Addr = tmpApi.Addr

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
		&rest.Route{"GET", "/api/:imgid/:count", GetIDstuff},
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
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var imgreturn []Image

	imgid, err := strconv.Atoi(r.PathParam("imgid"))
	if err != nil {
		rest.Error(w, "NYAN not found", 405)
		return
	}
	count, err := strconv.Atoi(r.PathParam("count"))
	if err != nil {
		rest.Error(w, "NYAN not found", 405)
		return
	}
	dbase, err := Db.NewDb()
	if err != nil {
		log.Printf("Api.GetIDstuff: %s\n", err.Error())
	}
	dbarray, err := dbase.GetLatestImages(imgid, count)

	if err != nil {
		log.Println(err.Error())
	}
	for _, ele := range dbarray {
		var tmpImage Image
		tmpImage.ID = strconv.Itoa(ele.Id)
		tmpImage.Img = ele.Name
		tmpImage.Thumb = ele.Thumbnail
		tmpImage.Date = ele.Timestamp.Unix()
		tmpImage.Nick = ele.User
		tmpImage.Chan = ele.Channel
		tmpImage.Link = ele.Url
		imgreturn = append(imgreturn, tmpImage)
	}
	w.WriteJson(
		&IDTest{
			ImageSrc: "http://aidskrebs.net/images/",
			ThumbSrc: "http://aidskrebs.net/images/",
			Images:   imgreturn,
		})
}
