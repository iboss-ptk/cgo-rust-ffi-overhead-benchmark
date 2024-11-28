use criterion::{criterion_group, criterion_main, Criterion};
use rustffi::{add, fibonacci};
use std::hint::black_box;

fn criterion_benchmark(c: &mut Criterion) {
    c.bench_function("add 1 + 2", |b| b.iter(|| add(black_box(1), black_box(2))));
    c.bench_function("fib 100000", |b| b.iter(|| fibonacci(black_box(100000))));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
