DATE = formatdate( "YYYY.MM.DD", timestamp() )
APP = "basalt"
SOURCE = "https://github.com/jesdga95/basalt"
variable "GIT_SHA" {}

variable "VERSION" {
  // renovate: datasource=github-releases depName=jesdga95/basalt versioning=regex:^v(?<major>\d+)\.(?<minor>\d+)\.(?<patch>\d+)$
  default = "v0.6.0"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  args = {
    VERSION = "${VERSION}"
  }
  labels = {
    "org.opencontainers.image.source" = "https://github.com/gabrielcosi/containers"
    "org.opencontainers.image.created" = "${DATE}"
    "org.opencontainers.image.revision" = "${GIT_SHA}"
    "org.opencontainers.image.title" = "${APP}"
    "org.opencontainers.image.url" = "${SOURCE}"
    "org.opencontainers.image.version" = "${VERSION}"
  }
  no-cache = true
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
  tags = ["${APP}:${VERSION}"]
}

target "image-all" {
  inherits = ["image"]
  platforms = ["linux/amd64"]
  tags = [
    "ghcr.io/gabrielcosi/containers/${APP}:rolling",
    "ghcr.io/gabrielcosi/containers/${APP}:sha-${GIT_SHA}",
    "ghcr.io/gabrielcosi/containers/${APP}:${VERSION}",
  ]
}

target "docker-metadata-action" {}
