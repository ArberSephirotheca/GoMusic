package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
)

type Song struct {
	tag.Metadata
	path string
}
const(
	mp3 Decode = iota
	wav 
	flac
	ogg
)

func (D Decode) String() string{
	switch D{
	case mp3:
		return ".mp3"
	case wav:
		return ".wav"
	case flac:
		return ".flac"
	case ogg:
		return ".ogg"
	default:
		return ""
	}
}
const defaultPath = "/Users/czy/Music/Music/Media.localized/Apple Music/Lube/Свои/"

func Load(path string) []*Song {
	songs := make([]*Song, 0, 0)
	files, err := os.ReadDir(defaultPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		freader, err := os.Open(path + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		meta, err := tag.ReadFrom(freader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		song := new(Song)
		song.Metadata = meta
		song.path = freader.Name()
		songs = append(songs, song)
	}

	return songs
}

func LoadDefault() []*Song {
	songs := make([]*Song, 0, 0)
	files, err := os.ReadDir(defaultPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		freader, err := os.Open(defaultPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		meta, err := tag.ReadFrom(freader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		song := new(Song)
		song.Metadata = meta
		song.path = freader.Name()
		songs = append(songs, song)
	}
	return songs
}

func 