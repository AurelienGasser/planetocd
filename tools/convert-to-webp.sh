#!/bin/bash

convert() {
    EXT=$1
    for i in $(find . -name '*'${EXT}); do
        OUTPUT_PATH=${i/$EXT/.webp}
        if ! [ -f "${OUTPUT_PATH}" ];then
            echo $OUTPUT_PATH
            cwebp $i -o ${OUTPUT_PATH}
        fi
    done
}

convert .png
convert .jpg

