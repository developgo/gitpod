// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package main

import (
	"context"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"

	"github.com/gitpod-io/gitpod/test/pkg/integration"
	test_context "github.com/gitpod-io/gitpod/test/pkg/integration/context"
)

func TestBuiltinUserExists(t *testing.T) {
	builtinUser := features.New("database").
		WithLabel("type", "check builtin user").
		Setup(func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			it := integration.NewTest(ctx, t, test_context.GetNamespace(ctx), cfg.Client())
			return test_context.SetIntegrationTest(ctx, it)
		}).
		Assess("builtin user workspace exists", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			it := test_context.GetIntegrationTest(ctx)

			db := it.API().DB()

			rows, err := db.Query(`SELECT count(1) AS count FROM d_b_user WHERE id ="builtin-user-workspace-probe-0000000"`)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()

			if !rows.Next() {
				t.Fatal("no rows selected - should not happen")
			}

			var count int
			err = rows.Scan(&count)
			if err != nil {
				t.Fatal(err)
			}

			if count != 1 {
				t.Fatalf("expected a single builtin-user-workspace-probe-0000000, but found %d", count)
			}

			return ctx
		}).
		Teardown(func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			it := test_context.GetIntegrationTest(ctx)
			it.Done()

			return ctx
		}).Feature()

	testenv.Test(t, builtinUser)
}
