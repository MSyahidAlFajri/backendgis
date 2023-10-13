package backendgis

import "encoding/json"

func GCHandlerFunc(Mongostring, dbname, colname string) []byte {
	koneksyen := GetConnectionMongo(Mongostring, dbname)
	datageo := GetAllGeoData(koneksyen, colname)

	jsonkuy, _ := json.Marshal(datageo)

	return jsonkuy
}