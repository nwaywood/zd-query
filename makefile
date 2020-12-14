.PHONY: build clear-cache

build:
	echo "Building zd-query binary"
	(cd src && go build -o ../zd-query)
	
clear-cache:
	echo "Removing cache files"
	(cd cache && rm *.json)