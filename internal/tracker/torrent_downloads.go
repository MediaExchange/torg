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

type TorrentDownloads struct { }

func (t *TorrentDownloads) Search(str string, category torrent.Category) []torrent.Result {
	c := t.convertCategory(category)
	searchUrl := fmt.Sprintf("https://www.torrentdownloads.me/search/?s_cat=%d&search=%s", c, url.QueryEscape(str))
	results := make([]torrent.Result, 0)
	doc, err := getPage(searchUrl)
	if err != nil {
		log.Error("Failed to retrieve page", log.String("host", "www.torrentdownloads.me"), log.Err(err))
		return results
	}

	doc.Find("div.inner_container div.grey_bar3").Each(func(i int, s *goquery.Selection) {
		link := s.Find("p a")
		href, exists := link.Attr("href")
		if !exists || !strings.HasPrefix(href, "/torrent/") {
			return
		}
		var r torrent.Result

		r.Magnet, err = torrent.MagnetFromLink(t.getMagnetLink("https://www.torrentdownloads.me" + href))
		if err != nil {
			log.Error("Failed to parse magnet link", log.Err(err))
		}

		s.Find("span").Each(func(i int, s *goquery.Selection) {
			if i == 2 {
				r.Seeders, _ = strconv.Atoi(s.Text())
			}
		})

		results = append(results, r)
	})

	return results
}

func (t *TorrentDownloads) convertCategory(c torrent.Category) int {
	switch c {
	case torrent.Movies, torrent.MoviesForeign, torrent.MoviesOther, torrent.MoviesSD, torrent.MoviesHD, torrent.Movies3D, torrent.MoviesBluray, torrent.MoviesDvd, torrent.MoviesWebDl:
		// "Movies"
		return 4
	case torrent.Tv, torrent.TvWebDl, torrent.TvHd, torrent.TvAnime, torrent.TvDocumentary, torrent.TvForeign, torrent.TvOther, torrent.TvSd, torrent.TvSport:
		// "TV Shows"
		return 8
	case torrent.Audio, torrent.AudioMP3, torrent.AudioVideo, torrent.AudioAudioBook, torrent.AudioLossless, torrent.AudioForeign, torrent.AudioOther:
		// "Music"
		return 3
	default:
		// "Any"
		return 0
	}
}

func (t *TorrentDownloads) getMagnetLink(detailsUrl string) string {
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