package syahid

import (
	"encoding/json"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datalokasi := GetAllBangunanLineString(mconn, collectionname)
	jsondatalokasi, _ := json.Marshal(datalokasi)
	return string(jsondatalokasi)
}