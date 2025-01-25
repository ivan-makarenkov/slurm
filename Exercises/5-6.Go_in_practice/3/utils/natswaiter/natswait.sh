#!/bin/sh
# wait.sh

set -e

host="$1"
shift
  
until ./natswaiter -u "$host"; do
  >&2 echo "NATS is unavailable - sleeping"
  sleep 1
done
  
>&2 echo "NATS is up - executing command"
exec "$@"