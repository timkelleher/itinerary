package db

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/timkelleher/itinerary/pkg/env"
	models "github.com/timkelleher/itinerary/pkg/models"
)

type DB struct {
	client   *sqlx.DB
	server   string
	user     string
	password string
	database string
}

var db *DB

func NewDB() (*DB, error) {
	if db == nil {
		db = &DB{
			client:   nil,
			server:   env.DbHost,
			user:     env.DbUsername,
			password: env.DbPassword,
			database: env.DbDatabase,
		}
		if err := db.Connect(); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func (db *DB) Connect() error {
	if db.client != nil {
		return nil
	}

	c, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true",
		db.user,
		db.password,
		db.server,
		db.database,
	))
	if err != nil {
		log.WithFields(log.Fields{
			"server": db.server,
			"user":   db.user,
			"db":     db.database,
			"error":  err.Error(),
		}).Error("Connect to MySQL")
		return err
	}

	log.WithFields(log.Fields{
		"server": db.server,
		"user":   db.user,
		"db":     db.database,
	}).Debug("Connect to MySQL")
	db.client = c
	return nil
}

// Used for tests
func (db *DB) ResetDB() error {
	tables := []string{"sets"}
	for _, table := range tables {
		sql := fmt.Sprintf("DELETE FROM `%s`", table)
		_, err := db.client.Exec(sql)
		if err != nil {
			log.WithFields(log.Fields{
				"sql":   sql,
				"error": err.Error(),
			}).Error("Reset database")
			return err
		}
	}

	log.Debug("Reset database")
	return nil
}

func (db *DB) Collection(id uint) (*models.Collection, error) {
	var collection models.Collection
	sql := "SELECT * FROM collections WHERE id = ? LIMIT 1"
	err := db.client.QueryRowx(sql, id).StructScan(&collection)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"id":    collection.ID,
			"error": err.Error(),
		}).Error("Get Collection")
		return nil, err
	}

	// Grab associated endpoints (without Outputs)
	var endpoints []models.Endpoint
	sql = "SELECT id, collection_id, name, method, path, response_status_code FROM endpoints WHERE collection_id = ?"
	err = db.client.Select(&endpoints, sql, collection.ID)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"id":    collection.ID,
			"error": err.Error(),
		}).Error("Get Collection")
		return nil, err
	}
	collection.Endpoints = endpoints

	log.WithFields(log.Fields{
		"id": collection.ID,
	}).Debug("Get Collection")
	return &collection, nil
}

type CollectionsFilter struct {
}

func (db *DB) Collections(filters CollectionsFilter) ([]*models.Collection, error) {
	sql := "SELECT * FROM collections"

	var collections []*models.Collection
	err := db.client.Select(&collections, sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"error": err.Error(),
		}).Error("Get Collections")
		return nil, err
	}

	log.Debug("Get Collections")
	return collections, nil
}

func (db *DB) InsertCollection(c *models.Collection) error {
	sql := "INSERT INTO collections (name, path_prefix, content_type) VALUES (?, ?, ?)"
	res, err := db.client.Exec(sql, c.Name, c.PathPrefix, c.ContentType)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"error": err.Error(),
		}).Error("Insert Collection")
		return err
	}
	c.ID, err = res.LastInsertId()
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"error": err.Error(),
		}).Error("Insert Collection")
		return err
	}

	log.WithFields(log.Fields{
		"id":   c.ID,
		"name": c.Name,
	}).Debug("Insert Collection")
	return nil
}

func (db *DB) UpdateCollection(c *models.Collection) error {
	sql := "UPDATE collections" +
		" SET" +
		" 	name = ?," +
		" 	path_prefix = ?," +
		" 	content_type = ?" +
		" WHERE id = ? LIMIT 1"
	_, err := db.client.Exec(
		sql,
		c.Name,
		c.PathPrefix,
		c.ContentType,
		c.ID,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"id":    c.ID,
			"error": err.Error(),
		}).Error("Update Collection")
	}

	return err
}

// updateCollections
// make sure to update computed path for all endpoints

func (db *DB) Endpoint(id uint) (*models.Endpoint, error) {
	var endpoint models.Endpoint
	sql := "SELECT endpoints.*, collections.path_prefix" +
		" FROM endpoints" +
		" JOIN collections ON endpoints.collection_id = collections.id" +
		" WHERE endpoints.id = ? LIMIT 1"
	err := db.client.QueryRowx(sql, id).StructScan(&endpoint)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":          sql,
			"attempted_id": id,
			"error":        err.Error(),
		}).Error("Get Endpoint")
		return nil, err
	}

	// Fetch Collection record next
	endpoint.Collection, err = db.Collection(uint(endpoint.CollectionID))
	if err != nil {
		log.WithFields(log.Fields{
			"sql":                     sql,
			"attempted_collection_id": endpoint.CollectionID,
			"error":                   err.Error(),
		}).Error("Get Endpoint")
		return nil, err
	}
	log.WithFields(log.Fields{
		"id": endpoint.ID,
	}).Debug("Get Endpoint")

	return &endpoint, nil
}

func (db *DB) EndpointFromPath(path, method string) (*models.Endpoint, error) {
	var sqlParams []interface{}
	sqlParams = append(sqlParams, path)
	requireMethodClause := ""
	method = strings.ToUpper(method)
	if method != "ANY" {
		sqlParams = append(sqlParams, method)
		requireMethodClause = " AND endpoints.method = ?"
	}

	var endpoint models.Endpoint
	sql := "SELECT endpoints.*, collections.path_prefix" +
		" FROM endpoints" +
		" JOIN collections ON endpoints.collection_id = collections.id" +
		" WHERE endpoints.computed_path = ?" +
		requireMethodClause +
		" LIMIT 1"
	err := db.client.QueryRowx(sql, sqlParams...).StructScan(&endpoint)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":              sql,
			"attempted_path":   path,
			"attempted_method": method,
			"error":            err.Error(),
		}).Error("Get Endpoint From Path")
		return nil, err
	}

	// Fetch Collection record next
	endpoint.Collection, err = db.Collection(uint(endpoint.CollectionID))
	if err != nil {
		log.WithFields(log.Fields{
			"sql":                     sql,
			"attempted_collection_id": endpoint.CollectionID,
			"error":                   err.Error(),
		}).Error("Get Endpoint From Path")
		return nil, err
	}

	log.WithFields(log.Fields{
		"id":   endpoint.ID,
		"path": endpoint.ComputePath(),
	}).Debug("Get Endpoint From Path")
	return &endpoint, nil
}

func (db *DB) InsertEndpoint(e models.Endpoint) error {
	sql := "INSERT INTO endpoints (name, collection_id, method, response_status_code, path, computed_path, output)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?)"
	res, err := db.client.Exec(
		sql,
		e.Name,
		e.CollectionID,
		e.Method,
		e.ResponseStatusCode,
		e.Path,
		e.ComputePath(),
		e.Output,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"id":    e.ID,
			"error": err.Error(),
		}).Error("Create Endpoint")
	}

	e.ID, err = res.LastInsertId()
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"error": err.Error(),
		}).Error("Insert Endpoint")
		return err
	}

	log.WithFields(log.Fields{
		"id":   e.ID,
		"name": e.Name,
	}).Debug("Insert Endpoint")
	return nil
}

func (db *DB) UpdateEndpoint(endpoint models.Endpoint) error {
	sql := "UPDATE endpoints" +
		" SET" +
		" 	name = ?," +
		" 	method = ?," +
		" 	response_status_code = ?," +
		" 	path = ?," +
		" 	computed_path = ?," +
		" 	output = ?" +
		" WHERE id = ? LIMIT 1"
	_, err := db.client.Exec(
		sql,
		endpoint.Name,
		endpoint.Method,
		endpoint.ResponseStatusCode,
		endpoint.Path,
		endpoint.ComputePath(),
		endpoint.Output,
		endpoint.ID,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":   sql,
			"id":    endpoint.ID,
			"error": err.Error(),
		}).Error("Update Endpoint")
	}

	return err
}
