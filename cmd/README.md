# Service Main

## build

```shell
docker build -t edm-service:test -f cmd/Dockerfile.x86_64 .
docker image ls -f 'dangling=true' -q | xargs docker image rm
```
