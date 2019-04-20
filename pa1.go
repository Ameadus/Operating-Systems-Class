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
    "sort"

)

type Process struct {
    name string 
    arrival int 
    burst int 
}


var (
    processcount int 
    runfor int 
    use string 
    quantum int //for RR 
)


 //main fucntion where readin and stuff occurs
func main(){ 
    
    fmt.Scan(&processcount)
    reader := bufio.NewReader(os.Stdin)
    processcount, _ := reader.ReadString('\n')
    fmt.Scan(&runfor)
    runfor, _ := reader.ReadString('\n')
    fmt.Scan(&use)
    use, _ := reader.ReadString('\n')

    fmt.Println(processcount)
    fmt.Println(runfor)
    fmt.Println(use)

   // if use == FCFS
     //FCFS()


}


func FCFS(){
    var( 
    i int 
    name []string 
    arrival []int 
    burst []int 

    )
    //var arrProcess [] Process //create array of Processes 
for i = 0; i < processcount; i++ {  
    fmt.Scan(&name)
    fmt.Scan(&arrival)
    fmt.Scan(&burst)
}
    //Process.name, _ := reader.ReadString('\n')
    //Process.arrival, _ := reader.ReadString('\n')
    //Process.burst, _ := reader.ReadString('\n')
    sort.Sort(sort.Reverse(sort.Ints(arrival))) //sorts the arrivals from greatest to least 
    fmt.Println(arrival)
}
    /*for i = 0 , count < processcount ; count++ (
        Process[i]
    )
}

func SJFP(){

}

func RR(){


}
*/
//func readFile(file, err)