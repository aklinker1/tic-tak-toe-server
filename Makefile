build:
	docker build . -t aklinker1/tic-tak-toe-server:dev
run: build
	docker-compose up
gen:
	swagger generate server -t gen -f api/swagger.yml --exclude-main -A tic-tak-toe

# Aliases
b: build
r: run
