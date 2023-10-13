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

func GetAllGeoData(MongoConnect *mongo.Database, colname string) []GeoJson {
	data := atdb.GetAllDoc[[]GeoJson](MongoConnect, colname)
	return data
}