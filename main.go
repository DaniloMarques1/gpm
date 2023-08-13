package main

import "log"

func main() {
	s, err := NewServer()
	if err != nil {
		log.Fatalf("ERROR creating server %v\n", err)
	}
	log.Fatal(s.Run())
}
