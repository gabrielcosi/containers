FROM docker.io/nvidia/cuda:13.5.0-devel-ubuntu24.04 AS build

ARG VERSION
ARG GCC_VERSION=14

RUN apt-get update && \
    apt-get install -y gcc-${GCC_VERSION} g++-${GCC_VERSION} build-essential cmake python3 python3-pip git libssl-dev libgomp1

ENV CC=gcc-${GCC_VERSION} CXX=g++-${GCC_VERSION} CUDAHOSTCXX=g++-${GCC_VERSION}

ADD --keep-git-dir=true https://github.com/ggml-org/llama.cpp.git#${VERSION} /app
COPY patches/ /patches/

WORKDIR /app

RUN for p in /patches/*.patch; do git apply --verbose "$p" || exit 1; done

RUN cmake -B build -DGGML_NATIVE=OFF -DGGML_CUDA=ON -DGGML_BACKEND_DL=ON -DGGML_CPU_ALL_VARIANTS=ON -DLLAMA_BUILD_TESTS=OFF -DLLAMA_USE_PREBUILT_UI=OFF -DCMAKE_EXE_LINKER_FLAGS=-Wl,--allow-shlib-undefined . && \
    cmake --build build --config Release -j$(nproc)

RUN mkdir -p /app/lib && \
    find build -name "*.so*" -exec cp -P {} /app/lib \;

FROM docker.io/nvidia/cuda:13.5.0-runtime-ubuntu24.04

RUN apt-get update \
    && apt-get install -y libgomp1 curl ffmpeg \
    && apt autoremove -y \
    && apt clean -y \
    && rm -rf /tmp/* /var/tmp/* \
    && find /var/cache/apt/archives /var/lib/apt/lists -not -name lock -type f -delete \
    && find /var/cache -type f -delete

COPY --from=build /app/lib/ /app

ENV LLAMA_ARG_HOST=0.0.0.0
ENV LLAMA_ARG_PORT=8080

COPY --from=build /app/build/bin/llama-server /app/

WORKDIR /app

HEALTHCHECK CMD [ "curl", "-f", "http://localhost:8080/health" ]

ENTRYPOINT [ "/app/llama-server" ]
