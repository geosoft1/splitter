// simple http splitter
// Copyright (C) 2018  geosoft1  geosoft1@gmail.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Ip      string   `json:"ip"`
	Port    string   `json:"port"`
	Handler string   `json:"handler"`
	Routes  []string `json:"routes"`
}

var c Config

var client = &http.Client{
	Timeout: 100 * time.Millisecond,
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RemoteAddr, r.RequestURI)
	for _, route := range c.Routes {
		if _, err := client.Get(route + r.RequestURI); err != nil {
			log.Print(err.Error())
		}
	}
}

func main() {
	log.Print("init logger")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		os.Exit(1)
	}

	log.Print("load configuration")
	f, err := os.Open(filepath.ToSlash(pwd + "/conf.json"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("%s > %v", c.Handler, c.Routes)
	http.HandleFunc(c.Handler, handler)

	log.Printf("start listening on %s:%s", c.Ip, c.Port)
	http.ListenAndServe(c.Ip+":"+c.Port, nil)
}
