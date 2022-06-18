package internal

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

//https://dev.to/go/using-boltdb-as-internal-database-39bd refer to this and figure out how to map SeadexEntry
// look into this -> https://github.com/asdine/storm
func dbInit() {
	db, err := bolt.Open("seadex.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		log.Fatalln("Error initializing db", err)
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Todo"))
		if err != nil {
			log.Fatalln("Error in creating db", err)
		}
		return bucket.Put([]byte("1"), []byte("Test"))
	})
}
