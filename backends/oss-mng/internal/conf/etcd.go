package conf

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	cfg "github.com/go-kratos/kratos/contrib/config/etcd/v2"
	"github.com/go-kratos/kratos/v2/config"
)

func NewEtcdSource(addr, path string, timeout time.Duration) config.Source {
	// create an etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		log.Fatal(err)
	}

	// configure the source, "path" is required
	source, err := cfg.New(client, cfg.WithPath(path), cfg.WithPrefix(true))
	if err != nil {
		log.Fatalln(err)
	}

	return source
}
