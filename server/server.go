package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"time"

	"fake.com/instore/storage"
	"fake.com/instore/storepb"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	"google.golang.org/grpc"
)

// initializeing filePath to for later use
var _, b, _, _ = runtime.Caller(0)

// basepath = "/path/to/server" the directory in which we find server.go
var basepath = filepath.Dir(b)

// filePath = /path/to/server/csvFiles/ the directory in which we find
// csv file to be used by server. ie where the data is stored
var filePath = basepath + "/csvFiles/" + filename

const (
	bSvr     string = "[svr]  "
	filename string = "data.csv"
)

type ItemBD struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type server struct {
	storepb.UnimplementedStoreServiceServer
}

var m map[string]string

func main() {

	// helps with detecting line in which we got an error <3 in case...
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s := grpc.NewServer()
	// listen on tcp
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Printf(bSvr+"%v\n", err)
		os.Exit(1)
	}
	// Registers a services and its implementations to the gRPC server.
	storepb.RegisterStoreServiceServer(s, &server{})

	m = make(map[string]string)

	go func() {
		fmt.Println(bSvr + "Oh you chose to run me :D!")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(bSvr + "Good choice ;)")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(bSvr + "waiting for the client...")
		err := s.Serve(lis)
		if err != nil {
			fmt.Printf(bSvr+"%v\n", err)
			os.Exit(1)
		}
	}()

	// wait for ctrl-C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is recieved
	<-ch

	s.Stop()
	lis.Close()
	fmt.Println(" ------Interrupted with Ctrl-C------")
	fmt.Println(bSvr + "Thank you for your stay! :D")
	fmt.Println(bSvr + "Have a good day!")
}

// UseCsv is a funcion that changes storing mode to CSV
func (*server) Use(ctx context.Context, req *storepb.UseRequest) (*storepb.UseResponse, error) {
	fmt.Println(bSvr + "Use function was invoked with " + req.Msg)
	storage.StoreType = req.Msg
	result := bSvr + "Using CSV mode now..."
	res := &storepb.UseResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Insert(ctx context.Context, req *storepb.InsertRequest) (*storepb.InsertResponse, error) {
	fmt.Printf(bSvr+"Insert function was invoked with %v\n", req)
	key := req.GetKey()
	value := req.GetValue()
	m[key] = value
	result := bSvr + "Inserted: {'" + key + "': '" + value + "'}"
	res := &storepb.InsertResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Select(ctx context.Context, req *storepb.SelectRequest) (*storepb.SelectResponse, error) {
	fmt.Printf(bSvr+"Select function was invoked with %v\n", req)

	resp := &storepb.SelectResponse{
		Result: fmt.Sprint(m),
	}
	return resp, nil
}
