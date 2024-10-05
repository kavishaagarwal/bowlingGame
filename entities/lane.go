package entities

type Lane struct {
	ID int
	//pins []int
	isOccupied bool
}

func NewLane(id int) *Lane {
	return &Lane{
		ID: id,
	}
}

func (l *Lane) GetID() int {
	return l.ID
}

//
//func (l *Lane) SetID(id int) {
//	l.ID = id
//}

//func (l *Lane) GetPins() []int {
//	return l.pins
//}
//
//func (l *Lane) SetPins(pins []int) {
//	l.pins = pins
//}
