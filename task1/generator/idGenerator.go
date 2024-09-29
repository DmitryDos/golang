package generator

type IdGenerator func() int

func GeneratorA() func() int {
	count := 1

	return func() int {
		count++
		return count
	}
}

func GeneratorB() func() int {
	count := -1

	return func() int {
		count--
		return count
	}
}
