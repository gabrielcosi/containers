ARG VERSION
FROM docker.io/lmsysorg/sglang:${VERSION}

COPY patches/ /patches/

RUN for p in /patches/*.patch; do git -C /sgl-workspace/sglang apply --verbose "$p" || exit 1; done

USER 1000:1000
