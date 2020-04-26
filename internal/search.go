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

package internal

import (
	"github.com/MediaExchange/log"
	"github.com/MediaExchange/torg/internal/tracker"
	"github.com/MediaExchange/torg/torrent"
	"sync"
	"time"
)

// Each tracker must implement this interface.
type Tracker interface {
	Search(str string, category torrent.Category) []torrent.Result
}

type Provider struct {
	Name string
	Impl Tracker
}

// trackers acts as a registry of all the supported trackers.
var trackers = []Provider {
	{"1337x", new(tracker.L337x)},
	{"SeedPeer", new(tracker.SeedPeer)},
	{"ThePirateBay", new(tracker.ThePirateBay)},
	{"TorrentDownloads", new(tracker.TorrentDownloads)},
}

// Search is the internal implementation of the torrent search aggregator.
func Search(str string, category torrent.Category) []torrent.Result {
	res := make([]torrent.Result, 0)
	var waiter sync.WaitGroup
	waiter.Add(len(trackers))

	for _, p := range trackers {
			go func(p Provider) {
				log.Info("Search started", log.String("provider", p.Name))
				start := time.Now()
				r := p.Impl.Search(str, category)
				end := time.Now()
				n := len(r)
				log.Info("Search finished", log.String("provider", p.Name), log.Int("count", n), log.String("duration", end.Sub(start).String()))
				res = append(res, r...)
				waiter.Done()
			}(p)
	}

	waiter.Wait()
	return res
}
