package mongodb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

type MongoOpts struct {
	Uri        string
	DB         string
	Collection string
	SSLConfig  *SSLConf
}

type SSLConf struct {
	CAFile   string
	CertFile string
	KeyFile  string
}

func New(opt MongoOpts) (*mongo.Collection, error) {
	ctx := context.TODO()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(opt.Uri).SetServerAPIOptions(serverAPI)

	if opt.SSLConfig != nil {
		// Loads CA certificate file
		caCert, err := os.ReadFile(opt.SSLConfig.CAFile)
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			return nil, fmt.Errorf("error: CA file must be in PEM format")
		}
		// Loads client certificate files
		cert, err := tls.LoadX509KeyPair(opt.SSLConfig.CertFile, opt.SSLConfig.KeyFile)
		if err != nil {
			return nil, err
		}
		// Instantiates a Config instance
		tlsConfig := &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{cert},
		}
		opts.SetTLSConfig(tlsConfig)
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database(opt.DB).Collection(opt.Collection), nil
}
