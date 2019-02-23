package main



//r表示数组最后一个下标，p表示第一个
func mergesort(a []int,p int ,r int)  {
	//lenth := len(a)
	if p >= r{
		return
	}

	m := (p+r)/2
	mergesort(a,r,m)
	mergesort(a,m+1,p)



	//s1 := make([]int,lenth/2)


}