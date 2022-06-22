# Cosmos-Gov - Never miss governance again

## Environment variables
- TELEGRAM_TOKEN used to communicate with the telegram api. Get a new one from [BotFather](https://t.me/BotFather)
- DATABASE_URL connection string for postgreSQL database
- SENTRY_DSN (optional) used to sent logs to [Sentry](https://sentry.io/).
- DEBUG (optional) if true log debug level 
- WEB_APP_EXTERNAL_URL (telegram only) links to web app (in browser)
- WEB_APP_URL (telegram only) links to web app (in-app)
- TELEGRAM_ENDPOINT (telegram only) ex. https://api.telegram.org/bot%s/test/%s - needed if bot runs on telegram test server
- JWT_SECRET_KEY (grpc only) jwt secret for webapp authentication
- ADMIN_IDS ex: 123523,4523443 - admins have more options in Telegram and in the webapp. Use the telegram user ID as admin ID.

## Local development
Run PostgreSQL db in docker.
```bash
docker-compose -f local-db.yml up
```
Go to [localhost:5050](http://localhost:5050) and login as `pgadmin4@pgadmin.org` (password: `admin`) 
and create a new database `cosmosgov`.

### Run services
```bash
go run main.go fetching   # service that fetches all proposals and chains
go run main.go telegram   # runs telegram bot
go run main.go discord    # runs discord bot
go run main.go grpc       # runs GRPC service for webapp
```

### Run frontend
See https://github.com/shifty11/cosmos-gov-web

## Edit database models
See [documentation](https://entgo.io/docs/getting-started)

- Edit schemas in ./ent/schemas
- Run ```go generate ./ent```