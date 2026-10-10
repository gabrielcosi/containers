FROM docker.io/library/golang:1.27.2-trixie@sha256:e58d6f83b3416618d8bcac2b3dde1b7f7e3c4a77d25e88637f8bbae81536c48d AS build

ARG VERSION

ADD --keep-git-dir=true https://github.com/mostlygeek/llama-swap.git#${VERSION} /src
COPY patches/ /patches/

WORKDIR /src

RUN for p in /patches/*.patch; do git apply --verbose "$p" || exit 1; done

RUN CGO_ENABLED=0 go build -trimpath -ldflags "-X main.version=${VERSION#v}" -o /kubeswap ./cmd/kubeswap

FROM docker.io/library/busybox:1.38.0@sha256:fd7dc98638c8e305f4dc34e979f1c0fdfdcaeb0fbf8fcff77ae834b6da3d7e6e
COPY --from=build /kubeswap /kubeswap
