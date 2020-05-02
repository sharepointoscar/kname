package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharepointoscar/kname/themes"
)

var _ = Describe("Theme", func() {

	var (
		_theme = themes.NewTheme("cocktails")
	)

	BeforeEach(func() {

	})

	Describe("Theme name", func() {

		It("Name should contain any of: 'yoga','cocktails' or 'national-parks", func() {
			Expect(_theme.Name).To(Equal("cocktails"))
		})
		It("It should generate cluster name based on theme", func() {

			Expect(_theme.ClusterName).ToNot(BeEmpty())
		})

	})
})
