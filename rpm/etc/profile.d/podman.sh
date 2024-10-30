# podman cli の通信先を root の podman デーモンに切り替える
[[ ! -r /run/podman.sock ]] || export CONTAINER_HOST=unix:/run/podman.sock
