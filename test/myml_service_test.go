package test

import (
	"github.com/mercadolibre/myml/src/api/services/myml"
	"testing"
)

func BenchmarkGetMyMLFromAPI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myml.GetMyMLFromAPI(1234567)
	}
}
