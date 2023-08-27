#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER dynus ENCRYPTED PASSWORD 'dynus' LOGIN;
	CREATE DATABASE dynus OWNER dynus;
EOSQL

psql -v ON_ERROR_STOP=1 --username "dynus" --dbname "dynus" -f /app/sql/init-db.sql
