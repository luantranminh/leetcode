package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i > j {
				continue
			}

			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {

		for k, l := 0, len(matrix[i])-1; k < len(matrix[i])/2; k, l = k+1, l-1 {
			matrix[i][k], matrix[i][l] = matrix[i][l], matrix[i][k]
		}
	}

}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
