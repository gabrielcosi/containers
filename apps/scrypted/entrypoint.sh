#!/usr/bin/env bash

ulimit -c 0

exec npm --prefix /server exec scrypted-serve "$@"
