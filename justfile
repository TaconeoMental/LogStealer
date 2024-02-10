_default:
    @just --list --unsorted

HERE        := justfile_directory()
PLUGINS_DIR := HERE / "handlers"
PLUGINS_SRC := PLUGINS_DIR / "src"
PLUGINS_LIB := PLUGINS_DIR / "build" # Compiled plugins

clean:
    @echo "Deleting shared libraries"
    -rm {{PLUGINS_LIB}}/*

build-handlers: clean
    #!/usr/bin/env bash
    set -euo pipefail

    find {{PLUGINS_SRC}} -type f -name "*.go" | while read path; do
        #src_filename="${path##*/}"
        #filename2="${path%.*}.so"
        src_filename=$(basename $path)
        so_filename=$(basename $src_filename .go).so
        echo "Compiling $src_filename -> $so_filename"
        go build -o {{PLUGINS_LIB}}/$so_filename -buildmode=plugin $path
    done
