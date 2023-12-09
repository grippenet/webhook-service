package webhookdb

import (
	"context"
	"time"

	"github.com/coneno/logger"
	"github.com/grippenet/webhook-service/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageDBService struct {
	DBClient     *mongo.Client
	timeout      int
	DBNamePrefix string
}

func NewMessageDBService(configs types.DBConfig) *MessageDBService {
	var err error
	dbClient, err := mongo.NewClient(
		options.Client().ApplyURI(configs.URI),
		options.Client().SetMaxConnIdleTime(time.Duration(configs.IdleConnTimeout)*time.Second),
		options.Client().SetMaxPoolSize(configs.MaxPoolSize),
	)
	if err != nil {
		logger.Error.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configs.Timeout)*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		logger.Error.Fatal(err)
	}

	ctx, conCancel := context.WithTimeout(context.Background(), time.Duration(configs.Timeout)*time.Second)
	err = dbClient.Ping(ctx, nil)
	defer conCancel()
	if err != nil {
		logger.Error.Fatal("fail to connect to DB: " + err.Error())
	}

	return &MessageDBService{
		DBClient:     dbClient,
		timeout:      configs.Timeout,
		DBNamePrefix: configs.DBNamePrefix,
	}
}

const MESSSAGE_DB_NAME = "messageDB"
const COLLECTION_WEBHOOK = "webhooks"

func (dbService *MessageDBService) GetDBName(instanceID string) string {
	return dbService.DBNamePrefix + instanceID + "_" + MESSSAGE_DB_NAME
}

// Collections
func (dbService *MessageDBService) CollectionRefWebhook(instanceID string) *mongo.Collection {
	return dbService.DBClient.Database().Collection(COLLECTION_WEBHOOK)
}

// DB utils
func (dbService *MessageDBService) GetContext() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(dbService.timeout)*time.Second)
}
