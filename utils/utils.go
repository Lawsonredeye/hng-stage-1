package utils

import "math"

func IsPrime(num int64) bool {
	var (
		i int64
		flag int = 0
	)

	if num < 0 {
		return false
	}

	if num == 1 || num == 2 {
		return true
	}

	for i = 2; i <= num/2; i++ {
		if num % i == 0 {
			flag = 1
			break
		}
	}
	return flag == 0
}

func IsPerfect(num int64) bool {
	if num < 1 {
		return false
	}

	var (
		sum int64 = 1
		i int64 
	)
	for i = 2; i*i <= num; i++ {
		if num % i == 0 {
			if i*i != num {
				sum += i + num/i
			} else {
				sum += i
			}
		}
	}

	return (sum == num && sum != 1)
}

func DigitSum(num int64) (int64,  int){
	var (
		sum int64
		i int64
		count int
	)

	num = int64(math.Abs(float64(num)))

	for i = num; num > 0; i++{
		c := num % 10
		sum += c
		num /= 10
		count += 1
	}
	return sum, count
}

func Properties(num int64) []string {
	var sum int64
	numCopy := num
	_, count:= DigitSum(num)

	for i := num; num > 0; i++ {
		c := num % 10
		sum += int64(math.Pow(float64(c),float64(count)))
		num /= 10
	}
	
	properties := make([]string, 0)
	println(numCopy)
	if numCopy == sum {
		properties = append(properties, "armstrong")
	}

	if numCopy % 2 == 1 {
		properties = append(properties, "odd")
	}else {
		properties = append(properties, "even")
	}
	return properties
}