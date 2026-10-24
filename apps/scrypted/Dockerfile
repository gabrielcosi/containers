ARG VERSION
FROM ghcr.io/koush/scrypted:${VERSION}

USER root

# Upstream's third-party apt repositories can go offline (coral-edgetpu returns 403); catatonit only needs the Ubuntu archive.
RUN apt-get update -o Dir::Etc::SourceParts=- -o Dir::Etc::SourceList=/etc/apt/sources.list.d/ubuntu.sources \
    && apt-get install -y --no-install-recommends catatonit \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && chown 1000:1000 /server/volume \
    && ln -s /tmp/.exit /server/.exit \
    && ln -s /tmp/.update /server/.update

COPY entrypoint.sh /entrypoint.sh

USER 1000:1000

ENTRYPOINT ["/usr/bin/catatonit", "-g", "--", "/entrypoint.sh"]
