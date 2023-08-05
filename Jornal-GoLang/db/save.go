package db

import (
	"context"
	"fmt"

)

func Insert(collection string, data interface{}) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("golang_jornal").Collection(collection)

	_, err := c.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("erro ao inserir dados: %v", err)
	}

	return nil
}
