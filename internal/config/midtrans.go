package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransConfig struct {
	ClientKey string
	ServerKey string
}

func NewMidtransClient(config MidtransConfig) snap.Client {
	var client snap.Client
	client.New(config.ServerKey, midtrans.Sandbox)
	return client
}
