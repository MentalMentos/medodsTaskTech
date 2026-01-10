#!/bin/bash
set -e

# ЯВНО указываем правильную базу данных
DSN="host=pg port=5432 dbname=postgres user=user password=1234 sslmode=disable"

echo "Starting migrations..."
echo "Using DSN: $DSN"
echo "Migration directory: ./migrations"

sleep 10  # Ждём PostgreSQL
goose -dir ./migrations postgres "$DSN" up -v
