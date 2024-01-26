package main

import "fmt"

type Codec struct {
}

func Constructor() Codec {

}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {

}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {

}

func main() {
	fmt.Println("vim-go")
}
