package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	//"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile("test.db", o)
	if err != nil {
		println(err)
	}

	batch := new(leveldb.Batch)
	batch.Put([]byte("1130-2014-01-06"), []byte("10"))
	batch.Put([]byte("1130-2014-01-06"), []byte("another value"))
	batch.Put([]byte("bar-3"), []byte("another value2"))
	batch.Delete([]byte("baz"))
	err = db.Write(batch, nil)

	iter := db.NewIterator(util.BytesPrefix([]byte("1130")), nil)
	for iter.Next() {
		// Use key/value.
		println(string(iter.Value()))

	}
	iter.Release()
	err = iter.Error()

}
