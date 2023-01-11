# Apito | Back-end

## Environment Variables

`apito-backend` needs that the listed environment variables in `.env-example` to be available to correctly run.

## Database

A database named in the `$DATABASE_URL` environment variable should be available.

## Endpoints

### Referees

| method | Pattern        | Code | Action                  |
| ------ | -------------- | ---- | ----------------------- |
| GET    | /referees      | 200  | Fetches all Referees.   |
| POST   | /referees      | 201  | Create a new Referee.   |
| GET    | /referees/{id} | 200  | Fetch a single Referee. |
| PUT    | /referees/{id} | 200  | Updates a Referee.      |
| DELETE | /referees/{id} | 204  | Deletes a Referee       |

- Default port is at `:5000`
- Prefix: `/api/v1`

## Installation

## Installation

Get all dependencies and install with:

    $ make deps && make build

## Ops

### Insomnia

Insomnia tasks are available to easy reproducibility of the API endpoints, the
latest files are at the ops folder.

### OpenAPI

API specification is generated at every release and can placed at the ops folder.

A Swagger UI is also available:

    $ make spec

PS: To generate the correct current api version, do export an environment variable `$APITO_BACKEND_VERSION` returning `.version` value.

### Container images

[Docker Hub](https://hub.docker.com/r/easbarbosa/apito)

### SQL

All SQL related tasks files are in `ops/sql`, you easily run those with the target prefixed by `db` in the `Makefile`.

    `make dbclean`

PS: Database engine will expect its password to be set. `PGPASSWORD=meh123`

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
