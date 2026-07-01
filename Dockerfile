# SimpleScript needs the Go toolchain at RUNTIME, not just at build time: the
# `run`, `build`, and `test` commands shell out to `go` to compile your program.
# So this image keeps Go in the final layer on purpose — a slim "scratch" image
# without Go could only do `check` and `reveal`. This is the honest trade-off
# for a transpiler-CLI rather than a self-contained service.

FROM golang:1.22

WORKDIR /src
COPY . .

# Install the friendly command onto PATH.
RUN go build -o /usr/local/bin/simplescript ./cmd/simplescript

# A scratch workspace for your .ss files; mount your own with -v.
WORKDIR /work

ENTRYPOINT ["simplescript"]
CMD ["help"]

# Usage:
#   docker build -t simplescript .
#   docker run --rm -v "$PWD":/work simplescript run hello.ss
#   docker run --rm -v "$PWD":/work simplescript check --af hello.ss
