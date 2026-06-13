package pkg

type Floor struct {
	number int
}

func NewFloor(num int) *Floor {
	return &Floor{
		number: num,
	}
}
