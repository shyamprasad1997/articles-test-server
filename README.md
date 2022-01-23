# White rabbit test

### Steps to run the migration:
* Setup your credentials in config/local.env
* Source config/local.env
* Run `go run *main*.go` from db_migration folder

### Steps to run the application:
* Setup your credentials in config/local.env
* Source config/local.env
* Run `go run main.go` from root
* Server will be available on port `8080`

### Endpoints:
* `PUT :8080/v1/article/decline/{article id}` - decline articles
* `PUT :8080/v1/article/approve/{article id}` - approve articles
* `GET :8080/v1/articles/search/{page}?key={key to search}` - search articles
* `POST :8080/v1/article` - post articles (data goes in body, sample json:`{"title": "title","description": "my new description"}`)
* `GET :8080/v1/articles/{page}` - get all articles
