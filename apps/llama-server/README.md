# llama-server

Upstream [llama.cpp](https://github.com/ggml-org/llama.cpp)'s CUDA `server` image (`.devops/cuda.Dockerfile`) without the web UI, plus the patches in [`patches/`](patches). It runs `/app/llama-server`. Tested only on the GPUs and models in my [LLMKube models](https://git.xcd.dev/gabrielcosi/home-ops/src/branch/main/kubernetes/apps/ai/llmkube/models); the patches can change with any update. [Issues](https://github.com/gabrielcosi/containers/issues) welcome.

## Patches

- `0001`: `GGML_CUDA_MMVQ_MAX_NVFP4=<n>` sets the largest batch size that uses the MMVQ kernel for NVFP4.
- `0002`: disk prompt cache, below.
- `0003`: when the first checkpoint of a request falls on the context checkpoint it just restored, keeps that one instead of saving the same state again.

## Disk prompt cache

States evicted from the RAM cache go to disk instead of being dropped, and the server writes the rest on a clean shutdown. After a restart, it loads them back instead of processing the prompt again.

- Enable it with `--cache-dir` and `--cache-disk` (see `--help`). It needs the RAM cache on.
- Use a persistent folder per server. Files the server can't use are deleted.
- A llama.cpp update, a model or context change, or a cache format change starts the disk cache empty.
