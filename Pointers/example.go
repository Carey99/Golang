//write a function that uses pointers to swap values

func swap(x, y *int) {
	temp := *x
	*x = y
	*y = temp
}