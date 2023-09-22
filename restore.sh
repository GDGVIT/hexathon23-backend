#!/bin/bash

# Check if a backup file is provided as an argument
if [ $# -ne 1 ]; then
  echo "Usage: $0 <backup_file>"
  exit 1
fi

BACKUP_FILE=$1

# Check if the backup file exists
if [ ! -f "$BACKUP_FILE" ]; then
  echo "Backup file '$BACKUP_FILE' does not exist."
  exit 1
fi

# Drop all the existing tables from the database
docker exec -i hexathon-postgres psql -U postgres postgres << EOF
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
EOF

# Restore the database from the backup file inside the PostgreSQL container
docker exec -i hexathon-postgres psql -U postgres postgres < $BACKUP_FILE

# Check the exit status of the psql command
if [ $? -eq 0 ]; then
  echo "Database successfully restored from '$BACKUP_FILE'."
else
  echo "Error: Failed to restore database from '$BACKUP_FILE'."
fi