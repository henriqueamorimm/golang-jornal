
# Requirements:

1. Download <a href="https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/instalacao-do-go#o-ambiente-go">Go Lang</a>
2. Download <a href="https://www.mongodb.com/try/download/community">Mongo</a>

# Tutorial:
## 🔵First step: Clone the Golang Jornal repository
```
git clone https://github.com/henriqueamorimm/golang-jornal
```
## 🔵Second step: Connect your Mongo URI into the project
```
func getConnection() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("YOUR_URI"))
	if err != nil {
		log.Fatal(err)
	}
```
# Enjoy! ❤


