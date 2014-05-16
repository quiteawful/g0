package util

import (
	"fmt"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	src, err := DownloadPage("http://4fuckr.com/robots.txt")
	fmt.Printf("%s\n", src)
	if err != nil {
		t.Fatal(err)
	}
	if src != "User-agent: *\nDisallow:" {
		fmt.Printf("Got: %s\n", src)
	}
}

func TestImgurGetImagesFromGallery(t *testing.T) {
	var id string = "1sXiY" // 1sXiY zH9qXSg
	fmt.Println("Test: " + id)
	links, err := ImgurGetImagesFromGallery(id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Found: %v links.\n", len(links))
	for _, s := range links {
		fmt.Println(s)
	}

}

func TestImgurGetImagesFromAlbum(t *testing.T) {
	var id string = "MjR2s" // 1sXiY zH9qXSg
	fmt.Println("Test: " + id)
	links, err := ImgurGetImagesFromAlbum(id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Found: %v links.\n", len(links))
	for _, s := range links {
		fmt.Println(s)
	}
}
