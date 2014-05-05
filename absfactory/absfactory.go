package absfactory

import (
	"fmt"
)

// User
type User struct {
	id   int
	name string
}

func (this *User) Id() int {
	return this.id
}
func (this *User) SetId(id int) {
	this.id = id
}

func (this *User) Name() string {
	return this.name
}
func (this *User) SetName(name string) {
	this.name = name
}

//Department
type Department struct {
	name string
}

func (this *Department) Name() string {
	return this.name
}
func (this *Department) SetName(name string) {
	this.name = name
}

type IDepartDao interface {
	update(department Department)
	delete(department Department)
}

type DepartMongoDB struct{}

func (this *DepartMongoDB) update(department Department) {
	sql := `db.depart.update({$set : {"name": "%s"}})`
	sql = fmt.Sprintf(sql, department.Name())
	fmt.Println(sql)
}
func (this *DepartMongoDB) delete(department Department) {
	sql := "db.depart.remove({name:`%s`})"
	sql = fmt.Sprintf(sql, department.Name())
	fmt.Println(sql)
}

type DepartMysql struct{}

func (this *DepartMysql) update(department Department) {
	sql := "update depart set name = '%s' where id = 1"
	sql = fmt.Sprintf(sql, department.Name())
	fmt.Println(sql)
}
func (this *DepartMysql) delete(department Department) {
	sql := "delete from depart where name = '%s'"
	sql = fmt.Sprintf(sql, department.Name())
	fmt.Println(sql)
}

type IUserDao interface {
	insert(user User)
	query()
}

//mongodb implements
type MongoDB struct{}

func (this *MongoDB) insert(user User) {
	sql := `db.user.insert({id:%d, name:'%s'})`
	sql = fmt.Sprintf(sql, user.Id(), user.Name())
	fmt.Println(sql)
}
func (this *MongoDB) query() {
	sql := "db.user.find()"
	fmt.Println(sql)
}

//mysql implements
type Mysql struct{}

func (this *Mysql) insert(user User) {
	sql := "insert into user(id, name) values(%d, '%s')"
	sql = fmt.Sprintf(sql, user.Id(), user.Name())
	fmt.Println(sql)
}
func (this *Mysql) query() {
	sql := "select * from user"
	fmt.Println(sql)
}

/*
type IFactory interface {
	CreateUserDao() IUserDao
	CreateDeparDao() IDepartDao
}

// UserFactory
// MongoDB Factory
type MongoDBFactory struct{}

func (this *MongoDBFactory) CreateUserDao() IUserDao {
	return new(MongoDB)
}
func (this *MongoDBFactory) CreateDeparDao() IDepartDao {
	return new(DepartMongoDB)
}

//Mysql Factory
type MysqlFactory struct{}

func (this *MysqlFactory) CreateUserDao() IUserDao {
	return new(Mysql)
}
func (this *MysqlFactory) CreateDeparDao() IDepartDao {
	return new(DepartMongoDB)
}
*/

type DataAccess struct {
	source string
}

func (this *DataAccess) CreateUser() (dao IUserDao) {
	if this.source == "mysql" {
		dao = new(Mysql)
	} else {
		dao = new(MongoDB)
	}

	return dao
}
func (this *DataAccess) CreateDepart() (dao IDepartDao) {
	if this.source == "mysql" {
		dao = new(DepartMysql)
	} else {
		dao = new(DepartMongoDB)
	}

	return dao
}
