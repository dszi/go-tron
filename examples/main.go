//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package main

import (
	"fmt"
	"log"

	"github.com/dszi/go-tron/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "grpc.nile.trongrid.io:50051"

	client := pkg.NewGrpcClient(address,
		pkg.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)

	if err := client.Start(); err != nil {
		log.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer client.Stop()

	testAddr := "TMWXhuxiT1KczhBxCseCDDsrhmpYGUcoA9"

	account, err := client.GetAccount(testAddr)
	if err != nil {
		log.Fatalf("GetAccount error: %v", err)
	}
	fmt.Printf("Account info: %+v\n", account)

	if len(account.Address) == 0 {
		log.Println("Expected non-empty account address")
	}
}
