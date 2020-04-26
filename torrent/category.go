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

type Category int

// Category names and values taken from the nZEDb project:
// https://github.com/nZEDb/nZEDb/blob/dev/docs/newznab_api_specification.txt
const (
	Other       Category = 0000
	OtherMisc   Category = 0010
	OtherHashed Category = 0020

	Console           Category = 1000
	ConsoleNds        Category = 1010
	ConsolePsp        Category = 1020
	ConsoleWii        Category = 1030
	ConsoleXbox       Category = 1040
	ConsoleXbox360    Category = 1050
	ConsoleWiiWare    Category = 1060
	ConsoleXbox360Dlc Category = 1070
	ConsolePs3        Category = 1080
	Console3ds        Category = 1110
	ConsolePsVita     Category = 1120
	ConsoleWiiU       Category = 1130
	ConsoleXboxOne    Category = 1140
	ConsolePs4        Category = 1180
	ConsoleOther      Category = 1999

	Movies        Category = 2000
	MoviesForeign Category = 2010
	MoviesOther   Category = 2020
	MoviesSD      Category = 2030
	MoviesHD      Category = 2040
	Movies3D      Category = 2050
	MoviesBluray  Category = 2060
	MoviesDvd     Category = 2070
	MoviesWebDl   Category = 2080

	Audio          Category = 3000
	AudioMP3       Category = 3010
	AudioVideo     Category = 3020
	AudioAudioBook Category = 3030
	AudioLossless  Category = 3040
	AudioForeign   Category = 3060
	AudioOther     Category = 3999

	Pc             Category = 4000
	Pc0Day         Category = 4010
	PcIso          Category = 4020
	PcMac          Category = 4030
	PcPhoneOther   Category = 4040
	PcGames        Category = 4050
	PcPhoneIds     Category = 4060
	PcPhoneAndroid Category = 4070

	Tv            Category = 5000
	TvWebDl       Category = 5010
	TvForeign     Category = 5020
	TvSd          Category = 5030
	TvHd          Category = 5040
	TvSport       Category = 5060
	TvAnime       Category = 5070
	TvDocumentary Category = 5080
	TvOther       Category = 5999

	Xxx         Category = 6000
	XxxDvd      Category = 6010
	XxxWmv      Category = 6020
	XxxXvid     Category = 6030
	XxxX264     Category = 6040
	XxxImageSet Category = 6060
	XxxPacks    Category = 6070
	XxxOther    Category = 6999

	Books          Category = 7000
	BooksMagazines Category = 7010
	BooksEbook     Category = 7020
	BooksComic     Category = 7030
	BooksTechnical Category = 7040
	BooksForeign   Category = 7060
	BooksUnknown   Category = 7999
)
