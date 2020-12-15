.PHONY: build clear-cache test

build:
	echo "Building zd-query binary"
	(cd src && go build -o ../zd-query)
	
test:
	(cd src && go test)

clear-cache:
	echo "Removing cache files"
	(cd .cache && rm *.json)