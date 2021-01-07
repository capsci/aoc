package helper

func Give3DArray(x, y, z int, init string) (multiDArr [][][]string) {
	multiDArr = make([][][]string, z)
	for i := 0; i < z; i++ {
		multiDArr[i] = make([][]string, x)
		for j := 0; j < x; j++ {
			multiDArr[i][j] = make([]string, y)
			for k := 0; k < y; k++ {
				multiDArr[i][j][k] = init
			}
		}
	}
	return
}

func GiveEmpty3DArray(x, y, z int) (multiDArr [][][]string) {
	multiDArr = make([][][]string, z)
	for i := 0; i < z; i++ {
		multiDArr[i] = make([][]string, x)
		for j := 0; j < x; j++ {
			multiDArr[i][j] = make([]string, y)
		}
	}
	return
}
