package context

import (
	"context"

	"github.com/gitpod-io/gitpod/test/pkg/integration"
)

const (
	integrationTest = "integration-test"
	namespace       = "namespace"
	username        = "username"
)

type contextKey string

func GetNamespace(ctx context.Context) string {
	return ctx.Value(contextKey(namespace)).(string)
}

func SetNamespace(ctx context.Context, ns string) context.Context {
	if ns == "" {
		ns = "default"
	}

	return context.WithValue(ctx, contextKey(namespace), ns)
}

func GetIntegrationTest(ctx context.Context) *integration.Test {
	return ctx.Value(contextKey(integrationTest)).(*integration.Test)
}

func SetIntegrationTest(ctx context.Context, it *integration.Test) context.Context {
	return context.WithValue(ctx, contextKey(integrationTest), it)
}

func GetUsername(ctx context.Context) string {
	return ctx.Value(contextKey(username)).(string)
}

func SetUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, contextKey(username), username)
}
