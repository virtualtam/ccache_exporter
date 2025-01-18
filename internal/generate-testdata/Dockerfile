# Copyright (c) VirtualTam
# SPDX-License-Identifier: MIT

# Generic Dockerfile for Debian- and Ubuntu-based images
ARG DOCKER_IMAGE
ARG DOCKER_TAG
FROM ${DOCKER_IMAGE}:${DOCKER_TAG}

ENV DEBIAN_FRONTEND=noninteractive

# Enable ccache system-wide for non-interactive logins (docker exec):
# - https://askubuntu.com/questions/470545/how-do-i-set-up-ccache
# - https://forums.debian.net/viewtopic.php?t=53053
ENV PATH="/usr/lib/ccache:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
ENV CCACHE_DIR="/home/cached/.ccache"

# Optionally configure a remote storage:
# - https://github.com/ccache/ccache/blob/master/doc/MANUAL.adoc#config_remote_storage
# - https://github.com/ccache/ccache/blob/master/doc/MANUAL.adoc#config_remote_only
ENV CCACHE_REMOTE_ONLY=""
ENV CCACHE_REMOTE_STORAGE=""

# Enable source repositories and install build dependencies to compile the
# ccache package ;-)
ADD enable-deb-src /usr/local/bin/

RUN enable-deb-src \
    && apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y \
        ccache \
        dpkg-dev \
        gcc \
        git \
    && update-ccache-symlinks \
    && apt build-dep -y ccache

# Create an unprivileged user to run compilation and cache operations
RUN groupadd --gid 1010 cached \
    && useradd \
        --create-home \
        --shell /bin/bash \
        --uid 1010 \
        --gid cached \
        cached

USER cached
WORKDIR /home/cached

# Pre-fetch the source package for ccache
RUN apt-get source ccache
