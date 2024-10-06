package entities

type Lane struct {
	ID         int
	isOccupied bool
}

// NewLane creates a new Lane with the specified ID.
func NewLane(id int) *Lane {
	return &Lane{
		ID: id,
	}
}

// GetID returns the ID of the lane.
func (l *Lane) GetID() int {
	return l.ID
}

// IsOccupied checks if the lane is currently occupied.
func (l *Lane) IsOccupied() bool {
	return l.isOccupied
}

// SetOccupied sets the occupancy status of the lane.
func (l *Lane) SetOccupied(occupied bool) {
	l.isOccupied = occupied
}

// Optional: If you plan to use pins in the future, you can uncomment and improve this section.
// func (l *Lane) GetPins() []int {
//     return l.pins
// }
//
// func (l *Lane) SetPins(pins []int) {
//     l.pins = pins
// }
