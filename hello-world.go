package main

import "fmt"
import "errors"



func main(){

//fmt.Println("hello world")

if i,j,k,l,err:=f(5,2); err==nil{
fmt.Println(i)
fmt.Println(j)
fmt.Println(k)
fmt.Println(l)
//fmt.Println(err)
}else{

fmt.Println("sorry function failed")
fmt.Println(err)
}
}

func f(x,y int)(int,int,int,float64,error){
if y==0{
return x*y,x+y,x-y,0,errors.New("can't divide by zero")
}

return x*y,x+y,x-y,float64(x)/float64(y),nil
}

