#!/bin/bash

DATABASE_CONNECTION="postgres://postgres:postgres@localhost:5432/chirpy"
SCHEMA_DIR="sql/schema"

execute_migration() {
  local command=$1
  local success_message=$2

  if ! cd "$SCHEMA_DIR" 2>/dev/null; then
    echo "Error: Could not change to schema directory: $SCHEMA_DIR"
    exit 1
  fi

  local output status
  output=$(set -o pipefail; goose postgres "$DATABASE_CONNECTION" "$command" 2>&1)
  status=$?

  cd - >/dev/null

  if [ $status -eq 0 ]; then
    echo "$success_message"
    echo "$output"
    echo
    return 0
  else
    echo "Migration failed with status code $status"
    echo "$output"
    echo
    return 1
  fi
}

show_help() {
  cat << EOF
  Usage: $(basename "$0") <command>

  Commands:
    up      - Migrate up one version
    down    - Migrate down one version
    upall   - Migrate up to latest version
    downall - Migrate down to version 0
    help    - Show this help message

EOF
}

if [ $# -eq 0 ]; then
  echo "Error: please provide a command"
  show_help
  exit 1
fi

case "$1" in
  "up")
    execute_migration "up-by-one" "Successfully migrated up one version"
    ;;
  "down")
    execute_migration "down" "Successfully migrated down one version"
    ;;
  "upall")
    execute_migration "up" "Successfully migrated up to latest version"
    ;;
  "downall")
    execute_migration "down-to 0" "Successfully migrated down to version 0"
    ;;
  "help")
    show_help
    exit 0
    ;;
  *)
    echo "Error: Unknown command '$1'"
    show_help
    exit 1
    ;;
esac
