# Apito | backend

## Environment Variables

`apito-backend` needs that the listed environment variables in `.env-example` to be available to correctly run.

## Database

A database named in the `$DATABASE_URL` environment variable should be available. 

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
