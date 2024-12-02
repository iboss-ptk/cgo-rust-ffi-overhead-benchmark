
# --------------------------------------------------------
# Rust Builder
# --------------------------------------------------------
FROM --platform=linux/arm64 rust:1.82.0 as rustbuilder

WORKDIR /ffibench

COPY . .

RUN cargo build --release --manifest-path=rust/Cargo.toml

# --------------------------------------------------------
# Go Builder
# --------------------------------------------------------

FROM --platform=linux/arm64 golang:1.22 as gobuilder

COPY --from=rustbuilder /ffibench /ffibench
WORKDIR /ffibench

ENTRYPOINT [ "/bin/bash" ]

RUN go test -c

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ubuntu:25.04
COPY --from=gobuilder /ffibench/ffibench.test /bin/ffibench.test

ENTRYPOINT ["/bin/ffibench.test"]
