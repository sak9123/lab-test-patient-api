#!/bin/sh

# https://github.com/docker-library/postgres/issues/151

set -x
POSTGRES="psql --username ${POSTGRES_USER}"

echo "Before"
echo "======"
$POSTGRES <<-SQL
\du
SQL

echo -n "[*] Creating database role: ${DB_USER}... "
$POSTGRES <<-SQL
CREATE USER ${DB_USER} WITH CREATEDB PASSWORD '${DB_PASSWORD}';
SQL

echo -n "[*] Creating database ${DB_NAME}... "
$POSTGRES <<-SQL
CREATE DATABASE ${DB_NAME} OWNER ${DB_USER};
SQL

echo
echo "After"
echo "====="
$POSTGRES <<-SQL
\du
SQL