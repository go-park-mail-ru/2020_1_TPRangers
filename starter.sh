#!/bin/bash

runner_func_authorization() {
  echo "authorization"
  # shellcheck disable=SC2164
  cd $1

  nohup ./authorization > author.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi  
}

runner_func_chats() {
  echo "chats"
  # shellcheck disable=SC2164
  cd $1

  nohup ./chats > chats.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_likes() {
  echo "likes"
  # shellcheck disable=SC2164
  cd $1
  nohup ./likes > likes.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_photos() {
  echo "photos"
  # shellcheck disable=SC2164
  cd $1
  nohup ./photos > photos.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_photo_save() {
  echo "photo_save"
  # shellcheck disable=SC2164
  cd $1
  nohup ./photo_save > photo_save.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

runner_func_main() {
  echo "main"
  # shellcheck disable=SC2164
  cd $1
  nohup ./main > server.log &
  if [[ "$1" != "." ]]; then
    cd ../../../
  fi
}

# shellcheck disable=SC2140
# shellcheck disable=SC1083

runner_func_authorization "."
runner_func_chats "."
runner_func_likes "."
runner_func_photos "."
runner_func_photo_save "."
runner_func_main "."




echo "all services are running"
