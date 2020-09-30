package main

import (
	"fmt"
	"time"

	"github.com/mateeullahmalik/go-schema-registry/srclient"
)

func main() {
	opts := srclient.SROpts{
		Timeout: 20 * time.Second,
		URL:     "http://data-infra-schema-reg-elb-dev-1632556799.ap-southeast-1.elb.amazonaws.com:8081",
	}
	sr := srclient.NewSchemaRegistryClient(opts)
	schema, err := sr.GetLatestSchema("demo-user-dev-value")
	if err != nil {
		panic(err)
	}

	fmt.Println(schema)
}
