package main

import "log"

func main() {

	var num int
	num = 100

	log.Println("FINDING THE NO. BY SIMPLE IF-ELSE")
	
	if num == 100 {
		log.Println("The number is 100 ")
	}else {
		log.Println("the number is not 100")
	}
   //  Doing same thing with Switch statement 
   
   log.Println("FINDING THE NO. BY SWITCH STATEMENT")
  
   switch num {
   case 100 :
	log.Println("the number is 100")
   default:
	log.Println("the number is not 100")
   }

}