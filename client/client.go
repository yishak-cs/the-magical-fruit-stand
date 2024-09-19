package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	the_magical_fruit_stand "github.com/yishak-cs/the-magical-fruit-stand/protogen/golang"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := the_magical_fruit_stand.NewFruitSageClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a fruit name (or 'quit' to exit): ")
		fruit, _ := reader.ReadString('\n')
		fruit = strings.TrimSpace(fruit)

		if strings.ToLower(fruit) == "quit" {
			break
		}

		r, err := c.AskAboutFruit(context.Background(), &the_magical_fruit_stand.FruitQuery{FruitName: fruit})
		if err != nil {
			log.Fatalf("could not get wisdom: %v", err)
		}
		fmt.Printf("Fruit Sage says: %s\n", r.GetWisdom())
	}
}
