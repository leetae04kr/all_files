package main

import "fmt"

func main() {
	var i1, i2, i3 int = 2, 2, 3
	fmt.Println(i1, "==", i2, " 결과", i1 == i2)
	fmt.Println(i1, "==", i3, " 결과", i1 == i3)

	var f1, f2, f3 float32 = 0.1, 0.1, 0.2
	fmt.Println(f1, "==", f2, " 결과", f1 == f2)
	fmt.Println(f1, "==", f3, " 결과", f1 == f3)

	var s1, s2, s3 string = "abc", "abc", "ABC"
	fmt.Println(s1, "==", s2, " 결과", s1 == s2)
	fmt.Println(s1, "==", s3, " 결과", s1 == s3)

	var hs1, hs2, hs3 string = "강감찬", "강감찬", "홍길동"
	fmt.Println(hs1, "==", hs2, " 결과", hs1 == hs2)
	fmt.Println(hs1, "==", hs3, " 결과", hs1 == hs3)
}
