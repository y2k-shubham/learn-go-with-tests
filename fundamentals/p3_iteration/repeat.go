package p3_iteration

const maxTimes = 5

func Repeat(s string) (repeated string) {
	for i := 1; i <= maxTimes; i++ {
		repeated += s
	}
	return
}
