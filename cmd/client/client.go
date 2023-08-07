package main

import (
	"context"
	"flag"
	"fmt"
	"library/client"
	"log"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	cl, err := client.NewLibraryClient(addr)

	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()

	resp, err := cl.GetBooksByAuthor(ctx, "Author1")

	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp != nil {
		for _, book := range resp.Books {
			fmt.Println(book)
		}
	}
}
