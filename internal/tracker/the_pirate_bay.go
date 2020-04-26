/*
   Copyright 2020 MediaExchange.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package tracker

import (
	"fmt"
	"github.com/MediaExchange/log"
	"github.com/MediaExchange/torg/torrent"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

type ThePirateBay struct { }

func (t *ThePirateBay) Search(str string, category torrent.Category) []torrent.Result {
	page := 0           // Do we want multiple pages, or is the first one enough?
	orderBy := 99       // Not sure what this orders the results by.
	c := t.convertCategory(category)
	searchUrl := fmt.Sprintf("https://thepiratebay.org/search/%s/%d/%d/%d", url.QueryEscape(str), page, orderBy, c)

	results := make([]torrent.Result, 0)
	doc, err := getPage(searchUrl)
	if err != nil {
		log.Error("Failed to retrieve page", log.String("host", "thepiratebay.org"), log.Err(err))
		return results
	}

	doc.Find("#SearchResults #content #main-content #searchResult tbody tr").Each(func(i int, s *goquery.Selection) {
		var r torrent.Result

		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 1:
				s.Find("a").Each(func(i int, s *goquery.Selection) {
					href, exists := s.Attr("href")
					if exists && strings.HasPrefix(href, "magnet:") {
						r.Magnet, err = torrent.MagnetFromLink(href)
						if err != nil {
							log.Error("Failed to parse magnet link", log.Err(err))
						}
					}
				})
			case 2:
				r.Seeders, _ = strconv.Atoi(s.Text())
			}
		})

		results = append(results, r)
	})

	return results
}

func (t *ThePirateBay) convertCategory(c torrent.Category) int {
	switch c {
	// Audio mappings
	case torrent.Audio:
		return 100
	case torrent.AudioMP3:
		return 101
	case torrent.AudioAudioBook:
		return 102
	case torrent.AudioLossless:
		return 104
	case torrent.AudioOther:
		return 199

	// Movie mappings
	case torrent.Movies:
		return 201
	case torrent.MoviesDvd:
		return 202
	case torrent.MoviesBluray, torrent.MoviesHD, torrent.MoviesWebDl:
		return 207

	// TV show mappings
	case torrent.Tv:
		return 205
	case torrent.TvHd, torrent.TvWebDl:
		return 208

	default:
		return 0
	}
}