package main

import (
	"fmt"
	"time"
)

func main(){
	a:=make([]int64,5)
	var b string
	for i:=0;i<=cap(a)-1;i++{
		_, _ = fmt.Scanf("%d\n", &a[i])
		fmt.Println("improt ok!")
	}
	_, _ = fmt.Scanf("%s", &b)
	if(b=="output"){
		fmt.Println("the results are:")
		for j:=0;j<=cap(a)-1;j++ {
			time1 :=time.Unix(a[j],1).Format("2006-01-02 15:04:05")
			fmt.Println(time1)
		}
	}
}


