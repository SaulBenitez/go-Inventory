#!/bin/bash
docker run -d \
    --name mariadb-learn \
    -p 127.0.0.1:33060:3306 \
    -e "MARIADB_USER=mariadb" \
    -e "MARIADB_PASSWORD=mariadbpass" \
    -e "MARIADB_ROOT_PASSWORD=mariadbpass" \
    mariadb:latest