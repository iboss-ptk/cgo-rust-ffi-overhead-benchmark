
# --------------------------------------------------------
# Rust Builder
# --------------------------------------------------------
FROM --platform=linux/arm64 rust:1.82.0 as rustbuilder

WORKDIR /ffibench

COPY rust rust

RUN cargo build --release --manifest-path=rust/Cargo.toml

# --------------------------------------------------------
# Go Builder
# --------------------------------------------------------

FROM --platform=linux/arm64 golang:1.22 as gobuilder


WORKDIR /ffibench
COPY . . 

COPY --from=rustbuilder /ffibench/rust/target/release/librustffi.a /ffibench/rust/target/release/librustffi.a
COPY --from=rustbuilder /ffibench/rust/target/release/librustffi.h /ffibench/rust/target/release/librustffi.h


ENTRYPOINT [ "/bin/bash" ]

RUN go test -c

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ubuntu:25.04
COPY --from=gobuilder /ffibench/ffibench.test /bin/ffibench.test

ENTRYPOINT ["/bin/ffibench.test"]
