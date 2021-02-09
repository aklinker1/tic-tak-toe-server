build:
	docker build . -t aklinker1/tic-tak-toe-server:dev
run: build
	docker-compose up
gen:
	mkdir -p package/server/gen
	swagger generate server -t package/server/gen -f api/swagger.yml --exclude-main -A tic-tak-toe

# Aliases
b: build
r: run
