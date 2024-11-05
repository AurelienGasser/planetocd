#!/bin/bash

set -ex

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
BASE_DIR="${SCRIPT_DIR}/../server/static/images/illustrations"
OUTPUT_DIR="${BASE_DIR}"
ORIGINALS_DIR="${BASE_DIR}/originals"
TMP_DIR="$(mktemp -d)"

# Create different sizes
magick convert -resize 1024x\> $1 -set filename:f '%t' "${TMP_DIR}/%[filename:f]_lg.webp"
magick convert -resize 500x\> $1 -set filename:f '%t' "${TMP_DIR}/%[filename:f]_md.webp"
magick convert -resize 100x\> $1 -set filename:f '%t' "${TMP_DIR}/%[filename:f]_sm.webp"

# Move to final desination
mv $1 "${ORIGINALS_DIR}" || true
mv "${TMP_DIR}"/* "${OUTPUT_DIR}"

