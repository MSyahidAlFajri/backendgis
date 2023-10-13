package backendgis

import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("string", "geojson", "bandaaceh")

	fmt.Printf("%+v", data)
}