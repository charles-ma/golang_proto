package book_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ginkgo_tests/book"
)

var _ = Describe("Book", func() {
	var foxInSocks, lesMis *book.Book

	BeforeEach(func() {
		foxInSocks = &book.Book{
			Title:  "Fox in Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}

		lesMis = &book.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
	})

	Describe("Categorizing books", func() {
		Context("with fewer than 300 pages", func() {
			It("should be short stories", func() {
				By("this is a description")

				Expect(foxInSocks.Category()).To(Equal(book.CatetoryShortStory))
			})
		})

		Context("with more than 300 pages", func() {
			It("should be noval", func() {
				Expect(lesMis.Category()).To(Equal(book.CategoryNoval))
			})
		})
	})
})
