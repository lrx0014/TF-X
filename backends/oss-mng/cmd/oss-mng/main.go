package main

import (
	"flag"
	"os"
	"time"

	"oss-mng/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// etcd configs
	etcdEntry   string
	etcdPath    string
	etcdTimeout int64 // in seconds

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf config.yaml")
	flag.StringVar(&etcdEntry, "etcd", "", "etcd endpoint, e.g, 127.0.0.1:2379")
	flag.StringVar(&etcdPath, "etcd_path", "/oss-mng", "etcd path, e.g, /app-conf")
	flag.Int64Var(&etcdTimeout, "etcd_timeout", 60, "etcd timeout in seconds, default to 60")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	c := config.New(
		config.WithSource(
			makeSources()...,
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func makeSources() (sources []config.Source) {
	if flagconf != "" {
		sources = append(sources, file.NewSource(flagconf))
	}
	if etcdEntry != "" {
		sources = append(sources, conf.NewEtcdSource(etcdEntry, etcdPath, time.Duration(etcdTimeout)*time.Second))
	}

	return
}
