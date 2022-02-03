package lib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"web-dev-with-golang-book-by-shiju/chapter-10/bdd-testing/lib"
)

var _ = Describe("Users", func() {
	BeforeEach(func() {

	})

	Describe("Get Users", func() {
		Context("Get all Users", func() {
			It("should get list of Users", func() {

			})
		})
	})

	Describe("Post a new User", func() {
		Context("Provide valid User data", func() {
			It("should create a new User and get HTTP status: 201", func() {})
		})
		Context("Provide User data that contains duplicate email id", func() {
			It("should get HTTP status: 400", func() {})
		})
	})
})
