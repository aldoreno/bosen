#!/usr/bin/env bash

docker run \
    --rm \
    -it \
    -v $(pwd)/scripts/.pgpass:/root/.pgpass \
    -v $(pwd)/migrations/postgresql:/tmp \
    --link postgres_local:postgredb \
    --network backend_default \
    postgres:15-alpine \
    sh -c "psql -h postgredb -U cmsuser cms < /tmp/001_initial_tables_definition.sql"