docker_project = be-main-template
docker_version = v1.0.0

sql-download:
	@command -v sqlc > /dev/null | go get github.com/kyleconroy/sqlc/cmd/sqlc

sql-init: sql-download
	sqlc init

sql-gen: sql-download
	sqlc generate

migrate-init:
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_userdt"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_onboarding"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_list_parameter"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_hist_onboarding"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_propinsi"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_kota"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_kecamatan"
	#migrate create -ext sql -dir migration/schema -seq "create-table-mb_kelurahan"

proto-gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./delivery/rpc/todo.proto

local-http:
	go run main.go

test-short:
	go test -short -coverprofile=coverage.out -v -count=1 ./...

test:
	go test -coverprofile=bin/cov.out -v -count=1 ./...

build: test-short
	docker build -t docker.mncbank.co.id:5000/$(docker_project):$(docker_version) -f Dockerfile .

run:
	docker run --rm --env-file ".env" -p 80:9002 --name $(docker_project) docker.mncbank.co.id:5000/$(docker_project):$(docker_version)

push:
	 docker push docker.mncbank.co.id:5000/$(docker_project):$(docker_version)

swag-download:
	@command -v sqlc > /dev/null | go get github.com/swaggo/swag/cmd/swag

swag: swag-download
	swag init