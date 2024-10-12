package generator

type IdGenerator func() int

func IncGeneratorID() func() int {
	count := 1

	return func() int {
		count++
		return count
	}
}

func DecGeneratorID() func() int {
	count := -1

	return func() int {
		count--
		return count
	}
}
