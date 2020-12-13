.PHONY: build 

build:
	echo "Building zd-query binary"
	(cd src && go build -o ../zd-query)