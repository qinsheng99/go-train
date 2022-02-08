package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var client = Etcd{}
func Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	get, err := client.Client.Get(ctx, key, opts...)
	if err != nil {
		return nil, err
	}
	return get, err
}

func Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	put, err := client.Client.Put(ctx, key, val, opts...)
	if err != nil {
		return nil, err
	}
	return put, nil
}
