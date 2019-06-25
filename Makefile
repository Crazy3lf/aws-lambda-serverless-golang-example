.PHONY: build clean deploy gomodgen

build: gomodgen
	set GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/create functions/main.go functions/create.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete functions/main.go functions/delete.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update functions/main.go functions/update.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	bash gomod.sh
