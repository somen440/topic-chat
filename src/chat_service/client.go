package main

type client struct {
	id   int
	send chan []byte
}

var clients = map[*client]bool{}
