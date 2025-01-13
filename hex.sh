#!/bin/sh

# Hexathon
# Handy set of commands to run to get a new server up and running
if [ "$1" = "local" ]; then
    shift # Discard the first argument
    environment="local"
    file="docker-compose-local.yaml"
else
    environment="production"
    file="docker-compose-prod.yaml"
fi
command=$1

if [ -z "$command" ]; then
    echo
    echo "    ██╗  ██╗███████╗██╗  ██╗"
    echo "    ██║  ██║██╔════╝╚██╗██╔╝"
    echo "    ███████║█████╗   ╚███╔╝ "
    echo "    ██╔══██║██╔══╝   ██╔██╗ "
    echo "    ██║  ██║███████╗██╔╝ ██╗"
    echo "    ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝"
    echo        
    echo "Environment: $environment"
    echo
    echo "Author: Dhruv Shah"
    echo
    echo "Usage: hex [env] [command]"
    echo
    echo "Available environments:"
    echo "  local: Local development environment (default)"
    echo "  prod: Production environment"
    echo
    echo "Available commands:"
    echo "  up: Start the server"
    echo "  down: Stop the server"
    echo "  restart: Restart the server"
    echo "  recreatedb: Recreate all the tables in database"
    echo "  copy: Copy a file to the hexathon api container"
    echo "  cli: Run a command inside the container"
    exit 1
fi

# Start server command
if [ "$command" = "up" ]; then
    echo "Starting server"
    docker compose -f "$file" up -d --build
    exit 1
fi

# Stop server command
if [ "$command" = "down" ]; then
    echo "Stopping server"
    docker compose -f "$file" down
    exit 1
fi

# Restart server command
if [ "$command" = "restart" ]; then
    echo "Restarting server"
    docker compose -f "$file" down
    docker compose -f "$file" up -d --build
    exit 1
fi

# Drop all DB tables
if [ "$command" = "recreatedb" ]; then
    echo "Dropping all tables"
    docker exec -it hexathon-postgres psql -U postgres -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
    echo "Restarting server"
    docker compose -f "$file" up -d --build
    exit 1
fi

# Copy file to hexathon api container
if [ "$command" = "copy" ]; then
    shift # Discard the first argument
    docker cp "$1" hexathon-api:/"$1"
    exit 1
fi

# Backup database
if [ "$command" = "backup" ]; then
    echo "Backing up database"
    shift # Discard the first argument
    # Make sure argument is provided
    if [ -z "$1" ]; then
        echo "Usage: hex backup <backup_name>"
        exit 1
    fi
    ./backup.sh "$1"
    exit 1
fi

# Restore database
if [ "$command" = "restore" ]; then
    echo "Restoring database"
    shift # Discard the first argument
    # Make sure argument is provided
    if [ -z "$1" ]; then
        echo "Usage: hex restore <backup_file>"
        exit 1
    fi
    ./restore.sh "$1"
    exit 1
fi

# Management commands
if [ "$command" = "cli" ]; then
    shift # Discard the first argument
    docker compose -f "$file" run --rm api ./bin/hex-api "$@"
    exit 1
fi

if [ "$command" = "psql" ]; then
    docker compose -f "$file" exec -it postgres psql -U postgres postgres
    exit 1
fi


docker compose -f "$file" "$@"