
run: build
	go run main.go

build: buildServer buildClient
buildServer:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
buildClient:
	BABEL_ENV=production gulp

watch: watchServer watchClient

watchServer:
	gin --port 8085

watchClient:
	watchify frontend/init.jsx -t babelify -p livereactload -o public/js/index.js

setup:
	go get github.com/codegangsta/gin
	go get github.com/golang/lint/golint
	go get
	rm -r ./node_modules
	npm install

test:
	golint ./...
	go tool vet ./
	go test
	flow check

cover:
	go test -coverprofile=coverage.out
	go tool cover  -html=coverage.out
