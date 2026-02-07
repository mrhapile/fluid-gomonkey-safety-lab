# fluid-gomonkey-safety-lab

A Proof-of-Work repository demonstrating the risks and correct usage of `gomonkey` in Ginkgo-based test suites.

## Problem: Global State Patching

In microservices and Kubernetes controllers (like Fluid), applications often rely on global runtime state for configuration, caching, or feature flags. Common examples include:
- Global configuration maps (e.g., `pkg/ddc` in Fluid).
- Singleton clients or factories.
- Time providers.

When writing unit tests for functions that depend on this global state, developers often turn to monkey patching libraries like `gomonkey` to swap out the dependency at runtime. While powerful, `gomonkey` modifies the executable code in memory, affecting the entire process.

## Unsafe Pattern (Why Tests Become Flaky)

If tests patch global variables or functions without explicit serialization, they become inherently unsafe when run in parallel.

In the `unsafe/` directory, the tests do not use the `Serial` decorator. If these tests are run in parallel (e.g., `ginkgo -p` or via concurrent goroutines in a single process), the following occurs:
1.  **Race Condition**: Test A patches `GlobalConfig` to "Alpha". Test B patches it to "Beta". Since they share the same memory space (if in the same process) or if the patches overlap in time, one test will see the wrong value.
2.  **State Pollution**: If a test fails or panics before calling `patches.Reset()`, the global state remains modified, causing subsequent tests to fail mysteriously.

This pattern is a major source of "flakiness" in CI pipelines.

## Safe Pattern (Serial Execution)

The correct way to handle global state patching in Ginkgo is to force serial execution for the affected specs.

In the `safe/` directory, we use the `Serial` decorator on the `Describe` block:
```go
var _ = Describe("Safe Pattern", Serial, func() { ... })
```


This ensures that:
1.  **Exclusive Access**: Ginkgo guarantees that no other specs in the suite run concurrently with these tests.
2.  **Deterministic State**: We can safely patch and reset globals without fear of interference from other tests.

## Execution Flow Comparison

### Without Serial (Unsafe)
Parallel execution leads to race conditions where multiple tests attempt to patch the same global memory simultaneously.

```text
Test A ─┐
        ├── modifies GlobalConfig (RACE CONDITION)
Test B ─┘
```

### With Serial (Safe)
Sequential execution ensures that each test has exclusive access to the global state, applies its patch, runs assertions, and resets the state before the next test begins.

```text
Test A ──▶ Patch GlobalConfig ──▶ Assert ──▶ Reset
             (Exclusive Access)

Test B ──▶ Patch GlobalConfig ──▶ Assert ──▶ Reset
             (Exclusive Access)
```


## Relevance to Fluid (pkg/ddc)

Fluid's `pkg/ddc` (Deep Learning Data Caching) logic relies heavily on global configurations and runtime checks. As we modernize the test suite to use Ginkgo, it is critical to identify all tests using `gomonkey` on global state and mark them as `Serial`. Failing to do so will result in a fragile CI pipeline where tests pass locally but fail randomly under load.

## Running the Tests

To run the unsafe suite (failures may depend on parallelism settings):
```bash
ginkgo -r ./unsafe
```

To run the safe suite:
```bash
ginkgo -r ./safe
```
