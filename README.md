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

## OpenAPI

| API name       | Path          |
|----------------|---------------|
| openapi {JSON} | /openapi      |
| openapi {YAML} | /openapi.yaml |
| openapi docs   | /openapi/ui   |

### Container images

[Docker Hub](https://hub.docker.com/r/easbarbosa/apito)

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
