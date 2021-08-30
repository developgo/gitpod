package main

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"

	"sigs.k8s.io/e2e-framework/klient"
	"sigs.k8s.io/e2e-framework/klient/conf"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	test_context "github.com/gitpod-io/gitpod/test/pkg/integration/context"
)

var (
	testenv env.Environment

	namespace = flag.String("namespace", "", `namespace to execute the test against. Defaults to the one configured in "kubeconfig".`)
	username  = flag.String("username", "", "username to execute the tests with. Chooses one automatically if left blank.")
)

func TestMain(m *testing.M) {
	flag.Parse()

	kubecfg, err := conf.New(conf.ResolveKubeConfigFile())
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	client, err := klient.New(kubecfg)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	conf := envconf.New()
	conf.WithClient(client)

	testenv = env.NewWithConfig(conf)
	testenv.Setup(
		checkGitpodIsRunning(),
		setup(),
		setupComponents(),
	)
	testenv.Finish(
		finish(),
	)

	os.Exit(testenv.Run(m))
}

func checkGitpodIsRunning() env.Func {
	return func(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
		return ctx, nil
	}
}

func setup() env.Func {
	return func(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
		ctx = test_context.SetNamespace(ctx, *namespace)
		ctx = test_context.SetUsername(ctx, *username)
		// setup gitpod components
		return ctx, nil
	}
}

func setupComponents() env.Func {
	return func(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
		return ctx, nil
	}
}

func finish() env.Func {
	return func(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
		return ctx, nil
	}
}
