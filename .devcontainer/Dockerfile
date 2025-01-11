FROM golang:1.23-bullseye as golang
FROM node:22-bullseye as node

FROM mcr.microsoft.com/vscode/devcontainers/base:bullseye

# nodejs
COPY --from=node /usr/local/bin/node /usr/local/bin/node

# golang
COPY --from=golang /usr/local/go /usr/local/go


ENV PATH $PATH:/usr/local/go/bin/