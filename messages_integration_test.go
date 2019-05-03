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
)

const integrationRequestId = 161917659

func TestMessagesService_Integration_ListByRepoId(t *testing.T) {
	opt := MessagesOption{Limit: 5}
	_, res, err := integrationClient.Messages.ListByRepoId(context.TODO(), integrationRepoId, integrationRequestId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestMessagesService_Integration_ListByRepoSlug(t *testing.T) {
	opt := MessagesOption{Limit: 5}
	_, res, err := integrationClient.Messages.ListByRepoSlug(context.TODO(), integrationRepoSlug, integrationRequestId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
