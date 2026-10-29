FROM docker.io/nvidia/cuda:13.4.3-devel-ubuntu24.04 AS build

ARG VERSION

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        cmake \
        git \
        ninja-build \
    && rm -rf /var/lib/apt/lists/*

ADD --keep-git-dir=true https://github.com/ggml-org/llama.cpp.git#${VERSION} /src
COPY patches/ /patches/

WORKDIR /src

RUN for p in /patches/*.patch; do git apply --verbose "$p" || exit 1; done

RUN cmake -S . -B build -G Ninja \
        -DCMAKE_BUILD_TYPE=Release \
        -DGGML_CUDA=ON \
        -DGGML_STATIC=ON \
        -DGGML_CUDA_NCCL=OFF \
        -DGGML_NATIVE=OFF \
        -DBUILD_SHARED_LIBS=OFF \
        -DLLAMA_CURL=OFF \
        -DLLAMA_OPENSSL=OFF \
        -DLLAMA_BUILD_TESTS=OFF \
        -DLLAMA_BUILD_EXAMPLES=OFF \
        "-DCMAKE_EXE_LINKER_FLAGS=-Wl,-rpath-link,/usr/local/cuda/lib64/stubs -lcuda" \
    && cmake --build build -j"$(nproc)" --target llama-server \
    && test -x build/bin/llama-server

FROM docker.io/library/busybox:1.38.0@sha256:a46257e06c82ad85242d0d38077d0a52d9896389ca60a1a59d3463db586e5fa4
COPY --from=build /src/build/bin/llama-server /llama-server
