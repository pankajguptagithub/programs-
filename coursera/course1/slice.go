package main
import ( 
	      "sort"
	      "fmt"
		  "strconv"		  
)
func main(){
	sli := make([]int,3)
	brk:
	for{
		var n string
		fmt.Println("Enter any integer ")
		fmt.Scan(&n)		
		if (n == "X"){
			fmt.Println("Program ends")
			break brk			
		}else{
			n1,_ := strconv.Atoi(n)
			sli = append(sli,n1)
			sort.Ints(sli)
			fmt.Println(sli)
		}	
	}
}