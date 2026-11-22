FROM docker.io/nvidia/cuda:13.5.0-devel-ubuntu24.04 AS build

ARG VERSION

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates cmake git python3 python3-venv \
    && rm -rf /var/lib/apt/lists/*

ADD --keep-git-dir=true https://github.com/jesdga95/basalt.git#${VERSION} /app

WORKDIR /app

RUN cmake -S . -B build -DCMAKE_BUILD_TYPE=Release \
    && cmake --build build -j"$(nproc)" \
    && python3 -m venv .venv \
    && .venv/bin/pip install --no-cache-dir -r requirements.txt

FROM docker.io/nvidia/cuda:13.5.0-runtime-ubuntu24.04

# Basalt dlopens FFmpeg's shared libraries for video input; it never runs the ffmpeg CLI.
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates python3 libavcodec60 libavformat60 libavutil58 libswscale7 \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir /data \
    && chown 1000:1000 /data

COPY --from=build /app/.venv /app/.venv
COPY --from=build /app/bin /app/bin
COPY --from=build /app/serve /app/serve
COPY --from=build /app/tools /app/tools
COPY --from=build /app/data /app/data
COPY --from=build /app/requirements.txt /app/LICENSE /app/THIRD_PARTY_LICENSES /app/

ENV HOME=/data \
    PATH=/app/bin:$PATH \
    PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

USER 1000:1000

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/bin/basalt"]
