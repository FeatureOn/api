package mocker

import (
	"github.com/FeatureOn/api/server/domain"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// GenerateMockFeature generates a Feature with totally random data
func GenerateMockFeature() domain.Feature {
	return domain.Feature{
		Name:         generateString(50),
		Key:          generateString(20),
		Description:  generateString(100),
		DefaultState: generateBool(),
		Active:       generateBool(),
	}
}

// GenerateMockEnvironment generates an Environment with totally random data
func GenerateMockEnvironment() domain.Environment {
	return domain.Environment{
		ID:   generateUUID(),
		Name: generateString(50),
	}
}

// GenerateMockProduct generates a Product with totally random data
func GenerateMockProduct(haveEnvironment, haveFeature bool) domain.Product {
	product := domain.Product{
		ID:   generateUUID(),
		Name: generateString(50),
	}
	if haveFeature {
		product.Features = make([]domain.Feature, generateCount())
		for i := 0; i < len(product.Features); i++ {
			product.Features[i] = GenerateMockFeature()
		}
	}
	if haveEnvironment {
		product.Environments = make([]domain.Environment, generateCount())
		for i := 0; i < len(product.Environments); i++ {
			product.Environments[i] = GenerateMockEnvironment()
		}
	}
	return product
}

// GenerateMockProductSlice generates a slice of Products with totally random data
func GenerateMockProductSlice() []domain.Product {
	products := make([]domain.Product, generateCount())
	for i := 0; i < len(products); i++ {
		products[i] = GenerateMockProduct(false, false)
	}
	return products
}

func generateUUID() string {
	return uuid.New().String()
}

func generateString(length int) string {
	seed := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	chars := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		chars[i] = seed[rand.Intn(52)]
	}
	return string(chars)
}

func generateBool() bool {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(1)
	if rnd == 0 {
		return false
	}
	return true
}

func generateCount() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(4) + 1

}
