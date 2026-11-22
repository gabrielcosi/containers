package main

import (
	"context"
	"testing"

	"github.com/gabrielcosi/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/gabrielcosi/containers/basalt:rolling")

	t.Run("basalt serve answers /health on port 8080", func(t *testing.T) {
		testhelpers.TestCommandSucceeds(t, ctx, image, nil, "/app/.venv/bin/python", "-c", `
import subprocess, sys, time, urllib.request
server = subprocess.Popen(["/app/bin/basalt", "serve", "--mock-engine", "--host", "0.0.0.0", "--port", "8080"])
try:
    deadline = time.monotonic() + 30
    while time.monotonic() < deadline:
        if server.poll() is not None:
            sys.exit(1)
        try:
            with urllib.request.urlopen("http://127.0.0.1:8080/health", timeout=2) as r:
                sys.exit(0 if r.status == 200 else 1)
        except OSError:
            time.sleep(0.5)
    sys.exit(1)
finally:
    server.terminate()
    server.wait(timeout=10)
`)
	})

	t.Run("Runs as a non-root user with a writable HOME", func(t *testing.T) {
		testhelpers.TestCommandSucceeds(t, ctx, image, nil, "sh", "-c", `test "$(id -u)" -ne 0 && mkdir -p "$HOME/.cache/basalt"`)
	})

	t.Run("FFmpeg libraries load the way basalt loads them", func(t *testing.T) {
		testhelpers.TestCommandSucceeds(t, ctx, image, nil, "/app/.venv/bin/python", "-c", `
import ctypes, sys
# The rows mirror ff::kAbiRows in include/basalt/vision/ffmpeg_abi.hpp at the pinned tag.
for fmt, codec, util, scale in ((63, 63, 61, 10), (62, 62, 60, 9), (61, 61, 59, 8), (60, 60, 58, 7)):
    try:
        avformat = ctypes.CDLL(f"libavformat.so.{fmt}")
    except OSError:
        continue
    libs = [avformat] + [ctypes.CDLL(f"lib{name}.so.{major}") for name, major in (("avcodec", codec), ("avutil", util), ("swscale", scale))]
    versions = [libs[0].avformat_version(), libs[1].avcodec_version(), libs[2].avutil_version(), libs[3].swscale_version()]
    sys.exit(0 if [v >> 16 for v in versions] == [fmt, codec, util, scale] else 1)
sys.exit(1)
`)
	})
}
