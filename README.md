# devops-on-azure-ep-1
Session 1 of DevOps on Azure series

## Dockerfile build
Using podman or docker:
```bash
$ docker build --tag docker.io/doa-ep1:v1 .
$ docker run -dt -p 3000:3000/tcp docker.io/library/doa-ep1:v1
```

Navigate to `localhost:3000`