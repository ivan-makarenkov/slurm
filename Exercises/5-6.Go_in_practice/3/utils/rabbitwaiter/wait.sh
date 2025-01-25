#!/bin/sh
# wait.sh

set -e

host="$1"
shift
  
until ./rabbitwaiter -u "amqp://guest:guest@$host/"; do
  >&2 echo "RabbitMQ is unavailable - sleeping"
  sleep 1
done
  
>&2 echo "RabbitMQ is up - executing command"
exec "$@"