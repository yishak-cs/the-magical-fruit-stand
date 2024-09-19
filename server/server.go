package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	the_magical_fruit_stand "github.com/yishak-cs/the-magical-fruit-stand/protogen/golang"
	"google.golang.org/grpc"
)

type server struct {
	the_magical_fruit_stand.UnimplementedFruitSageServer
}

func (s *server) AskAboutFruit(ctx context.Context, in *the_magical_fruit_stand.FruitQuery) (*the_magical_fruit_stand.FruitWisdom, error) {
	fruit := strings.ToLower(in.GetFruitName())
	var wisdom string

	switch fruit {
	case "apple":
		wisdom = "An apple a day keeps the doctor away, but a pear is also fair!"
	case "banana":
		wisdom = "Bananas are nature's energy bars, curved for your convenience!"
	default:
		wisdom = fmt.Sprintf("Ah, the %s! A fruit of mystery and wonder!", fruit)
	}

	return &the_magical_fruit_stand.FruitWisdom{Wisdom: wisdom}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	the_magical_fruit_stand.RegisterFruitSageServer(s, &server{})
	fmt.Println("Fruit Sage is ready to share wisdom on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
