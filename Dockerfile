FROM debian:stretch-slim
WORKDIR /
COPY app /usr/local/bin
CMD ["app"]