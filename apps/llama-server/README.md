# llama-server

Upstream [llama.cpp](https://github.com/ggml-org/llama.cpp) `llama-server`, built with CUDA for llama.cpp's default architectures, plus the patches in [`patches/`](patches). The image holds only the binary at `/llama-server`. Tested only on the GPUs and models in my [llama-swap config](https://git.xcd.dev/gabrielcosi/home-ops/src/branch/main/kubernetes/apps/ai/llama-swap/app/config.yaml); the patches can change with any update. [Issues](https://github.com/gabrielcosi/containers/issues) welcome.

## Patches

- `0001`: `GGML_CUDA_MMVQ_MAX_NVFP4=<n>` sets the largest batch size that uses the MMVQ kernel for NVFP4.
- `0002`: applies the NVFP4 scales to the gate and FFN weights of Muse Glimmer models.
- `0003`: disk prompt cache, below.
- `0004`: moves context checkpoints into the prompt cache instead of copying them.
- `0005`: for recurrent and hybrid models with speculative decoding such as MTP (not EAGLE3), takes the last prompt checkpoint without decoding the end of the prompt in a separate batch.

## Disk prompt cache

States evicted from the RAM cache go to disk instead of being dropped, and the server writes the rest on a clean shutdown. After a restart, it loads them back instead of processing the prompt again.

- Enable it with `--cache-dir` and `--cache-disk` (see `--help`). It needs the RAM cache on.
- Use a persistent folder per server. Files the server can't use are deleted.
- A llama.cpp update, a model or context change, or a cache format change starts the disk cache empty.
