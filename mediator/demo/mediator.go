package demo

// Mediator
// 中介者interface
type UnitedNations interface {
	Declare(msg string, country Country)
}

// Mediator Concrete
// implements
type UnitedNationsA struct {
	usa  USA
	iraq Iraq
}

func (this *UnitedNationsA) setUSA(usa USA) {
	this.usa = usa
}
func (this *UnitedNationsA) setIraq(iraq Iraq) {
	this.iraq = iraq
}

func (this *UnitedNationsA) Declare(msg string, country *interface{}) {

	if _, ok := country.(USA); ok {
		this.usa.getMessage(msg)
	} else {
		this.iraq.getMessage(msg)
	}

	/*switch value := country.(type) {
	case this.iraq:
		this.usa.getMessage(msg)
	case this.usa:
		this.iraq.getMessage(msg)
	default:
		fmt.Println(country)
	}*/
}
