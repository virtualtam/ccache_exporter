#!/bin/bash -ex
#
# Generate sample ccache output testdata for a given Debian or Ubuntu version.
#
# This script:
#
# - creates a Docker container
# - outputs ccache data with an empty cache
# - runs a first build to populate the cache, then outputs ccache data
# - runs a second build to reuse the cache, then outputs ccache data
#
# Usage
#
#   ./run.sh <docker image> <docker image tag> <container name>
#
# Example
#
#  ./run.sh local/ccache debian-11 ccache-testdata-debian-11

# Docker image
IMAGE=$1

# Docker image tag
TAG=$2

# Docker container name
NAME=$3

function _version {
    # Convert a version code (up to 4 digits) to a integer for version
    # comparisons.
    #
    # Reference:
    # - https://stackoverflow.com/questions/4023830/how-to-compare-two-strings-in-dot-separated-version-format-in-bash
    # - https://apple.stackexchange.com/questions/83939/compare-multi-digit-version-numbers-in-bash/123408#123408
    echo "$@" | awk -F. '{ printf("%d%03d%03d%03d\n", $1,$2,$3,$4); }'
}

function _print_stats() {
    # Print ccache stats to a local directory.
    local output_dir=$1
    local version=$2
    local name=$3

    # ccache < 3.7:  machine-readable stats
    # ccache >= 3.7: human-readable stats
    docker exec $NAME ccache --show-stats > $output_dir/$name

    # ccache >= 3.7: machine-readable stats (tab-separated values)
    # https://ccache.dev/releasenotes.html#_ccache_3_7
    if [[ $(_version $version) -ge $(_version "3.7") ]]
    then
        docker exec $NAME ccache --print-stats > $output_dir/$name.tsv
    fi
}

# Start ccache container
docker run --rm --name $NAME -d -it $IMAGE:$TAG

# Retrieve ccache version code
docker exec $NAME ccache --version
VERSION=$(docker exec $NAME bash -c 'ccache --version | grep "ccache version" | cut -d " " -f 3')

OUTPUT_DIR=output/$TAG-ccache-$VERSION
mkdir -p $OUTPUT_DIR

# Clear cache and reset stats
docker exec $NAME ccache --clear --zero-stats
_print_stats $OUTPUT_DIR $VERSION empty

# First build
docker exec $NAME apt source --compile ccache
_print_stats $OUTPUT_DIR $VERSION firstbuild

# Second build
docker exec $NAME apt source --compile ccache
_print_stats $OUTPUT_DIR $VERSION secondbuild

# Stop and remove ccache container
docker stop $NAME
