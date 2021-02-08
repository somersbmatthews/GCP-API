package pg

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {

}

func createUser(user User) {
	ctx := context.Background()
	payload, ok := CreateUser(ctx, user)

}
