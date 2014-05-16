package util

import (
	"errors"
	"log"
	"regexp"
)

func ImgurGetImagesFromAlbum(id string) ([]string, error) {
	var links []string
	var urlregex = regexp.MustCompile(`data\-src="\/\/i\.imgur\.com\/(?P<id>(.*)[^s]\.(jpg|jpeg|png|gif|apng|tiff|bmp))" `)

	if id == "" {
		err := errors.New("Empty id.")
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	src, err := DownloadPage("http://imgur.com/a/" + id)
	if err != nil {
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	if urlregex.MatchString(src) {
		raw := urlregex.FindAllStringSubmatch(src, -1)
		for _, n := range raw {
			links = append(links, "http://i.imgur.com/"+n[1])
		}
	}
	return links, nil
}

func ImgurGetImagesFromGallery(id string) ([]string, error) {
	var links []string
	var urlregex = regexp.MustCompile(`<img src="\/\/i\.imgur\.com\/(?P<id>(.*)\.(jpg|jpeg|png|gif|apng|tiff|bmp))`)

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

	if urlregex.MatchString(src) {
		raw := urlregex.FindAllStringSubmatch(src, -1)
		for _, n := range raw {
			links = append(links, "http://i.imgur.com/"+n[1])
		}
	}
	return links, nil
}
