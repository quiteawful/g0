//usage:
//	imgtest, _ := img.GetImageFromFile("audrey.jpg")
//	thmtest, _ := img.MakeThumbnail(imgtest, 150, 150)
//	img.SaveImageAsJPG("thumb-audjey.jpg", thmtest)

package img

import (
	"code.google.com/p/graphics-go/graphics"
	_ "code.google.com/p/vp8-go/webp"
	"fmt"
	"github.com/quiteawful/g0/conf"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

var (
	_img *ConfImg = nil // config holder for package instance
)

type ConfImg struct {
	Imagepath string
	Thumbpath string
}

func init() {
	if _img == nil {
		_img = new(ConfImg)
	}
	tmpConf := new(ConfImg)
	conf.Fill(tmpConf)

	_img.Imagepath = tmpConf.Imagepath
	_img.Thumbpath = tmpConf.Thumbpath
}

func GetImageFromFile(f string) (image.Image, error) {
	fi, ferr := os.Open(_img.Imagepath + f)
	if ferr != nil {
		fmt.Println(ferr.Error())
		return nil, ferr
	}
	pic, perr := GetImage(fi)
	if perr != nil {
		fmt.Println(perr.Error())
		return nil, perr
	}

	defer func() {
		fi.Close()
	}()

	return pic, nil

}

func GetImage(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return img, nil
}

func MakeThumbnail(src image.Image, x int, y int) (image.Image, error) {
	tgt := image.NewRGBA(image.Rect(0, 0, x, y))
	err := graphics.Thumbnail(tgt, src)
	return tgt, err
}

func SaveImageAsJPG(f string, src image.Image) error {
	fi, ferr := os.Create(_img.Thumbpath + f) // thumbpath. richtig? oder imgpath?
	if ferr != nil {
		fmt.Println(ferr.Error())
		return ferr
	}
	defer func() {
		fi.Chmod(0644)
		fi.Close()
	}()
	return jpeg.Encode(fi, src, nil)

}
