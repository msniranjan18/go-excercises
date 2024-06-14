package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := deck{}
	if len(d) != 0 {
		t.Errorf("Expected deck lenth 52 but got %v", len(d))
	}

	d = newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck lenth 52 but got %v", len(d))
	}
}

func TestSaveToFileAndLoadDeckFromFile(t *testing.T) {
	testFilename := "_decktestfile"
	//Remove the existing file
	os.Remove(testFilename)
	d := newDeck()
	d.saveToFile(testFilename)
	loadedDeck, err := loadDeckFromFile(testFilename)
	if err != nil {
		if len(loadedDeck) != 52 {
			t.Errorf("Expected deck lenth 52 but got %v", len(loadedDeck))
		}
	}
	//Remove the file
	os.Remove(testFilename)
}
