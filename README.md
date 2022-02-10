## Run postgres
docker run --name postgres -e POSTGRES_USER=dieptv -e POSTGRES_PASSWORD=dieptv -p 5432:5432 -v /home/dieptv/data:/var/lib/postgres/data -d postgres:latest# worker-transaction
