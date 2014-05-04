// util project util.go
package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const MAX_SIZE = 1048576

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func DownloadImage(link string) (filename string, errret error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	fmt.Println(u.Fragment)
	var bufa [64]byte
	var b []byte
	size := 0
	buf := bufa[:]
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if err != nil {
		fmt.Println("asdf")
		return "", err
	}
	for {
		n, err := res.Body.Read(buf)
		if size == 0 {
			mime := http.DetectContentType(buf)
			if strings.Contains(mime, "image") == false {
				return "", fmt.Errorf("not an image: %q", mime)
			}
		}
		size += n
		if size > MAX_SIZE {
			return "", fmt.Errorf("image too large")
		}
		b = append(b, buf[:n]...)
		if err == io.EOF {
			ioutil.WriteFile(newLenChars(6, StdChars)+".jpg", b, 0644)
			return filename, nil
		}
	}
	return filename, nil
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
