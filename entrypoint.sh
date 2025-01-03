#!/bin/sh

set -e

echo "Running database migrations..."
./goose -dir ./db/migrations mysql "$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" up

echo "Starting application..."
exec ./main
