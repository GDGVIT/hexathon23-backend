#!/bin/bash

# Check if a backup file name is provided as an argument
if [ $# -ne 1 ]; then
  echo "Usage: $0 <backup_name>"
  exit 1
fi

# Set the current date and time along with input argument name as the backup file name
BACKUP_FILENAME="$1-$(date '+%Y-%m-%d-%H-%M-%S').sql"

# Check if the backups directory exists
if [ ! -d "./backups" ]; then
  mkdir ./backups
fi


# Run the backup command inside the PostgreSQL container
docker exec -it hexathon-postgres pg_dump -U postgres postgres > ./backups/$BACKUP_FILENAME