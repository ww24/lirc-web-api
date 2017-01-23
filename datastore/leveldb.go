package datastore

import (
	"bytes"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/ugorji/go/codec"
)

const datastorePath = "datastore/leveldb"

// LevelDBStore structure
type LevelDBStore struct {
	msgh *codec.MsgpackHandle
	db   *leveldb.DB
}

// NewLevelDBStore constructor
func NewLevelDBStore() (ld *LevelDBStore) {
	ld = &LevelDBStore{
		msgh: &codec.MsgpackHandle{RawToString: true},
	}

	// open LevelDB
	db, err := leveldb.OpenFile(datastorePath, nil)
	if err != nil {
		panic(err)
	}
	ld.db = db

	return
}

// Save method
func (ld *LevelDBStore) Save(id string, item interface{}) (err error) {
	// encode to MessagePack and save into LevelDB
	buff := &bytes.Buffer{}
	err = codec.NewEncoder(buff, ld.msgh).Encode(item)
	if err != nil {
		return
	}
	err = ld.db.Put([]byte(id), buff.Bytes(), nil)
	return
}

// Keys method
func (ld *LevelDBStore) Keys(prefixes ...string) (list []string, err error) {
	prefix := ""
	if len(prefixes) > 0 {
		prefix = strings.Join(prefixes, ":")
	}

	list = make([]string, 0, 100)

	iter := ld.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	for iter.Next() {
		key := iter.Key()
		list = append(list, string(key))
	}
	err = iter.Error()
	if err != nil {
		return
	}
	return
}

// Fetch method
func (ld *LevelDBStore) Fetch(id string, item interface{}) (err error) {
	var data []byte
	data, err = ld.db.Get([]byte(id), nil)
	if err == leveldb.ErrNotFound {
		err = nil
		return
	}
	if err != nil {
		return
	}

	// read data from LevelDB and decode MessagePack
	err = codec.NewDecoderBytes(data, ld.msgh).Decode(item)

	return
}

// Exists method
func (ld *LevelDBStore) Exists(id string) (exists bool, err error) {
	exists, err = ld.db.Has([]byte(id), nil)
	return
}

// Remove method
func (ld *LevelDBStore) Remove(id string) (err error) {
	err = ld.db.Delete([]byte(id), nil)
	return
}
