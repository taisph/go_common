#!/usr/bin/env sh

set -eu

module=$(grep -m1 -E "^module .+" go.mod | cut -d " " -f 2-)
exec goimports -local "${module}" "$@"
