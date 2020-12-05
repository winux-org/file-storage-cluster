docker rm --force pg-db
docker build --tag file-db-pg:1.0 .
docker run --publish 5432:5432 --detach --name pg-db -e POSTGRES_PASSWORD=password  file-db-pg:1.0

