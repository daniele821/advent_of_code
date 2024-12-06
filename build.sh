#!/bin/bash

SCRIPT_PWD="$(realpath "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(dirname "${SCRIPT_PWD}")"
TEMPLATES="${SCRIPT_DIR}/template"

! [[ -d "${TEMPLATES}" ]] && echo 'templates not exist!' && exit 1

echo -n 'write year: '
read -r year
! [[ "${year::2}" == "20" && "${year}" -gt 2000 ]] && echo 'invalid year!' && exit 1
echo -n 'write day: '
read -r day
! [[ "${day}" -gt 0 && "${day}" -lt 26 ]] && echo 'invalid day!' && exit 1

DIR="${year}/${day}"
echo -n "conferm you want to create: $DIR"
read -r answer
[[ "${answer,,}" == "y" ]] && exit 0

mkdir -p "${SCRIPT_DIR}/${year}"
cp -r "${TEMPLATES}" "${SCRIPT_DIR}/${DIR}"

exit 0
