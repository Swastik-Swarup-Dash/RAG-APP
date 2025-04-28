package pinecone

import (
	"context"

	"log"
	"os"

	pineconego "github.com/pinecone-io/go-pinecone/v3/pinecone" // Proper alias usage
)

var (
	Client *pineconego.Client      // Use alias for type
	Index  *pineconego.IndexConnection // Use alias for type
)

func InitPinecone() error {
	clientParams := pineconego.NewClientParams{ // Use alias
		ApiKey: os.Getenv("PINECONE_API_KEY"),
	}
	
	pc, err := pineconego.NewClient(clientParams) // Use alias
	if err != nil {
		return err
	}
	Client = pc

	indexName := os.Getenv("PINECONE_INDEX_NAME")
	dimension := int32(768)
	metric := pineconego.Cosine // Use alias
	cloud := pineconego.Aws     // Use alias
	region := "us-east-1"

	// Check if index exists
	idx, err := Client.DescribeIndex(context.Background(), indexName)
	if err != nil {
		log.Printf("Creating new serverless index: %s", indexName)
		
		// Create serverless index
		_, err = Client.CreateServerlessIndex(context.Background(), &pineconego.CreateServerlessIndexRequest{
			Name:      indexName,
			Dimension: &dimension,
			Metric:    &metric,
			Cloud:     cloud,
			Region:    region,
		})
		if err != nil {
			return err
		}

		// Refresh index description
		idx, err = Client.DescribeIndex(context.Background(), indexName)
		if err != nil {
			return err
		}
	}

	// Connect to index
	Index, err = Client.Index(pineconego.NewIndexConnParams{ // Use alias
		Host: idx.Host,
	})
	return err
}
