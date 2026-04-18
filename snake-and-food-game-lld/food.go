package main

type FoodItem interface {
	getRow() int
	getColumn() int
	getPoints() int
}

type NormalFood struct {
	row    int
	column int
	points int
}

func NewNormalFood(row int, column int) *NormalFood {
	return &NormalFood{
		row:    row,
		column: column,
		points: 1,
	}
}

func (nf *NormalFood) getRow() int {
	return nf.row
}

func (nf *NormalFood) getColumn() int {
	return nf.column
}

func (nf *NormalFood) getPoints() int {
	return nf.points
}

type SpecialFood struct {
	row    int
	column int
	points int
}

func NewSpecialFood(row int, column int) *SpecialFood {
	return &SpecialFood{
		row:    row,
		column: column,
		points: 3,
	}
}

func (sf *SpecialFood) getPoints() int {
	return sf.points
}

func (sf *SpecialFood) getRow() int {
	return sf.row
}

func (sf *SpecialFood) getColumn() int {
	return sf.column
}

func FoodFactory(foodType string, pair Pair) FoodItem {
	switch foodType {
	case "normal":
		return NewNormalFood(pair.row, pair.col)
	case "bonus":
		return NewSpecialFood(pair.row, pair.col)
	default:
		return NewNormalFood(pair.row, pair.col)
	}
}
