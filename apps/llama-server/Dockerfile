FROM docker.io/nvidia/cuda:13.4.1-devel-ubuntu24.04 AS build

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
        -DGGML_STATIC=ON \
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

FROM docker.io/library/busybox:1.38.0@sha256:fd7dc98638c8e305f4dc34e979f1c0fdfdcaeb0fbf8fcff77ae834b6da3d7e6e
COPY --from=build /src/build/bin/llama-server /llama-server
