package db

import (
	"fmt"
	"os"
	"testing"
)

var (
	dbfile string = "test.db"
)

func TestNewDb(t *testing.T) {
	defer os.Remove(dbfile)
	testDb, err := NewDb(dbfile)
	if err != nil {
		t.Fatal(err)
	}
	defer testDb.Close()

	if _, err := os.Stat(testDb.DbFile); os.IsNotExist(err) {
		t.Fatal("Failed to create database.")
	}
}

func TestNewImage(t *testing.T) {
	defer os.Remove(dbfile)
	testDb, err := NewDb(dbfile)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}
}

func TestGetImage(t *testing.T) {
	defer os.Remove(dbfile)
	testDb, err := NewDb(dbfile)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	img, err := testDb.GetImage(1)
	if err != nil {
		t.Fatal("Failed to select image")
	}
	fmt.Printf("ID: %s\n", img.id)
	fmt.Printf("Hash: %s\n", img.hash)
	fmt.Printf("Name: %s\n", img.name)
	fmt.Printf("Thumbnail: %s\n", img.thumbnail)
	fmt.Printf("Timestamp: %s\n", img.timestamp.String())
	fmt.Printf("Url: %s\n", img.url)
	fmt.Printf("Network: %s\n", img.network)
	fmt.Printf("Channel: %s\n", img.channel)
	fmt.Printf("User: %s\n", img.user)
}

func TestGetImages(t *testing.T) {
	//defer os.Remove(dbfile)
	testDb, err := NewDb(dbfile)
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
