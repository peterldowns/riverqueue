# How to run the tests

```bash
# start a postgres server on localhost port 5432
docker compose up -d
# export the test database url
export TEST_DATABASE_URL="postgres://postgres:postgres@127.0.0.1:5432/river_testdb?sslmode=disable"
# create the test databases
PGHOST=127.0.0.1 PGPORT=5432 PGUSER=postgres PGPASSWORD=postgres PGSSLMODE=disable go run ./internal/cmd/testdbman create
# Run the tests -- takes about 13 seconds uncached
# If you add -race -count=1 to all the test commands,
# takes about:
#
#   make test  70.00s user 9.86s system 78% cpu 1:41.22 total
#
time make test
```
