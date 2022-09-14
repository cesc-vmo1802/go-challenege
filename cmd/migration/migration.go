package migration

import (
	"context"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-challenege/common"
	"go-challenege/pkg/database"
	"go-challenege/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const DbUri = "db-uri"

var (
	userCollection = "users"
	appCollection  = "applications"
)
var jsonSchema = bson.M{
	"bsonType":             "object",
	"required":             []string{"name", "enabled", "type"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"name": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
			"minLength":   1,
		},
		"description": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
			"minLength":   1,
			"maxLength":   150,
		},
		"enabled": bson.M{
			"bsonType":    "bool",
			"description": "must be a boolean and is required",
		},
		"type": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
	},
}

var jsonSchemaUser = bson.M{
	"bsonType":             "object",
	"required":             []string{"login_id", "password", "salt", "refresh_token_id"},
	"additionalProperties": true,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"login_id": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
			"minLength":   10,
			"maxLength":   50,
		},
		"password": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
			"minLength":   10,
		},
		"salt": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
			"minLength":   10,
		},
		"refresh_token_id": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
			"minLength":   10,
		},
	},
}

type MgoMigration struct {
	Directory  string `json:"directory"`
	Conn       *mongo.Client
	Database   string
	Collection string
}

func NewMgoMigration(dir string, client *mongo.Client, dbName string) *MgoMigration {
	return &MgoMigration{
		Directory: dir,
		Conn:      client,
		Database:  dbName,
	}
}

func (m *MgoMigration) ExtractData() ([]bson.M, error) {
	var results []bson.M
	err := filepath.Walk(m.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		jsonFile, err := os.Open(path)
		if err != nil {
			return err
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var data []bson.M
		err = json.Unmarshal([]byte(byteValue), &data)
		if err != nil {
			return err
		}

		results = append(results, data...)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m *MgoMigration) Exec() error {
	data, err := m.ExtractData()
	if err != nil {
		return err
	}

	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}

	db := m.Conn.Database(m.Database)
	err = db.Collection("applications").Drop(context.Background())
	if err != nil {
		return err
	}
	opts := options.CreateCollection().SetValidator(validator)

	if err := db.CreateCollection(context.Background(), appCollection, opts); err == nil {
		connection := db.Collection("applications")

		var ui []interface{}
		for _, t := range data {
			ui = append(ui, t)
		}

		_, err = connection.InsertMany(context.Background(), ui)

		if err != nil {
			return err
		}
	}

	validator = bson.M{
		"$jsonSchema": jsonSchemaUser,
	}

	opts = options.CreateCollection().SetValidator(validator)

	if err := db.CreateCollection(context.Background(), userCollection, opts); err != nil {
		log.Println(err)
		return nil
	}

	return nil
}

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command used to migrate database",
	Long:  "This command used to migrate database",
}

func newMigrateUpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "This command used to migrate database",
		Long:  "This command used to migrate database",
		RunE: func(cmd *cobra.Command, args []string) error {
			var l = logger.Init(
				logger.WithLogDir("logs/"),
				logger.WithDebug(true),
				logger.WithConsole(true),
			)
			defer l.Sync()

			dbCnf := database.MongoConfig{
				Uri: viper.GetString(DbUri),
			}

			appMongo := database.NewAppDB(&dbCnf)
			ctx := context.Background()
			if err := appMongo.Start(ctx); err != nil {
				panic(err)
			}

			migration := NewMgoMigration("./dbmigration", appMongo.GetDB(), common.DefaultDatabase)

			if err := migration.Exec(); err != nil {
				panic(err)
			}

			appMongo.Stop(ctx)

			return nil
		},
	}
	return cmd
}

func RegisterCommandRecursive(parent *cobra.Command) {
	cmd := newMigrateUpCommand()
	migrationCmd.AddCommand(cmd)
	parent.AddCommand(migrationCmd)
}
