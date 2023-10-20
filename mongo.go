package backendgis

import (
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func GetConnectionMongo(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func GetAllData(MongoConnect *mongo.Database, colname string) []GeoJson {
	data := atdb.GetAllDoc[[]GeoJson](MongoConnect, colname)
	return data
}

func InsertDataLonlat(MongoConn *mongo.Database, colname string, coordinate []float64, name, volume, tipe string) (InsertedID interface{}) {
	req := new(LonLatProperties)
	req.Type = tipe
	req.Coordinates = coordinate
	req.Name = name
	req.Volume = volume

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
}