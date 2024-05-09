package etcd

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"gin-app/pkg/log"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(NewEtcdClient)

var config clientv3.Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Log.Error("Loading .env file error.tsx: %s", err)
		panic("Error loading .env file")
	}

	config = clientv3.Config{
		Endpoints: []string{os.Getenv("ETCD_ENDPOINTS")},
		//Username:  os.Getenv("ETCD_USERNAME"),
		//Password:    os.Getenv("ETCD_PASSWORD"),
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		DialTimeout: time.Second * 5,
	}
}

func NewEtcdClient() *clientv3.Client {
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	return client
}

func LoadConfigFromPath[T any](client *clientv3.Client, path string) (t *T, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, err
	}

	if err = json.Unmarshal(resp.Kvs[0].Value, &t); err != nil {
		return nil, err
	}

	return
}
