# Presentation

This is a repository for a demo application.
The demo demonstrates implementing a simple reconciliation loop.

## Demo

Start docker

```
docker run -d --name postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres:11
docker exec -it postgres psql -U postgres
```

List PostgreSQL databases

```
watch 'psql -U postgres -h localhost -c "SELECT datname from pg_database"'
```

List active connections

```
postgres=# SELECT * FROM pg_stat_activity;
```