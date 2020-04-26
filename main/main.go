package main

import (
	"fmt"
	"github.com/MediaExchange/log"
	"github.com/MediaExchange/torg"
	"github.com/MediaExchange/torg/torrent"
)

func main() {
	log.Info("Starting")
	res := torg.Search("The Avengers 2012 1080p", torrent.MoviesHD)
	fmt.Printf("%+v\n", res)
}
