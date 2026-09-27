# CUDA major.minor must match the CUDA runtime in the llama-swap image that runs this binary
FROM docker.io/nvidia/cuda:13.3.1-devel-ubuntu24.04 AS build

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

RUN for p in /patches/*.patch; do git apply --verbose "$p"; done

RUN cmake -S . -B build -G Ninja \
        -DCMAKE_BUILD_TYPE=Release \
        -DGGML_CUDA=ON \
        -DCMAKE_CUDA_ARCHITECTURES=120 \
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

FROM docker.io/library/busybox:1.37.0@sha256:bdf57e528e45e4433820e045b29b4597825a1c9e38353532d90a01445013f82e
COPY --from=build /src/build/bin/llama-server /llama-server
