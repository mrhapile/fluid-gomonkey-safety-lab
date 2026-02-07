package unsafe_test

import (
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fluid-gomonkey-safety-lab/unsafe"
)

var _ = Describe("Unsafe Pattern: Global State Patching", func() {
	// WHY THIS IS UNSAFE:
	// 1. **Concurrent Modification**: Ginkgo allows tests to run in parallel (e.g., via goroutines or separate processes if configured).
	//    Even with separate processes, if test execution is interleaved or state is shared (e.g. in older versions or specific setups),
	//    gomonkey patches modify the running process memory globally. Two concurrent tests trying to patch the same
	//    global variable will race, causing undefined behavior where one test sees the other's patch.
	// 2. **State Pollution**: If a test fails or panics before cleanup, global state remains modified, affecting subsequent tests.
	//
	// Without the 'Serial' decorator, these tests are prone to flakiness if run in parallel or if execution order changes.

	It("Scenario A: Expects Config 'Alpha'", func() {
		// UNSAFE: Patching global variable without enforcing serial execution.
		patches := gomonkey.ApplyGlobalVar(&unsafe.GlobalConfig, "Alpha")
		defer patches.Reset()

		Expect(unsafe.GetConfig()).To(Equal("Alpha"))
	})

	It("Scenario B: Expects Config 'Beta'", func() {
		// UNSAFE: Patching global variable without enforcing serial execution.
		patches := gomonkey.ApplyGlobalVar(&unsafe.GlobalConfig, "Beta")
		defer patches.Reset()

		Expect(unsafe.GetConfig()).To(Equal("Beta"))
	})
})
