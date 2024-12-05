#!/bin/bash

SCRIPT_PWD="$(realpath "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(dirname "${SCRIPT_PWD}")"

echo 'ADVENT OF CODE:' >/dev/tty
for dir in "${SCRIPT_DIR}/20"*; do
    echo "Year $(basename "${dir}"):" >/dev/tty
    [[ -d "${dir}" ]] && for subdir in "${dir}/"*; do
        if [[ -d "${subdir}" ]]; then
            echo "Day $(basename "${subdir}"):" >/dev/tty
            cd "${subdir}" || exit 1
            go run ./main.go
            echo >/dev/tty
        fi
    done
done
