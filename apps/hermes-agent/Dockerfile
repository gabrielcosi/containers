ARG VERSION
FROM docker.io/nousresearch/hermes-agent:${VERSION}

# renovate: datasource=github-tags depName=eleboucher/memini
ARG MEMINI_VERSION=0.7.39

USER root:root

RUN git clone --depth 1 --branch "v${MEMINI_VERSION}" https://github.com/eleboucher/memini /tmp/memini-src \
    && cp -a /tmp/memini-src/integrations/hermes/plugin/memini /opt/hermes/plugins/memory/memini \
    && rm -rf /tmp/memini-src

COPY plugins/crawl4ai /opt/hermes/plugins/web/crawl4ai
COPY defaults/mise.sh /etc/profile.d/mise.sh

RUN groupmod -g 1000 hermes \
    && usermod -u 1000 -g 1000 hermes \
    && chown 1000:1000 /opt/data

USER 1000:1000
