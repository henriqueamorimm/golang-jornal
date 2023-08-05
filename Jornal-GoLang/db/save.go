package db

func Insert(collection string, data interface{}) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("webform").Collection(collection) 

	_, err := c.InsertOne(ctx, data) 

	return err
}
