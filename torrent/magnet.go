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

package torrent

import (
	"errors"
	"net/url"
	"strings"
)

// Magnet represents the components of a magnet link.
type Magnet struct {
	Link string
	Name string
	Hash struct {
		Type  string
		Value string
	}
	Trackers []string
}

func MagnetFromLink(link string) (m Magnet, err error) {
	// Store the link
	m.Link = link

	// A magnet link is just a query string with a scheme of `magnet`
	var values url.Values
	values, err = url.ParseQuery(strings.TrimPrefix(link, "magnet:?"))
	if err != nil {
		return
	}

	// Exact Topic. The magnet spec states that multiple topics may be present in the form `xt.1`, `xt.2` and so on,
	// but in practice this is rarely used.
	var xt string
	xt, err = url.QueryUnescape(values.Get("xt"))
	if err != nil {
		return
	}

	hash := strings.Split(xt, ":")
	if len(hash) != 3 {
		return m, errors.New("invalid xt")
	}

	if hash[0] != "urn" {
		return m, errors.New("xt does not contain urn")
	}

	m.Hash.Type = hash[1]
	m.Hash.Value = hash[2]

	// Display Name
	m.Name, err = url.QueryUnescape(values.Get("dn"))
	if err != nil {
		return
	}

	// Trackers
	m.Trackers = values["tr"]
	return
}
