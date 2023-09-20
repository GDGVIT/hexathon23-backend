@echo off
REM Vitty
REM Handy set of commands to run to get a new server up and running
if "%1"=="local" (
    shift
    set "environment=local"
    set "file=docker-compose-local.yaml"
) else (
    set "environment=production"
    set "file=docker-compose-prod.yaml"
)
set "command=%1"

if "%command%"=="" (
    echo.
    echo "    ██╗  ██╗███████╗██╗  ██╗"
    echo "    ██║  ██║██╔════╝╚██╗██╔╝"
    echo "    ███████║█████╗   ╚███╔╝ "
    echo "    ██╔══██║██╔══╝   ██╔██╗ "
    echo "    ██║  ██║███████╗██╔╝ ██╗"
    echo "    ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝"
    echo.
    echo Environment: %environment%
    echo.
    echo Author: Dhruv Shah
    echo.
    echo Usage: hex [env] [command]
    echo.
    echo Available environments:
    echo   local: Local development environment
    echo   production: Production environment (default)
    echo.
    echo Available commands:
    echo   up: Start the server
    echo   down: Stop the server
    echo   restart: Restart the server
    echo   recreatedb: Recreate all the tables in database
    echo   cli: Run a command inside the container
    exit /b 1
)

REM Start server command
if "%command%"=="up" (
    echo Starting server
    docker compose -f "%file%" up -d --build
    exit /b 1
)

REM Stop server command
if "%command%"=="down" (
    echo Stopping server
    docker compose -f "%file%" down
    exit /b 1
)

REM Restart server command
if "%command%"=="restart" (
    echo Restarting server
    docker compose -f "%file%" down
    docker compose -f "%file%" up -d --build
    exit /b 1
)

REM Drop all DB tables
if "%command%"=="recreatedb" (
    echo Dropping all tables
    docker compose -f "%file%" run --rm postgres psql -U postgres -d postgres -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
    echo Restarting server
    docker compose -f "%file%" up -d --build
    exit /b 1
)

REM Management commands
if "%command%"=="cli" (
    shift
    docker compose -f "%file%" run --rm vitty-api ./bin/hex-api %*
    exit /b 1
)
