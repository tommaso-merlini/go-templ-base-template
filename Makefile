run: build
	@./bin/app #set the project name here

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	# @go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest
	@npx tailwindcss init -p

build:
	npx tailwindcss -i view/css/app.css -o public/styles.css 
	@templ generate view
	@go build -o bin/app main.go 

up: ## Database migration up
	@go run cmd/migrate/main.go up

reset:
	@go run cmd/reset/main.go up

down: ## Database migration down
	@go run cmd/migrate/main.go down

migration: ## Migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go
