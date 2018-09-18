package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/sarathsp06/gorest/db/crud"
)

// FieldKey is the struct tag for field names
var FieldKey = "bson"

// PageSize if the default page size applied if none supplied for search
var PageSize = 100000

//Insert to insert a doecument ot the database
func (db *DB) Insert(name string, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}
	err := db.C(name).Insert(values...)
	if err != nil {
		return err
	}
	return nil
}

// Update updates rows given filter
func (db *DB) Update(name string, ID string, update interface{}, result interface{}) error {
	updateMap, err := StructToMap(update, FieldKey, true)
	if err != nil {
		return err
	}
	collection := db.C(name)
	if err = collection.UpdateId(ID, bson.M{"$set": bson.M(updateMap)}); err != nil {
		if isNotFound(err) {
			return crud.ErrNotFound
		}
		return err
	}
	return collection.FindId(ID).One(result)
}

// Delete delets a row given id and table name
func (db *DB) Delete(name string, ID string, result interface{}) error {
	collection := db.C(name)
	if err := collection.FindId(ID).One(result); err != nil {
		if isNotFound(err) {
			return crud.ErrNotFound
		}
		return err
	}
	if err := collection.RemoveId(ID); err != nil {
		if isNotFound(err) {
			return crud.ErrNotFound
		}
		return err
	}
	return nil
}

// Get fiters the table given filter and pagination logic
// Current implementation does not support search with other types of predicates
// The request is paginated  with pagenum and page size parameters
func (db *DB) Get(name string, filter interface{}, result interface{}, pageNum int, pageSize int) error {
	filterMap, err := StructToMap(filter, FieldKey, true)
	if err != nil {
		return err
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = PageSize
	}

	skip := pageSize * (pageNum - 1)
	return db.C(name).Find(bson.M(filterMap)).Skip(skip).Limit(pageSize).All(result)
}
