package entities

type Arena struct {
	Lanes []Lane
}

func NewArena() *Arena {
	return &Arena{
		Lanes: make([]Lane, 0),
	}
}

func (a *Arena) AddLane() {
	l := NewLane(len(a.Lanes))
	a.Lanes = append(a.Lanes, *l)
}

func (a *Arena) GetLanes() []Lane {
	return a.Lanes
}

func (a *Arena) GetAvailableLanes() *Lane {
	for _, lane := range a.Lanes {
		if !lane.isOccupied {
			return &lane
		}
	}
	return nil
}

func (a *Arena) RemoveLane(laneNumber int) {
	a.Lanes = append(a.Lanes[:laneNumber], a.Lanes[laneNumber+1:]...)
}
