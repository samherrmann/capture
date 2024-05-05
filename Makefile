# Disable CGO since we have no Go code that calls into C.
export CGO_ENABLED := 0
gobuild = mkdir -p dist/$$GOOS-$$GOARCH && go build -ldflags "-s -w" -o dist/$$GOOS-$$GOARCH .
tar = (cd dist && tar -czvf $$GOOS-$$GOARCH.tar.gz $$GOOS-$$GOARCH/*)
zip = (cd dist && zip -r $$GOOS-$$GOARCH.zip $$GOOS-$$GOARCH/*)

build:
	mkdir -p dist && go build -o dist .

build.all:
	export GOOS=linux; export GOARCH=386; $(gobuild) && $(tar)
	export GOOS=linux; export GOARCH=amd64; $(gobuild) && $(tar)
	export GOOS=windows; export GOARCH=amd64; $(gobuild) && $(zip)
	export GOOS=windows; export GOARCH=386; $(gobuild) && $(zip)

lint:
	staticcheck -checks=all ./...

test:
	CGO_ENABLED=1 go test ./... -race -cover

clean:
	rm -rf dist

# Resources:
# List of available target OSs and architectures:
# https://golang.org/doc/install/source#environment
