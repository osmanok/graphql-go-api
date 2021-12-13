package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Println("Graphql Tutorial")

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	// defines the object config
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	// defines a schema config
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	// creates schema
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new Graphql Schema err %v", err)
	}

	query := `
		{
			hello
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, erros %+v", r.Errors)
	}

	rJson, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJson)

}
