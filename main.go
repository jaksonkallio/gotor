package main

import "github.com/jaksonkallio/go-torrent-tracker/server"

func main() {
	server := server.Server{}
	server.Run()
}
