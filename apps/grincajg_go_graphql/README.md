## migrations

```
migrate -path database/migration/ -database "postgresql://postgres:postgres@localhost:5432/grincajg_graphql_dev?sslmode=disable" -verbose up
```

## create migrations

```
migrate create -ext sql -dir database/migration/ -seq create-category
migrate -path database/migration/ -database "postgresql://postgres:postgres@localhost:5432/grincajg_graphql_dev?sslmode=disable" down 1

```

## deploy docker

1. This command creates a new builder instance and sets it as the current builder.

```

docker buildx create --use
```

2. This command inspects the current builder instance and bootstraps it, ensuring that it's ready to use.

```

docker buildx inspect --bootstrap
```

3.This command builds and pushes a multi-architecture image to Docker Hub. The --platform flag specifies the target platforms for the build (in this case, linux/amd64 and linux/arm64).

```
docker buildx build --platform linux/amd64,linux/arm64 -t dinodanic/grincajg_go_graphql:latest . --push
```

If you've previously set up a builder instance, you can use just the last command to build and push your image. Otherwise, you should run all three commands to ensure that everything is set up correctly.
