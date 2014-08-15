// util project util.go
package util

import (
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"

	"github.com/quiteawful/g0/conf"
)

var (
	_util               *ConfImg = nil
	imgurregex                   = regexp.MustCompile("(http://)?imgur.com/gallery/[A-Za-z0-9]*")
	imgurSubredditRegex          = regexp.MustCompile("(http://)?imgur.com/r/[A-Za-z0-9]*/[A-Za-z0-9]*")
	imguralbumregex              = regexp.MustCompile("(http://)?imgur.com/a/[A-Za-z0-9]*")
	idregex                      = regexp.MustCompile("[A-Za-z0-9]*")
)

type ConfImg struct {
	Imagepath string
	Thumbpath string
}

func init() {
	if _util == nil {
		_util = new(ConfImg)
	}
	tmpConf := new(ConfImg)
	conf.Fill(tmpConf)

	_util.Imagepath = tmpConf.Imagepath
	_util.Thumbpath = tmpConf.Thumbpath
}

const MAX_SIZE = 10485760

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var imageregex = regexp.MustCompile(`image\/(.+)|video\/webm`)

func DownloadImage(link string) (filename, hash string, errret error) {
	_, err := url.Parse(link)
	if err != nil {
		return "", "", err
	}
	var (
		bufa    [64]byte
		b       []byte
		urlType []string
		mime    string
	)

	size := 0
	buf := bufa[:]
	/* extract image links from imgur */
	if imgurregex.MatchString(link) {
		arr := idregex.FindAllString(link, -1)
		id := arr[len(arr)-1]
		galleryUrlString, err := ImgurGetImagesFromGallery(id)
		if err != nil {
			log.Printf("Util parse Imgur: %s\n", err.Error())
			return "", "", err
		}
		link = galleryUrlString[0]
	}

	if imguralbumregex.MatchString(link) {
		arr := idregex.FindAllString(link, -1)
		id := arr[len(arr)-1]
		albumUrlString, err := ImgurGetImagesFromAlbum(id)
		if err != nil {
			log.Printf("util.DownloadImage: Could not fetch images from imgur album. %s\n", err.Error())
			return "", "", err
		}
		link = albumUrlString[0]
	}

	if imgurSubredditRegex.MatchString(link) {
		arr := idregex.FindAllString(link, -1)
		var id string
		if len(arr) == 8 {
			id = arr[7]
		} else {
			return "", "", err
		}

		subredditlink, err := ImgurGetImageFromSubreddit(id)
		if err != nil {
			log.Printf("util.DownloadImage: Could not fetch image url from subreddit. %s\n", err.Error())
			return "", "", err
		}
		link = subredditlink
	}

	res, err := http.Get(link)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	if err != nil {
		return "", "", err
	}
	for {
		n, err := res.Body.Read(buf)
		if size == 0 {
			mime = http.DetectContentType(buf)
			urlType = imageregex.FindStringSubmatch(mime)
			if urlType == nil {
				return "", "", fmt.Errorf("not an image: %q", mime)
			}
		}
		size += n
		if size > MAX_SIZE {
			return "", "", fmt.Errorf("image too large")
		}
		b = append(b, buf[:n]...)
		if err == io.EOF {
			h := md5.New()
			h.Write(b)
			if urlType[1] == "" {
				urlType[1] = "webm"
			}
			filename = newLenChars(6, StdChars) + "." + urlType[1]
			ioutil.WriteFile(_util.Imagepath+filename, b, 0644)
			if mime == "video/webm" {
				out, err := exec.Command("ffmpeg", "-y", "-i", _util.Imagepath+filename, "-ss", "2", "-vframes", "1", _util.Imagepath+"tmp.jpeg").CombinedOutput()
				if err != nil {
					fmt.Println(err.Error() + string(out))
				}
			}
			return filename, fmt.Sprintf("%x", h.Sum(nil)), nil
		}
	}
	return filename, "", nil
}

// NewLenChars stolen from https://github.com/dchest/uniuri , thx
func newLenChars(length int, chars []byte) string {
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, r); err != nil {
			panic("error reading from random source: " + err.Error())
		}
		for _, c := range r {
			if c >= maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
	panic("unreachable")
}

func IsDirWriteable(path string) bool {
	// can be used for setup/startup to check wether we can write to imagepath.
	return false
}

func DownloadPage(link string) (string, error) {
	if link == "" {
		err := errors.New("Empty url.")
		log.Printf("Util.DownloadPage: %s\n", err.Error())
		return "", err
	}

	resp, err := http.Get(link)
	if err != nil {
		log.Printf("Util.DownloadPage: %s\n", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Util.DownloadPage: %s\n", err.Error())
		return "", err
	}

	return string(body), nil
}
