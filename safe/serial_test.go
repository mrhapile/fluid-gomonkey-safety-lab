package safe_test

import (
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fluid-gomonkey-safety-lab/safe"
)

// Mark as Serial so Ginkgo never runs these in parallel with other specs in this suite.
// This ensures that global state modifications are exclusive to one test at a time.
var _ = Describe("Safe Pattern: Serial Execution", Serial, func() {
	It("Scenario A: Expects Config 'Alpha'", func() {
		// SAFE: Patching global variable safely within a serial spec.
		patches := gomonkey.ApplyGlobalVar(&safe.GlobalConfig, "Alpha")
		defer patches.Reset()

		Expect(safe.GetConfig()).To(Equal("Alpha"))
	})

	It("Scenario B: Expects Config 'Beta'", func() {
		// SAFE: Patching global variable safely within a serial spec.
		patches := gomonkey.ApplyGlobalVar(&safe.GlobalConfig, "Beta")
		defer patches.Reset()

		Expect(safe.GetConfig()).To(Equal("Beta"))
	})
})
