package mp3player

import (
	"errors"
)

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (*MusicEntry, int) {
	if len(m.musics) == 0 {
		return nil, -1
	}

	for index, v := range m.musics {
		if v.Name == name {
			return &v, index
		}
	}
	return nil, -1
}

func (m *MusicManager) FindIndexByName(name string) int {
	for k, v := range m.musics {
		if v.Name == name {
			return k
		}
	}
	return -1
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	//从切片中删除元素
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}
