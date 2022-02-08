package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	get, err := EClient.Get(ctx, key, opts...)
	if err != nil {
		return nil, err
	}
	return get, err
}

func Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	put, err := EClient.Put(ctx, key, val, opts...)
	if err != nil {
		return nil, err
	}
	return put, nil
}
