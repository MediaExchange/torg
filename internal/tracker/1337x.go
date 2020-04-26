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

type L337x struct { }

func (t *L337x) Search(str string, category torrent.Category) []torrent.Result {
	c := t.convertCategory(category)
	searchUrl := fmt.Sprintf("https://www.1337x.to/category-search/%s/%s/1/", url.PathEscape(str), c)
	results := make([]torrent.Result, 0)
	doc, err := getPage(searchUrl)
	if err != nil {
		log.Error("Failed to retrieve page", log.String("host", "www.1337x.to"), log.Err(err))
		return results
	}

	doc.Find("tbody tr").Each(func(i int, tr *goquery.Selection){
		// Result appended to the slice of results.
		r := torrent.Result{}

		// The magnet link is found on a details page for the torrent.
		href, exists := tr.Find("td.name a:nth-of-type(2)").First().Attr("href")
		if !exists || !strings.HasPrefix(href, "/torrent/") {
			return
		}

		r.Magnet, err = torrent.MagnetFromLink(t.getMagnetLink("https://www.1337x.to" + href))
		if err != nil {
			// If the magnet link couldn't be found/parsed then there's no reason to continue with this selector.
			log.Error("Failed to parse magnet link", log.Err(err))
			return
		}

		// Number of seeders
		r.Seeders, _ = strconv.Atoi(tr.Find("td.seeds").First().Text())

		// Size of torrent
		r.Size = tr.Find("td.size").Nodes[0].FirstChild.Data

		// Add to the slice
		results = append(results, r)
	})

	return results
}

// convertCategory converts from a category constant to the string used in the query URL.
func (t *L337x) convertCategory(c torrent.Category) string {
	switch c {
	case torrent.Audio, torrent.AudioMP3, torrent.AudioVideo, torrent.AudioAudioBook, torrent.AudioLossless, torrent.AudioForeign, torrent.AudioOther:
		return "Music"
	case torrent.Movies, torrent.MoviesForeign, torrent.MoviesOther, torrent.MoviesSD, torrent.MoviesHD, torrent.Movies3D, torrent.MoviesBluray, torrent.MoviesDvd, torrent.MoviesWebDl:
		return "Movies"
	case torrent.Tv, torrent.TvWebDl, torrent.TvForeign, torrent.TvSd, torrent.TvHd, torrent.TvSport, torrent.TvAnime, torrent.TvDocumentary, torrent.TvOther:
		return "TV"
	default:
		return ""
	}
}

// getMagnetLink retrieves the child page to extract the magnet link.
func (t *L337x) getMagnetLink(detailsUrl string) (magnetLink string) {
	doc, err := getPage(detailsUrl)
	if err != nil {
		log.Error("Failed to retrieve page", log.String("url", detailsUrl), log.Err(err))
		return
	}

	link := doc.Find("div.page-content div.torrent-detail-page div:nth-child(2) div:first-child ul:first-child > li:first-child > a").First()
	href, exists := link.Attr("href")
	if !exists || !strings.HasPrefix(href, "magnet:") {
		return
	}

	magnetLink = href
	return
}
