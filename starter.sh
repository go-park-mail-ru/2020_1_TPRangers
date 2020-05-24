#!/usr/bin/env bash
runner_func_authorization() {
  echo "cd to $1"
  # shellcheck disable=SC2164
  cd $1
  go build authorization.go
  nohup ./authorization > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi  
}

runner_func_chats() {
  echo "cd to $1"
  # shellcheck disable=SC2164
  cd $1
  go build chats.go
  nohup ./chats > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_likes() {
  echo "cd to $1"
  # shellcheck disable=SC2164
  cd $1
  go build likes.go
  nohup ./likes > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_photos() {
  echo "cd to $1"
  # shellcheck disable=SC2164
  cd $1
  go build photos.go
  nohup ./photos > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_photo_save() {
  echo "cd to $1"
  # shellcheck disable=SC2164
  cd $1
  go build photo_save.go
  nohup ./photo_save > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_main() {
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

runner_func_authorization "internal/cmd/authorization"
runner_func_chats "internal/cmd/chats/"
runner_func_likes "internal/cmd/likes/"
runner_func_photos "internal/cmd/photos/"
runner_func_photo_save "internal/cmd/photo_save"
runner_func_main "."




echo "all services are running"