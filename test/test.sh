#!/bin/bash

script_dir=$(dirname "$0")

source ${script_dir}/test.env

INPUT_FILE="test.mp4"
OUTPUT_FILE="test-out.mp4"

${script_dir}/../build/uploader s3 \
    --source="${script_dir}/assets/${INPUT_FILE}" \
    --dest="test/${OUTPUT_FILE}"

${script_dir}/../build/uploader download \
    --url="${S3_URL}/test/${OUTPUT_FILE}" \
    --dest="${script_dir}/assets/${OUTPUT_FILE}"
