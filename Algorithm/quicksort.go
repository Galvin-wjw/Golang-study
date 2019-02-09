package main

import "fmt"

// int Partition(int arr[],int low,int high){
//    int temp =arr[low];
//    while(low<high){    //low和high从待排序列两端向中间移动
//        while(low<high&&arr[high]>=temp) --high;
//        arr[low] = arr[high];    //将 比temp小的移到低端
//        while(low<high&&arr[low]<=temp) --low;
//        arr[high] = arr[low];    //将比temp大的移到高端
//    }
//    arr[low] = temp;
//    return low;
//}

//void QSort(int arr[],int s,int t){
//    int pirvotloc;
//    if(s<t){
//        pirvotloc = Partition(arr,s,t); //进行划分并返回基数位置
//        QSort(arr,s,pirvotloc-1);       //对基数前的子序列进行递归排序
//        QSort(arr,pirvotlov+1,t);      //对基数后的子序列进行递归排序
//    }
//}

func Quick3Sort(a []int,left int, right int)  {

	if left >= right {
		return
	}

	BaseIndex := left

	for i := left + 1; i <= right ; i++ {

		if a[left] >= a[i]{

			//分割位定位++
			BaseIndex ++;
			fmt.Println(i,BaseIndex)
			a[i],a[BaseIndex] = a[BaseIndex],a[i]
			fmt.Println("for cycle:",a)

		}

	}

	//起始位和分割位
	a[left], a[BaseIndex] = a[BaseIndex],a[left]

	Quick3Sort(a,left,BaseIndex - 1)
	Quick3Sort(a,BaseIndex + 1,right)

}

func main() {
	//var arr = []int{6,2,7,3,8,9}
	var arr = []int{6,2,3,1,4,5}
	Quick3Sort(arr,0,5)
	fmt.Println(arr)
}
