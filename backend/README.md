## Overview

Bosen were meant to be a proof of concept of monorepo project.
That being said **backend** is the APIs providing application.
The end goal of this project is to be a reference (or template) project.

This project tries to implements Clean Code e.g. dependency direction are going
inward hence no inner components depends on outer components.

## How to run

1. Create new PostgreSQL instance `docker-compose up -d db`

2. Initialise database with [schema](./migrations/postgresql/).

   Using Docker:

   ```shell scripts/import_db.sh
   docker run \
        --rm \
        -it \
        -v $(pwd)/scripts/.pgpass:/root/.pgpass \
        -v $(pwd)/migrations/postgresql:/tmp \
        --link postgres_local:postgredb \
        --network backend_default \
        postgres:15-alpine \
        sh -c "psql -h postgredb -U cmsuser cms < /tmp/001_initial_tables_definition.sql"
   ```

   NOTE: See [docker-compose.yml](docker-compose.yml) for default database credentials and informations. If you make adjustments, please reflect you changes in [.pgpass](scripts/.pgpass) and `scripts/*.sh`.

3. Duplicate `.env.example` to `.env` then update its values

4. `make run`