#!/usr/bin/env bash

docker run \
    --rm \
    -it \
    -v $(pwd)/scripts/.pgpass:/root/.pgpass \
    --link postgres_local:postgredb \
    --network backend_default \
    postgres:15-alpine \
    sh -c "psql -h postgredb -U cmsuser --no-password cms"