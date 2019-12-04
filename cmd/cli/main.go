package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/MohammadTaheri304/discovery/rpc"
	"github.com/MohammadTaheri304/discovery/service"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Sample args: <discovery-address localhost:9090>")
	address := "0.0.0.0:9090"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	client(address)
}

func client(address string) {
	dServer := dialDiscoveryServer(address)

	_, pubKey := service.GenerateKeyPair()
	pubString := service.PublicKeyToString(pubKey)
	log.Println("Your public key is: " + pubString)

	dServer.Register(context.Background(), &rpc.Message{
		Key:   pubString,
		Value: GetOutboundIP().String() + ":" + strconv.Itoa(randomPort()),
	})

	res, err := dServer.Get(context.Background(), &rpc.Message{
		Key: pubString,
	})
	if err != nil {
		log.Fatal("Error in get value %+v", err)
	}
	log.Printf("Result was %+v \n", res)
}

func dialDiscoveryServer(address string) rpc.DiscoveryServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to Discovery server. %v\n", err)
	}

	return rpc.NewDiscoveryServiceClient(conn)
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func randomPort() int {
	return 3000 + rand.New(rand.NewSource(time.Now().Unix())).Intn(5000)
}
