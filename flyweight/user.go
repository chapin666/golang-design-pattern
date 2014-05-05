package flyweight

type User struct {
	name string
}

func (this *User) Name() string {
	return this.name
}
func (this *User) SetName(name string) {
	this.name = name
}
