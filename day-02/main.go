// In this program we will use  pointer to change the value of a variable 
//(It is cool because that variable is created in main function and it cannot be used or changed in a different function but by  using its location we can change the data present there).
package main

import "log"

func main() {

	var mountains string
	mountains = "HIMALAYAS"

    log.Println("printing value of var mountains before calling the Use_pointer function =",mountains)
	log.Println("The location of mountain is: ",&mountains)// This prints the location of mountains variable and you will see that this output will be equal to the value of s 
	Use_Pointer(&mountains) // This &variable is used to point to the location where the variable's content exist
	log.Println("printing value of var mountains after  calling the Use_pointer function =",mountains)
	
}

func Use_Pointer(s *string) { // Here s is a pointer to a variable of type string

	log.Println("Value of s is: ",s)
	var desert string
	desert = "THAR"
	

	*s = desert // This stores the value of desert in the location of variable s

	
}
