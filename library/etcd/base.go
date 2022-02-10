package etcd

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Etcd struct {
	Client *clientv3.Client
}

func init()  {
	_, err := GetEtcd()
	if err != nil {
		fmt.Printf("Etcd connect failed , error is %v\n", err)
		panic(err)
	}
}
var EClient  = new(clientv3.Client)

func GetEtcd() (*Etcd, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 3,
	})

	if err != nil {
		return nil ,err
	}
	EClient = client
	return &Etcd{Client: client}, nil
}


