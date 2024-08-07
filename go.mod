module github.com/2Nemanja/RSSAggregator

go 1.22.4

require github.com/joho/godotenv v1.5.1 // indirect -- go get plus the gitlink gets this so i can get data from .env --

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/cors v1.2.1
)
