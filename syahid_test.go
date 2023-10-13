package backendgis

import (
	"fmt"
	"testing"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "geojson")
	datalokasi := GetAllBangunanLineString(mconn, "bandaaceh")
	fmt.Println(datalokasi)
}