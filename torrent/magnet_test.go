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
	"github.com/MediaExchange/assert"
	"testing"
)

func TestFromLink(t *testing.T) {
	Assert := assert.With(t)

	link := "magnet:?xt=urn:btih:5C453A86DA27F7A47AD6ADBCB23FEFC7F3F6287C&dn=Bad.Boys.for.Life.2020.720p.HDRip.900MB.x264-GalaxyRG+⭐&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969&tr=udp%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=udp%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.demonii.si%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.pirateparty.gr%3A6969%2Fannounce&tr=udp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce"
	m, err := MagnetFromLink(link)

	Assert.That(err).IsOk()
	Assert.That(m.Link).IsEqualTo(link)
	Assert.That(m.Name).IsEqualTo("Bad.Boys.for.Life.2020.720p.HDRip.900MB.x264-GalaxyRG ⭐")
	Assert.That(m.Hash.Type).IsEqualTo("btih")
	Assert.That(m.Hash.Value).IsEqualTo("5C453A86DA27F7A47AD6ADBCB23FEFC7F3F6287C")
	Assert.That(len(m.Trackers)).IsEqualTo(17)
}