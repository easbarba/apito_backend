# Apito | Back-end

## Installation

Get all dependencies and install with:

    $ make deps && make build 

### Container images

[Docker Hub](https://hub.docker.com/r/easbarbosa/apito)

### SQL
    All SQL related tasks files are in `ops/sql`, you easily run those with the target prefixed by `db` in the `Makefile`.

    `make dbclean`

PS: Database engine will expect its password to be set. `PGPASSWORD=meh123`

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)


