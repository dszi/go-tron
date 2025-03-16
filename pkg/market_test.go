package pkg

import (
	"fmt"
	"testing"
)

func TestGrpcClient_GetMarketPairList(t *testing.T) {
	client := setupGrpcClient(t)
	defer client.Stop()

	account, err := client.GetMarketPairList()
	if err != nil {
		t.Fatalf("GetMarketPairList error: %v", err)
	}
	fmt.Println(account)
}

func TestGrpcClient_GetMarketPriceByPair(t *testing.T) {
	client := setupGrpcClient(t)
	defer client.Stop()

	account, err := client.GetMarketPriceByPair("1000094", "1000095")
	if err != nil {
		t.Fatalf("GetMarketPairList error: %v", err)
	}
	fmt.Println(account)
}

func TestGrpcClient_GetMarketOrderById(t *testing.T) {
	client := setupGrpcClient(t)
	defer client.Stop()

	account, err := client.GetMarketOrderById("1000094")
	if err != nil {
		t.Fatalf("GetMarketPairList error: %v", err)
	}
	fmt.Println(account)
}

func TestGrpcClient_GetMarketOrderListByPair(t *testing.T) {
	client := setupGrpcClient(t)
	defer client.Stop()

	account, err := client.GetMarketPriceByPair("1000094", "1000095")
	if err != nil {
		t.Fatalf("GetMarketPairList error: %v", err)
	}
	fmt.Println(account)
}

func TestGrpcClient_GetMarketOrderByAccount(t *testing.T) {
	client := setupGrpcClient(t)
	defer client.Stop()

	testAddr := "TDS7NjQwQn7iNBN1UxxpsFD5UB3pxB7msL"
	tx, err := client.GetMarketOrderByAccount(testAddr)
	if err != nil {
		t.Fatalf("UpdateAccount error: %v", err)
	}
	t.Logf("UpdateAccount tx: %+v", tx)
}
