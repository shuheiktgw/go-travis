// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestBroadcastsService_Integration_List(t *testing.T) {
	cases := []*BroadcastsOption{
		nil,
		{Active: true},
		{Active: false},
	}

	for i, opt := range cases {
		_, res, err := integrationClient.Broadcasts.List(context.TODO(), opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		// Be nice to the API
		time.Sleep(2 * time.Second)
	}
}
