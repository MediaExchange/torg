# torg

[![Build Status](https://travis-ci.org/mediaexchange/torg.svg)](https://travis-ci.org/mediaexchange/torg)
[![GoDoc](https://godoc.org/github.com/mediaexchange/torg/github?status.svg)](https://godoc.org/github.com/mediaexchange/torg)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.13-green.svg)](https://golang.org/doc/devel/release.html#go1.13)
[![Go version](https://img.shields.io/badge/go-~%3E1.14-green.svg)](https://golang.org/doc/devel/release.html#go1.14)

The Torrent Aggregator (torg) project provides a search API for Go applications that targets multiple Torrent trackers.
The API accepts a string, such as `The Terminator 1984 1080p`, and returns a collection of results that include magnet 
link, torrent name, health of the torrent (seeders vs. leechers), and size of the torrent.

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Media Exchange](http://github.com/MediaExchange)

## License

   Copyright 2019 MediaExchange.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
