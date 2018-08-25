// +build integration

package travis

import (
	"context"
	"testing"
)

func TestBranchesService_ListFromRepository(t *testing.T) {
	t.Parallel()

	_, _, err := integrationClient.Branches.ListFromRepository(context.TODO(), integrationRepo)
	ok(t, err)
}
