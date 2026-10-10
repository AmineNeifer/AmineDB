package client

import (
	"context"
	"fmt"
	"os"

	"fake.com/instore/storepb"
)

const bClt string = "[clt] "

// UseCsv is a funcion that changes storing mode to CSV
func Use(c storepb.StoreServiceClient, storageType string) {
	req := &storepb.UseRequest{
		Msg: storageType,
	}
	res, err := c.Use(context.Background(), req)

	if err != nil {
		fmt.Printf(bClt+"%v\n", err)
		os.Exit(1)
	}
	fmt.Println(res.Result)
}

func Insert(c storepb.StoreServiceClient, key string, value string) {
	req := &storepb.InsertRequest{
		Key:   key,
		Value: value,
	}
	res, err := c.Insert(context.Background(), req)
	if err != nil {
		fmt.Printf(bClt+"%v\n", err)
		os.Exit(1)
	}
	fmt.Println(res.Result)
}

func Select(c storepb.StoreServiceClient, key string) {
	req := &storepb.SelectRequest{
		Key: key,
	}
	res, err := c.Select(context.Background(), req)
	if err != nil {
		fmt.Printf(bClt+"%v\n", err)
		os.Exit(1)
	}
	fmt.Println(res.Result)
}