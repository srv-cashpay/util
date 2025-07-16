package elasticsearch

import (
	"context"
	"fmt"
	"os"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var (
	ELASTIC_INDEX_NAME = getEnv("ELASTIC_INDEX_NAME", "default_index")
	Client             *el.Client
)

// Init initializes the Elasticsearch client
func Init() {
	var err error
	Client, err = el.NewClient(
		el.SetURL(os.Getenv("ELASTIC_URL_1")),
		el.SetSniff(false),
		// Uncomment if auth is needed:
		// el.SetBasicAuth(os.Getenv("ELASTIC_USERNAME"), os.Getenv("ELASTIC_PASSWORD")),
	)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to Elasticsearch")
	}
	logrus.Info("Elasticsearch successfully connected")
}

// Insert a new document into the specified index
func Insert(ctx context.Context, index string, log interface{}) error {
	if Client == nil {
		return fmt.Errorf("elasticsearch client is not initialized")
	}

	result, err := Client.Index().
		Index(index).
		Type("_doc").
		BodyJson(log).
		Do(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "Insert Failed",
			"Index":         index,
			"Error":         err,
			"Data":          log,
		}).Error("Elasticsearch insert error")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Index":  result.Index,
		"ID":     result.Id,
		"Result": result.Result,
	}).Info("Document inserted into Elasticsearch")

	return nil
}

// Update an existing document by ID
func Update(ctx context.Context, index, ID string, update map[string]interface{}) error {
	if Client == nil {
		return fmt.Errorf("elasticsearch client is not initialized")
	}

	_, err := Client.Update().
		Index(index).
		Type("_doc").
		Id(ID).
		Doc(update).
		Do(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "Update Failed",
			"ID":            ID,
			"Index":         index,
			"Data":          update,
		}).Error("Elasticsearch update error")
		return err
	}

	logrus.WithField("ID", ID).Info("Document updated in Elasticsearch")
	return nil
}

// Search for documents matching a given query
func Search(ctx context.Context, index string, searchSource *el.SearchSource) (*el.SearchResult, error) {
	if Client == nil {
		return nil, fmt.Errorf("elasticsearch client is not initialized")
	}

	result, err := Client.Search().
		Index(index).
		SearchSource(searchSource).
		Do(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "Search Failed",
			"Index":         index,
			"Error":         err,
		}).Error("Elasticsearch search error")
		return nil, err
	}

	return result, nil
}

// Get is a placeholder for future implementation
func Get(ctx context.Context) (interface{}, error) {
	return nil, nil
}

// Helper function for fallback env values
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
