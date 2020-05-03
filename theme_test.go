package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharepointoscar/kname/themes"
)

var _ = Describe("Theme", func() {

	Describe("Yoga Theme", func() {

		var (
			_theme = themes.NewTheme("yoga")
		)

		It("Name should contain any of: yoga", func() {
			Expect(_theme.Name).To(Equal("yoga"))
		})
		It("It should generate cluster name based on theme", func() {

			Expect(_theme.ClusterName).ToNot(BeEmpty())
		})

	})

	Describe("Cocktail Theme", func() {

		var (
			_theme = themes.NewTheme("cocktails")
		)

		It("Name should contain any of: cocktails", func() {
			Expect(_theme.Name).To(Equal("cocktails"))
		})
		It("It should generate cluster name based on theme", func() {

			Expect(_theme.ClusterName).ToNot(BeEmpty())
		})

	})

	Describe("National Parks Theme", func() {

		var (
			_theme = themes.NewTheme("national-parks")
		)

		It("Name should contain any of: cocktails", func() {
			Expect(_theme.Name).To(Equal("national-parks"))
		})
		It("It should generate cluster name based on theme", func() {

			Expect(_theme.ClusterName).ToNot(BeEmpty())
		})

	})
})
