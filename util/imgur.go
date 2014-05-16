package util

import (
	"errors"
	"log"
	"regexp"
)

func ImgurGetImagesFromAlbum(id string) (string, error) {
	return "", nil
}

func ImgurGetImagesFromGallery(id string) ([]string, error) {
	var links []string
	if id == "" {
		err := errors.New("Empty id.")
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	src, err := DownloadPage("http://imgur.com/gallery/" + id)
	if err != nil {
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	// id:zH9qXSg
	// <img src="//i.imgur.com/zH9qXSg.jpg" alt="" />
	// http://imgur.com/gallery/1sXiY <- many
	//var urlregex = regexp.MustCompile(`<img src="//i.imgur.com/[a-zA-Z0-9]{5-7}.(jpg|jpeg|png|gif|apng|tiff|bmp)" alt="(.*)?" />`)
	var urlregex = regexp.MustCompile(`<img src="\/\/i\.imgur\.com\/(?P<id>(.*)\.(jpg|jpeg|png|gif|apng|tiff|bmp))`)

	if urlregex.MatchString(src) {
		raw := urlregex.FindAllStringSubmatch(src, -1)
		for _, n := range raw {
			links = append(links, n[1])
		}

	}
	return links, nil
}
