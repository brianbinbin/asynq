// Copyright 2020 Kentaro Hibino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package asynq

import (
	"sort"
	"testing"

	h "github.com/brianbinbin/asynq/internal/asynqtest"
	"github.com/go-redis/redis/v7"
	"github.com/google/go-cmp/cmp"
)

// This file defines test helper functions used by
// other test files.

func setup(tb testing.TB) *redis.Client {
	tb.Helper()
	r := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   14,
	})
	// Start each test with a clean slate.
	h.FlushDB(tb, r)
	return r
}

var sortTaskOpt = cmp.Transformer("SortMsg", func(in []*Task) []*Task {
	out := append([]*Task(nil), in...) // Copy input to avoid mutating it
	sort.Slice(out, func(i, j int) bool {
		return out[i].Type < out[j].Type
	})
	return out
})
