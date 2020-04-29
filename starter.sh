#!/bin/bash

runner_func() {
  echo "cd to $1"
  # shellcheck disable=SC2164
  cd $1
  go build main.go
  nohup ./main > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi  
}

# shellcheck disable=SC2140
# shellcheck disable=SC1083
for services in "." "internal/cmd/authorization/" "internal/cmd/chats/" "internal/cmd/likes/" "internal/cmd/photos/"; do
  runner_func "$services" 
done



echo "all services are running"
