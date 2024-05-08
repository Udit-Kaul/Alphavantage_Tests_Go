package test_test

import (
	"go-alphavantage-tests/pkg/client"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AlphaVantage API", func() {
	var (
		apiClient *client.AlphavantageClient
		apiKey    string
	)

	BeforeEach(func() {
		// Load environment variables from .env file
		dir, err1 := os.Getwd()
		if err1 != nil {
			log.Fatal(err1)
		}
		log.Println("Current working directory:", dir)
		err := godotenv.Load("../.env")
		Expect(err).NotTo(HaveOccurred())

		// Get the API key from environment variable
		apiKey = os.Getenv("API_KEY")
		Expect(apiKey).NotTo(BeEmpty())

		// Initialize the API client
		apiClient = client.NewClient(apiKey)
	})

	Describe("Global Quote", func() {
		Context("When requesting stock information for a valid symbol", func() {
			It("should return HTTP 200", func() {
				resp, err := apiClient.FetchStock("IBM")
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("When requesting stock information for an invalid symbol", func() {
			It("should not return HTTP 200", func() {
				resp, err := apiClient.FetchStock("WronfSymbol")
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()
				log.Println("status code:", resp.Body)
				Expect(resp.Status).ToNot(Equal(http.StatusOK))
			})
		})
	})
})
