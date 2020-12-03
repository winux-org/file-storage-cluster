FILES=serv.go database.go models.go rpc.go api.go

build:
	go build -o serv $(FILES)

run:
	go run $(FILES) -- env=prod