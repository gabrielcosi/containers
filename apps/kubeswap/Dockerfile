FROM docker.io/library/golang:1.27.2-trixie@sha256:47667934874913a449b96faa7ef4777fa285d9ec29126ee80ed61c45b7cccf0e AS build

ARG VERSION

ADD --keep-git-dir=true https://github.com/mostlygeek/llama-swap.git#${VERSION} /src
COPY patches/ /patches/

WORKDIR /src

RUN for p in /patches/*.patch; do git apply --verbose "$p" || exit 1; done

RUN CGO_ENABLED=0 go build -trimpath -ldflags "-X main.version=${VERSION#v}" -o /kubeswap ./cmd/kubeswap

FROM docker.io/library/busybox:1.38.0@sha256:fd7dc98638c8e305f4dc34e979f1c0fdfdcaeb0fbf8fcff77ae834b6da3d7e6e
COPY --from=build /kubeswap /kubeswap
