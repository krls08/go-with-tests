package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// func NewFileSystemPlayerStore(aDatabase io.ReadWriteSeeker) *FileSystemPlayerStore {
func NewFileSystemPlayerStore(aFile *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(aFile)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	aLeague, err := NewLeague(aFile)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", aFile.Name(), err)
	}
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{aFile}),
		league:   aLeague,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetSortedLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)

}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info form file %s", file.Name())
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil

}
