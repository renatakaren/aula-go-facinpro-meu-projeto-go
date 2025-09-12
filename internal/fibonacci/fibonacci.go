package fibonacci

func Fibonacci(n int) []int {
	if n < 0 {
		return []int{}
	}

	seq := make([]int, n+1)
	seq[0] = 0
	if n >= 1 {
		seq[1] = 1
		for i := 2; i <= n; i++ {
			seq[i] = seq[i-1] + seq[i-2]
		}
	}
	return seq
}
