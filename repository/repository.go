package staff_repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindStaffByID(ctx context.Context, accountId string) ([]byte, error)
	CreateAccount(ctx context.Context, staff any) (*mongo.InsertOneResult, error)
}
