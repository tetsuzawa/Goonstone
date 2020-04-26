#!/usr/bin/env bash
# Check argument count
if [ $# -ne 2 ]; then
  echo "ARGS: $#" 1>&2
  echo "Error: Require argument=1" 1>&2
  exit 1
fi

cd $(dirname $(cd $(dirname $0); pwd))

# AWS Login
$(aws ecr get-login --no-include-email)

# Build
ECR_REPOSITORY_URL_FRONTEND=$1 ECR_REPOSITORY_URL_API=$2 docker-compose -f docker-compose.prod.yml build --no-cache
docker push "$ECR_REPOSITORY_URL_FRONTEND":latest
docker push "$ECR_REPOSITORY_URL_API":latest

exit 0
