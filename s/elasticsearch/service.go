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
)

var Client *el.Client

// Init initializes the Elasticsearch client
func Init() {
	var err error
	Client, err = el.NewClient(
		el.SetURL(os.Getenv("ELASTIC_URL_1")),
		el.SetSniff(false),
		// el.SetBasicAuth(os.Getenv("ELASTIC_USERNAME"), os.Getenv("ELASTIC_PASSWORD")),
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to elasticsearch: %w", err))
	}
	logrus.Info("Elasticsearch successfully connected")
}

// Insert a new document into an index
func Insert(ctx context.Context, index string, log interface{}) error {
	if Client == nil {
		return fmt.Errorf("elasticsearch client is not initialized")
	}

	_, err := Client.Index().
		Index(index).
		BodyJson(log).
		Do(ctx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot insert data",
			"Index":         index,
			"Data":          log,
		}).Error(err.Error())
		return err
	}

	return nil
}

// Update a document by ID in the index
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
			"ElasticSearch": "cannot update data",
			"ID":            ID,
			"Index":         index,
			"Data":          update,
		}).Error(err.Error())
		return err
	}

	return nil
}

// Search performs a search on a given index
func Search(ctx context.Context, index string, searchSource *el.SearchSource) (*el.SearchResult, error) {
	if Client == nil {
		return nil, fmt.Errorf("elasticsearch client is not initialized")
	}

	results, err := Client.Search().
		Index(index).
		SearchSource(searchSource).
		Do(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot search data",
			"Index":         index,
		}).Error(err.Error())
		return nil, err
	}

	return results, nil
}

// Get placeholder
func Get(ctx context.Context) (interface{}, error) {
	// Implementasi belum tersedia
	return nil, nil
}

// Helper function to get env with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
