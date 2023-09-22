#!/bin/bash

# Set the current date for the backup filename
BACKUP_FILENAME="backup_$(date +%Y%m%d%H%M%S).sql"

# Check if the backups directory exists
if [ ! -d "./backups" ]; then
  mkdir ./backups
fi


# Run the backup command inside the PostgreSQL container
docker exec -it hexathon-postgres pg_dump -U postgres postgres > ./backups/$BACKUP_FILENAME