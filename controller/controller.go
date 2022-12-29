package staff_ctl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	staff_config "github.com/lifeentify/staff/config"
	staff_repo "github.com/lifeentify/staff/repository"
	staff_db "github.com/lifeentify/staff/repository/db"
	staff "github.com/lifeentify/staff/staff/v1"
)

type Controller struct {
	DB     staff_repo.Repository
	Config *staff_config.Config
}

const (
	Mongo    = "MONGODB"
	MySQL    = "MYSQL"
	PostGres = "POSTGRES"
	Category = "STAFF"
)

func NewController(config *staff_config.Config) (*Controller, error) {
	dbType := config.DatabaseType

	if dbType == Mongo {
		uri := config.MongoUrl
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}
		db := staff_db.NewMongoDB(config)
		return &Controller{
			db,
			config,
		}, nil
	}
	return nil, fmt.Errorf("no match fund for db type :%s", dbType)
}

func (c *Controller) Save(newStaff *staff.Staff) (*staff.Staff, error) {
	staffJson, err := newStaff.ToJson()
	if err != nil {
		return nil, err
	}
	result, err := c.DB.CreateAccount(context.TODO(), staffJson)
	if err != nil {
		return nil, err
	}
	savedStaffByte, err := c.DB.FindStaffByID(context.TODO(), fmt.Sprintf("%s", result.InsertedID))
	if err != nil {
		return nil, err
	}
	var savedStaff staff.Staff
	json.Unmarshal(savedStaffByte, &savedStaff)
	return &savedStaff, err
}
