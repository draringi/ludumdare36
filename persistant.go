package main

import (
	"compress/lzw"
	"encoding/gob"
	"os"
)

var PersistantData = new(struct {
	GameSettings *Settings
	SavedGames   [10]*gamestate
})

const persistantDataFile = ".traildata.persist"
const persistantDataOrder = lzw.LSB

func savePersistantToDisk() error {
	home := os.Getenv("HOME")
	err := os.Chdir(home)
	if err != nil {
		return err
	}
	os.Remove(persistantDataFile)
	f, err := os.Create(persistantDataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	compressedStream := lzw.NewWriter(f, persistantDataOrder, 8)
	defer compressedStream.Close()
	enc := gob.NewEncoder(compressedStream)
	return enc.Encode(PersistantData)
}

func loadPersistantFromDisk() error {
	home := os.Getenv("HOME")
	err := os.Chdir(home)
	if err != nil {
		return err
	}
	f, err := os.Open(persistantDataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	compressedStream := lzw.NewReader(f, persistantDataOrder, 8)
	defer compressedStream.Close()
	dec := gob.NewDecoder(compressedStream)
	return dec.Decode(&PersistantData)
}
