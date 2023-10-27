proto:
	protoc ./internal/proto/*.proto --go_out=. --go-grpc_out=.
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1079 -d postgres:15
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres my_osm_database

dropdb:
	docker exec -it postgres dropdb my_osm_database
