package root

var ROOT_COUNT int64 = 0

func InsertAndCount(increment bool) int64 {
	if increment {
		ROOT_COUNT++
	}
	return ROOT_COUNT
}
