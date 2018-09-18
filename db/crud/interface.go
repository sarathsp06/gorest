package crud

// Interface Interface defines all the essential methods require for a CRUD compatible issue
type Interface interface {
	Insert(name string, values ...interface{}) error
	Update(name string, ID string, update interface{}, result interface{}) error
	Delete(name string, ID string, result interface{}) error
	Get(name string, filter interface{}, result interface{}, pageNum, pageSize int) error
}

//default singltone implementation of crud apis
var crud Interface

//SetDefault sets default crud implementation
func SetDefault(crudImpl Interface) {
	crud = crudImpl
}

func Insert(name string, values ...interface{}) error {
	return crud.Insert(name, values...)
}
func Update(name string, ID string, update interface{}, result interface{}) error {
	return crud.Update(name, ID, update, result)
}
func Delete(name string, ID string, result interface{}) error {
	return crud.Delete(name, ID, result)
}
func Get(name string, filter interface{}, result interface{}, pageNum, pageSize int) error {
	return crud.Get(name, filter, result, pageNum, pageSize)
}
