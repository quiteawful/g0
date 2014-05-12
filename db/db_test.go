package Db

import (
	"testing"
)

func TestNewDb(t *testing.T) {

	err := Open()
	if err != nil {
		t.Fatal(err)
	}
	Close()
}

func TestImageSetup(t *testing.T) {
	_, err := ImageInit()
	if err != nil {
		t.Fatal(err)
	}

}

/*
func TestQuery(t *testing.T) {

	err := Open()
	if err != nil {
		t.Fatal(err)
	}
	_, err = Query("select 1;")
	if err != nil {
		t.Fatal(err)
	}

	Close()
}


func TestNewImage(t *testing.T) {
	testDb, err := NewDb("newImage.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}
}

func TestGetImage(t *testing.T) {
	testDb, err := NewDb("getImage.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	_, err = testDb.GetImage(1)
	if err != nil {
		t.Fatal("Failed to select image")
	}
}

func TestGetImages(t *testing.T) {
	testDb, err := NewDb("newImages.db")
	if err != nil {
		t.Fatal(err)
	}

	// insert paar bilder
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	img, err := testDb.GetImages(1, 5)
	if err != nil {
		t.Fatal("Failed to select image")
	}
	if len(img) != 5 {
		t.Fatal("Requested 5 Images, got %d", len(img))
	}
}

func TestDeleteImage(t *testing.T) {
	os.Remove("deleteImage.db")
	testDb, err := NewDb("deleteImage.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	if !testDb.DeleteImage(1) {
		t.Fatal("Failed to delete image.")
	}
}

func TestGetImageCount(t *testing.T) {
	os.Remove("getImageCount.db")
	testDb, err := NewDb("getImageCount.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	c, err := testDb.GetImageCount()
	if err != nil {
		t.Fatal("Failed to count images in db.")
	}
	if c != 1 {
		t.Fatalf("Failed to count images correctly, got %d expected 1\n", c)
	}
}
*/
