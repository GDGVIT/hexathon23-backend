#!/bin/sh

# Vitty
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
    echo "  local: Local development environment"
    echo "  production: Production environment (default)"
    echo
    echo "Available commands:"
    echo "  up: Start the server"
    echo "  down: Stop the server"
    echo "  restart: Restart the server"
    echo "  recreatedb: Recreate all the tables in database"
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

# Management commands
if [ "$command" = "cli" ]; then
    shift # Discard the first argument
    docker compose -f "$file" run --rm hexathon-api ./bin/hex-api "$@"
    exit 1
fi