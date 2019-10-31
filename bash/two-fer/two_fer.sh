#!/usr/bin/env bash

set -o errexit
set -o nounset

main() {
  local name="${1:-you}"
  printf "%s" "One for $name, one for me."
  printf "\n"
}

main "$@"
