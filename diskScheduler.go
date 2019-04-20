/*“I Alexander Meade al565538 affirm that this program is entirely my own work and
that I have neither developed my code together with any another person, nor copied any code from any
other person, nor permitted my code to be copied or otherwise used by any other person, nor have I copied,
modified, or otherwise used programs created by others. I acknowledge that any violation of the above terms
will be treated as academic dishonesty.”
*/

//scheduling algs FCFS SJFP RR

//go imports
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


var (
	lowerCYL int       //valid lower num 
	upperCYL       int       //valid upper num must be > than lower
	use          string    //which alg to use 
	cylreq      	[]int       //cylinder request 
	initCYL        int  //inital position must be 0<init<upper
)

//main fucntion where readin and stuff occurs
func main() {

	fileRead := os.Args[1]
	file, _ := os.Open(fileRead) //open file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) //scan in file as words
	var currCyl int 
	for scanner.Scan() {
		checkStr := scanner.Text()
		//all info save use will need to have text converted to string
		if checkStr == "use" {
			scanner.Scan()
			use = scanner.Text()
			fmt.Printf("Seek algorithm: %s\n",use)
		}else if checkStr == "lowerCYL" {
			scanner.Scan()
			lowerCYL, _ = strconv.Atoi(scanner.Text())
			fmt.Printf("\t Lower cylinder:\t %d\n",lowerCYL)

		} else if checkStr == "upperCYL" {
			scanner.Scan()
			upperCYL, _ = strconv.Atoi(scanner.Text())
			if upperCYL < lowerCYL{
				fmt.Printf("Error upper bound not greater than lower bound")
			}else{
			fmt.Printf("\t Upper cylinder:\t %d\n",upperCYL)
				}

		} else if checkStr == "initCYL" {
			scanner.Scan()
			initCYL, _ = strconv.Atoi(scanner.Text())
			if initCYL < 0 || initCYL > upperCYL{
				fmt.Printf("Error inital cycle is outside of specifed range")
			}else{
			fmt.Printf("\t Init cylinder:\t %d\n",initCYL)
			}
		}else if checkStr == "cylreq"{
			scanner.Scan()
			currCyl, _ = strconv.Atoi(scanner.Text())
			cylreq[currCyl] = currCyl
			if currCyl < lowerCYL || currCyl > upperCYL{
				fmt.Printf("Error reqested cycle is outside of specifed range")
			} else{
				fmt.Printf("Cylinder\t%d",currCyl) //print the request
			}

		}else if checkStr == "end"{
			break //once end done scanning 
		}
	}

	switch use {
	case "fcfs":{
			fcfs(initCYL,lowerCYL,upperCYL,cylreq)
			break
		}
	case "sstf":{
			
			break
		}
	case "SCAN":{
			break
		}
	case "CSCAN":{
			break
		}
	case "LOOK":{
			break
		}
	case "CLOOK":{
			break
		}

	}
}



func fcfs(inital int ,lower int , upper int, requests []int ){
	//for time < len(requests) {  //number of requests
	traversal := 0 //set inital traversal to 0 
	requestQ := make([]int, len(requests)) //make a queue the size of requests
	requestQ[0] = inital //start at init value 
	curr := inital //current location vs next 
		//populate the request queue
		for i := 1; i < len(requests); i++ {
				requestQ[i] = requests[i]  //add procc to queue
				//amount++               // added to queue knows it has things to run
		}
	//work the requests
	for i:= 0; i < len(requestQ); i++{
		curr = requestQ[i]
		fmt.Printf("Servicing\t %d", requestQ[i])//servicing current one
		for j:=i+1; j < len(requestQ); j++{ //need to keep vals positive 
			if curr > requestQ[j]{
				traversal += curr - requestQ[j]
			} else if curr < requestQ[j]{
					traversal += requestQ[j] - curr 
			}
		}
		fmt.Printf("FCFS traversal count = %d", traversal)
	}
}


func sstf(inital int ,lower int , upper int, requests []int){
	//for time < len(requests) {  //number of requests
	traversal := 0 //set inital traversal to 0 
	requestQ := make([]int, len(requests)) //make a queue the size of requests
	requestQ[0] = inital //start at init value 
	curr := inital //current location vs next 
		//populate the request queue
		for i := 1; i < len(requests); i++ {
				requestQ[i] = requests[i]  //add procc to queue
				//amount++               // added to queue knows it has things to run
		}
	sortedQ := selSort(requestQ)	
	//work the requests
	for i:= 0; i < len(requestQ); i++{
		curr = requestQ[i]
		fmt.Printf("Servicing\t %d", requestQ[i])//servicing current one
		for j:=i+1; j < len(requestQ); j++{ //need to keep vals positive 
			if curr > requestQ[j] || sortedQ[j] < sortedQ[curr] { //if the next seek time is shorter 
				traversal += curr - requestQ[j]
			} else if curr < requestQ[j] || sortedQ[j] < sortedQ[curr]{
					traversal += requestQ[j] - curr 
			}
		}
		fmt.Printf("SSTF traversal count = %d", traversal)
	}
}

func SCAN(inital int ,lower int , upper int, requests []int){
	//for time < len(requests) {  //number of requests
	traversal := 0 //set inital traversal to 0 
	requestQ := make([]int, len(requests)) //make a queue the size of requests
	requestQ[0] = inital //start at init value 
	curr := inital //current location vs next
	temp := inital + 1 //next value after curr  
		//populate the request queue
		for i := 1; i < len(requests); i++ {
				requestQ[i] = requests[i]  //add procc to queue
				//amount++               // added to queue knows it has things to run
		}
	sortedQ := selSort(requestQ) //at this point sorted from least to greatest 	
	amount := len(requestQ)
	for i:= curr; i >= 0 ; i--{
		if i == 0 && amount != 0{
			break
		}
		curr = requestQ[i]
		amount--
		fmt.Printf("Servicing\t %d", requestQ[i])//servicing current one
		traversal += curr - requestQ[i-1]
		//above handles from curr to 0 below should do curr+1 to end 
		if i == 0{
		for temp; temp < len(requestQ); temp++{ //need to keep vals positive 
			fmt.Printf("Servicing\t %d", requestQ[temp])
			travesal += requestQ[temp+1] - requestQ[temp] // where we go - where we are rn    
		}
	}
		fmt.Printf("SCAN traversal count = %d", traversal)
	}
}

func CSCAN(inital int ,lower int , upper int, requests []int){

}

func LOOK(inital int ,lower int , upper int, requests []int){

}

func CLOOK(inital int ,lower int , upper int, requests []int){

}

//selection sort 
func selSort(reqs []int) (sortedReqs []int) {
	sortedReqs = make([]int, len(reqs))
	sortedReqs = reqs
		for i := 0; i < (len(sortedReqs) - 1); i++ {
			selNum := i
			for j := i + 1; j < len(sortedReqs); j++ {
				if sortedReqs[j] < sortedReqs[selNum] {
					selNum = j
				}
			}
			temp := sortedReqs[selNum]
			sortedReqs[selNum] = sortedReqs[i]
			sortedReqs[i] = temp
			}
			return sortedReqs
}

//find the min of slice 
func Min(array []int) (int) {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min, max
}	
//find the max of slice 
func Max(array []int) (int) {
    var max int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
    }
    return min, max
}