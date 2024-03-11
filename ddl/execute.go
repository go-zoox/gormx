package ddl

import (
	"context"
	"fmt"

	"github.com/go-zoox/gormx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func execute(engine, dsn, sql string) error {
	db, err := gormx.Connect(engine, dsn)
	if err != nil {
		return fmt.Errorf("connecting database failed: %s", err.Error())
	}
	conn, err := db.DB()
	if err != nil {
		return fmt.Errorf("getting database connection failed: %s", err.Error())
	}
	defer conn.Close()

	var f interface{}
	return db.Raw(sql).Scan(&f).Error
}

func executeMongoDB(dsn string, do func(ctx context.Context, client *mongo.Client) error) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return fmt.Errorf("connecting database failed: %s", err.Error())
	}
	defer client.Disconnect(ctx)

	return do(ctx, client)
}
