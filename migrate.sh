#!/bin/bash

ARG="$1"
DATABASE_CONNECTION="postgres://postgres:postgres@localhost:5432/chirpy"

if [ ! $1 ]; then
  echo "please pass in argument"
  exit 1
fi

if [ $ARG == "up" ]; then
  output=$(cd sql/schema && goose postgres $DATABASE_CONNECTION up-by-one 2>&1)
  if [ $? -eq 0 ]; then
    echo "successfully migrated up one"
    echo $output
    echo ""
    exit 0
  else
    echo "migration failed with status code $?"
    echo $output
    echo ""
    exit 1
  fi
fi

if [ $ARG == "down" ]; then
  output=$(cd sql/schema && goose postgres $DATABASE_CONNECTION down 2>&1)
  if [ $? -eq 0 ]; then
    echo "successfully migrated down one"
    echo $output
    echo ""
    exit 0
  else
    echo "migration failed with status code $?"
    echo $output
    echo ""
    exit 1
  fi
fi

if [ $ARG == "upall" ]; then
  output=$(cd sql/schema && goose postgres $DATABASE_CONNECTION up 2>&1)
  if [ $? -eq 0 ]; then
    echo "successfully migrated up all"
    echo $output
    echo ""
    exit 0
  else
    echo "migration failed with status code $?"
    echo $output
    echo ""
    exit 1
  fi
fi

if [ $ARG == "downall" ]; then
  output=$(cd sql/schema && goose postgres $DATABASE_CONNECTION down-to 0 2>&1)
  if [ $? -eq 0 ]; then
    echo "successfully migrated down all"
    echo $output
    echo ""
    exit 0
  else
    echo "migration failed with status code $?"
    echo $output
    echo ""
    exit 1
  fi
fi

if [ $ARG == "help" ]; then
  echo "arguments to use:"
  echo "  * up - migrates up one file"
  echo "  * down - migrates down one file"
  echo "  * upall - migrates up all files"
  echo "  * downall - migrates down all files"
  echo ""
  exit 0
fi

echo "argument doesn't exist"
exit 1
