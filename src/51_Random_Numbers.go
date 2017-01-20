package main

import (
	"time"
	"fmt"
	"math/rand"
)

// 特别注意：rand是非线程安全
func main() {

	// Intn(n) 随机生成[0, 100)之间的整数
	fmt.Println(rand.Intn(100), ",", rand.Intn(100))

	// 随机生成[0.0, 1.0)之间的64位小数
	fmt.Println(rand.Float64())

	// 使用time.Now().UnixNano()生成的随机种子很难复现
	s1 := rand.NewSource(time.Now().UnixNano())		// 生成随机种子
	r1 := rand.New(s1)		//根据随机种子，生成随机生成器

	fmt.Println(r1.Intn(100), ",", r1.Intn(100))

	s2 := rand.NewSource(42)	// 根据42生成随机种子
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100), ",", r2.Intn(100))

	s3 := rand.NewSource(42)	// 和22行使用了相同的随机因子42
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100), ",", r3.Intn(100))	// 和24行产生的随机数相同

}
