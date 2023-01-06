# Apito | Back-end

## Environment Variables

`apito-backend` needs that the listed environment variables in `.env-example` to be available to correctly run.

## Database

A database named in the `$DATABASE_URL` environment variable should be available. 

## Endpoints

### Referees

| HTTP method | Path                  | Status Code      | Description                      |
|-------------|-----------------------|------------------|----------------------------------|
| GET         | /api/v1/referees      | 200 (OK)         | Fetches all Referees resources.  |
| POST        | /api/v1/referees      | 200 (CREATED)    | Create a new Referee resource.   |
| GET         | /api/v1/referees/{id} | 200 (OK)         | Fetch a single Referee resource. |
| PUT         | /api/v1/referees/{id} | 200 (OK)         | Updates a Referee resource.      |
| DELETE      | /api/v1/referees/{id} | 204 (No content) | Deletes a Referee resource.      |

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


