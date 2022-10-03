package server

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed wisdoms.txt
var wisdomsSource []byte

type Wisdom struct {
	Text   string
	Author string
}

type Db [][]byte

func (db Db) getWisdom() (*Wisdom, error) {
	wisdomRaw := db[rand.Intn(len(db))]
	comps := strings.Split(string(wisdomRaw), "~")
	if len(comps) < 2 {
		return nil, fmt.Errorf("bad wisdom")
	}

	return &Wisdom{comps[0], comps[1]}, nil
}

func newDb() Db {
	return bytes.Split(wisdomsSource, []byte("\n"))
}

func init() {
	rand.Seed(time.Now().Unix())
}
