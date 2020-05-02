package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type YogaTheme struct {
	Name string
}

var _ = Describe("Theme", func() {

	var (
		_theme YogaTheme
	)

	BeforeEach(func() {

		_theme = YogaTheme{
			Name: "yoga",
		}

	})

	Describe("Theme name", func() {

		Context("initially", func() {

			It("It should load based on theme chosen", func() {})

			It("Name should contain any of: 'yoga','cocktails' or 'national-parks", func() {
				Expect(_theme.Name).To(Equal("yoga"))
			})
		})

		Context("It should generate name", func() {

		})

	})
})
