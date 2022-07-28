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

# Enable source repositories and install build dependencies to compile the
# ccache package ;-)
RUN sed -n 'p; s/^deb/deb-src/p' -i /etc/apt/sources.list \
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
RUN groupadd --gid 1000 cached \
    && useradd \
        --create-home \
        --shell /bin/bash \
        --uid 1000 \
        --gid cached \
        cached

USER cached
WORKDIR /home/cached

# Pre-fetch the source package for ccache
RUN apt-get source ccache
