package database

import (
	"github.com/hashicorp/go-memdb"
)

type IP struct {
	IP      string
	Lip     string
	City    string
	CitySub string
	Lat     float64
	Lon     float64
	Status  int
}

func InitMemDb() *memdb.MemDB {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"ip": &memdb.TableSchema{
				Name: "ip",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "IP"},
					},
					"lip": &memdb.IndexSchema{
						Name:    "lip",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Lip"},
					},
					"city": &memdb.IndexSchema{
						Name:    "city",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "City"},
					},
					"citysub": &memdb.IndexSchema{
						Name:    "citysub",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CitySub"},
					},
					"lat": &memdb.IndexSchema{
						Name:    "lat",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Lat"},
					},
					"lon": &memdb.IndexSchema{
						Name:    "lon",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Lon"},
					},
					"status": &memdb.IndexSchema{
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Status"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return db
}
