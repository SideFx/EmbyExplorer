// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// Save/load preferences (Window size/position & Emby access data)
// ---------------------------------------------------------------------------------------------------------------------

package ui

import (
	"Emby_Explorer/api"
	"Emby_Explorer/assets"
	"Emby_Explorer/settings"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const preferencesFileName = "org.janbuchholz.embyexplorer.json"

func SavePreferences() error {
	s := settings.GetPreferences()
	p := s.EmbyPassword
	s.EmbyPassword = encode(p)
	j, err := json.Marshal(s)
	if err == nil {
		dir, _ := os.UserConfigDir()
		dir = filepath.Join(dir, assets.AppName)
		_, err := os.Stat(dir)
		if err != nil {
			if err := os.Mkdir(dir, os.ModePerm); err != nil {
				return err
			}
		}
		fname := filepath.Join(dir, preferencesFileName)
		err = os.WriteFile(fname, j, 0644)
	}
	if err == nil {
		s.EmbyPassword = p
		settings.SetPreferences(s)
	}
	return err
}

func LoadPreferences() error {
	var s settings.Settings
	dir, err := os.UserConfigDir()
	dir = filepath.Join(dir, assets.AppName)
	fname := filepath.Join(dir, preferencesFileName)
	j, err := os.Open(fname)
	if err == nil {
		byteValue, _ := io.ReadAll(j)
		_ = j.Close()
		err = json.Unmarshal(byteValue, &s)
	}
	if err == nil {
		s.EmbyPassword = decode(s.EmbyPassword)
		settings.SetPreferences(s)
		if s.EmbyServer != "" {
			api.CheckEmby(s.EmbyServer)
		}
	}
	return err
}

const bits = 8

func encode(b []byte) []byte {
	var c = make([]byte, len(b))
	copy(c, b)
	j := 0
	for i, e := range c {
		if i > bits-1 {
			i = 1
		}
		c[j] = ror(e, i)
		j++
	}
	return c
}

func decode(b []byte) []byte {
	var c = make([]byte, len(b))
	copy(c, b)
	j := 0
	for i, d := range c {
		if i > bits-1 {
			i = 1
		}
		c[j] = rol(d, i)
		j++
	}
	return c
}

func ror(x byte, n int) byte {
	return (x >> n) | (x << (bits - n))
}

func rol(x byte, n int) byte {
	return (x << n) | (x >> (bits - n))
}
