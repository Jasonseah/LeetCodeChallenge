package challenge_1476

type SubrectangleQueries struct {
	array2d [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	class := SubrectangleQueries{
		array2d: rectangle,
	}

	return class
}

func (sq *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	//var MostLeft = sq.array2d[row1][col1]
	//var MostRight = sq.array2d[row2][col2]

	//[{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}]
	for parentKey, parentValue := range sq.array2d {
		// key 0
		// val {1, 2, 3, 4}
		var updateValue = false
		if parentKey >= row1 && parentKey <= row2 {
			updateValue = true
		}

		if updateValue {
			for childKey := range parentValue {
				// key 0
				// val 1

				updateValue = false
				if childKey >= col1 && childKey <= col2 {
					updateValue = true
				}

				if updateValue {
					sq.array2d[parentKey][childKey] = newValue
				}

			}
		}
	}
}

func (sq *SubrectangleQueries) GetValue(row int, col int) int {
	return sq.array2d[row][col]
}

func (sq *SubrectangleQueries) GetArray() [][]int {
	return sq.array2d
}
