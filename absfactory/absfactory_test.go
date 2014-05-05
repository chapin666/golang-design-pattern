package absfactory

import (
	"testing"
)

/*
func Test_MongoDB(t *testing.T) {
	user := User{}
	user.SetId(1)
	user.SetName("张三")

	factory := new(MysqlFactory)
	dao := factory.CreateUserDao()

	dao.insert(user)
}

func TestDepart(t *testing.T) {
	department := Department{}
	department.SetName("scitc")

	factory := new(MongoDBFactory)
	depart := factory.CreateDeparDao()
	depart.delete(department)
}
*/
func TestAccess(t *testing.T) {
	access := DataAccess{source: "mysql"}
	dao := access.CreateDepart()

	department := Department{}
	department.SetName("scitc")
	dao.update(department)
}
