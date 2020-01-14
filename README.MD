### if golang wire already installed skip this step
go get github.com/google/wire/cmd/wire
and ensuring that $GOPATH/bin is added to your $PATH.

### Run Containerized Postgres if have no postgres installed
docker run --name postgres -p 5432:5432 -d postgres

## Run migration 
go run cmd/migrate/main.go

## Run Application 
go run cmd/api/main.go cmd/api/wire_gen.go
