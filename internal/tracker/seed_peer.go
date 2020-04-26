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
	"strings"
)

type SeedPeer struct { }

func (t *SeedPeer) Search(str string, category torrent.Category) []torrent.Result {
	c := t.convertCategory(category)
	searchUrl := fmt.Sprintf("https://www.seedpeer.me/search/%s?category=%s?sort=seeds", url.PathEscape(str), c)
	results := make([]torrent.Result, 0)
	doc, err := getPage(searchUrl)
	if err != nil {
		log.Error("Failed to retrieve page", log.String("host", "www.seedpeer.me"), log.Err(err))
		return results
	}

	doc.Find("div.table-container table.table tbody tr[data-reactid]").Each(func(i int, s *goquery.Selection) {
		var r torrent.Result
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				link := doc.Find("a")
				href, exists := link.Attr("href")
				if !exists || !strings.HasPrefix(href, "/details/") {
					return
				}
				r.Magnet, err = torrent.MagnetFromLink(t.getMagnetLink("https://www.seedpeer.me" + href))
				if err != nil {
					log.Error("Failed to parse magnet link", log.Err(err))
				}
			}
		})
		results = append(results, r)
	})

	return results
}

func (t *SeedPeer) convertCategory(c torrent.Category) string {
	switch c {
	case torrent.Audio, torrent.AudioMP3, torrent.AudioVideo, torrent.AudioAudioBook, torrent.AudioLossless, torrent.AudioForeign, torrent.AudioOther:
		return "musics"
	case torrent.Movies, torrent.MoviesForeign, torrent.MoviesOther, torrent.MoviesSD, torrent.MoviesHD, torrent.Movies3D, torrent.MoviesBluray, torrent.MoviesDvd, torrent.MoviesWebDl:
		return "movies"
	case torrent.Tv, torrent.TvWebDl, torrent.TvForeign, torrent.TvSd, torrent.TvHd, torrent.TvSport, torrent.TvAnime, torrent.TvDocumentary, torrent.TvOther:
		return "tv"
	default:
		return ""
	}
}

func (t *SeedPeer) getMagnetLink(detailsUrl string) string {
	var magnetLink string
	doc, err := getPage(detailsUrl)
	if err != nil {
		log.Error("Failed to retrieve page", log.Err(err))
		return ""
	}

	doc.Find("div.inner_container div.grey_bar1").Each(func(i int, s *goquery.Selection) {
		link := s.Find("p a")
		href, exists := link.Attr("href")
		if !exists || !strings.HasPrefix(href, "magnet:") {
			return
		}
		magnetLink = href
	})
	return magnetLink
}
