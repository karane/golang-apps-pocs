#!/bin/sh
set -e

DB_FILE="/app/data/urls.db"

echo "Running migration..."
sqlite3 "$DB_FILE" < /app/migrations/init.sql

echo "Starting app..."
./url-shortener
