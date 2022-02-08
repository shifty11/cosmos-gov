# Cosmos-Gov - Never miss governance again

## Edit models
See [documentation](https://entgo.io/docs/getting-started)

- Edit schemas in ./ent/schemas
- Run ```go generate ./ent```

## Environment variables
- TELEGRAM_TOKEN used to communicate with the telegram api. Get a new one from [BotFather](https://t.me/BotFather)
- DATABASE_URL connection string for postgreSQL database
- SENTRY_DSN (optional) used to sent logs to [Sentry](https://sentry.io/).
- DEBUG (optional) if true log debug level

## Local development
Run PostgreSQL db in docker.
```bash
docker-compose -f local-db.yml up
```
Go to [localhost:5050](http://localhost:5050) and login as `pgadmin4@pgadmin.org` (password: `admin`) 
and create a new database `cosmosgov`.

