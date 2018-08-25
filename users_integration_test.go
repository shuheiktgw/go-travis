// +build integration

package travis

import (
	"context"
	"testing"
	"time"
)

func TestUsersService_GetAuthenticated(t *testing.T) {
	t.Parallel()

	if auth := integrationClient.IsAuthenticated(); !auth {
		t.Skip("test client is unauthenticated. skipping.")
	}

	user, _, err := integrationClient.Users.GetAuthenticated(context.TODO())
	ok(t, err)

	assert(
		t,
		user != nil,
		"UsersService.GetAuthenticated returned nil user",
	)
}

func TestUsersService_Get(t *testing.T) {
	t.Parallel()

	if auth := integrationClient.IsAuthenticated(); !auth {
		t.Skip("test client is unauthenticated. skipping.")
	}

	authenticatedUser, _, err := integrationClient.Users.GetAuthenticated(context.TODO())
	userId := authenticatedUser.Id

	user, _, err := integrationClient.Users.Get(context.TODO(), userId)
	ok(t, err)

	assert(
		t,
		authenticatedUser != nil,
		"UsersService.Get returned nil user",
	)

	assert(
		t,
		user.Id == userId,
		"UsersService.Get returned user with id %d; expected %d", user.Id, userId,
	)
}

func TestUsersService_Sync(t *testing.T) {
	t.Parallel()

	if auth := integrationClient.IsAuthenticated(); !auth {
		t.Skip("test client is unauthenticated. skipping.")
	}

	userNow, _, err := integrationClient.Users.GetAuthenticated(context.TODO())
	_, err = integrationClient.Users.Sync(context.TODO())
	ok(t, err)
	time.Sleep(5 * time.Second)

	userThen, _, err := integrationClient.Users.GetAuthenticated(context.TODO())
	if userThen != nil {
		// User might have been destroyed since last sync
		assert(
			t,
			userThen.SyncedAt > userNow.SyncedAt,
			"UsersService.Sync does not have updated the user synced_at marker",
		)
	}
}
