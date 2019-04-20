
package main

import (
     "fmt"
     "os"
     "bufio"
     "strconv"
)


var (
  lowerCYL int
  upperCYL int
  initCYL int
  use string
  cylinders []Cylinder
)

type Cylinder struct {
  index int
  id int
  complete bool

}

func main() {

  inputFile := os.Args[1] //input file 
  
  lowerCYL, upperCYL, initCYL, use, cylinders = getInput(inputFile) //read in inputs 

   if(use == "fcfs") {
    fcfs(lowerCYL, upperCYL, initCYL, cylinders)
  } else if (use == "sstf") {
    // sstf(lowerCYL, upperCYL, initCYL, cylinders)
  } else if (use == "scan") {
    SCAN(lowerCYL, upperCYL, initCYL, cylinders)
  } else if (use == "c-scan") {
    CSCAN(lowerCYL, upperCYL, initCYL, cylinders)
  } else if (use == "look") {
    LOOK(lowerCYL, upperCYL, initCYL, cylinders)
  } else if (use == "c-look") {
    CLOOK(lowerCYL, upperCYL, initCYL, cylinders)
  } 

}
   



//read input 
func getInput (filename string) (lower int, upper int, initial int, use string, reqList []Cylinder){

  reqList = make([]Cylinder, 0, 20)
  var tempCyclinder Cylinder

  //open the input file and get a pointer to it
  file, _ := os.Open(filename)

  //make a pointer for the text inside the file
  fileScanner := bufio.NewScanner(file)

  //split the file into words
  fileScanner.Split(bufio.ScanWords)

  //iterate through the words in the file
  for fileScanner.Scan() {

    //get the first word as a string and set it to word variable 
    word := fileScanner.Text()

    //break the loop if "end" is encountered in the input file
    if(word == "end"){
      break
    }

    if(word == "cylreq") {

      //increment pointer
      fileScanner.Scan()

      tempCyclinder.id, _ = strconv.Atoi(fileScanner.Text())
      
      reqList = append(reqList, tempCyclinder)
      

    } else if(word == "use") {
      //increment pointer
      fileScanner.Scan()
      //assign the algorithm name to algorithmType variable
      use = fileScanner.Text()

    } else if(word == "lowerCYL") {

      //increment pointer
          fileScanner.Scan()
          //assign lower cylinder value to lower variable
          lower,_ = strconv.Atoi(fileScanner.Text())

        } else if(word == "upperCYL") {

          //increment pointer
          fileScanner.Scan()
          //assign upper cylinder value to upper variable
          upper,_ = strconv.Atoi(fileScanner.Text())

        } else if(word == "initCYL"){

          //increment pointer
          fileScanner.Scan()
          //assign initial cylinder value to initial variable
          initial,_ = strconv.Atoi(fileScanner.Text())
        } 

  }

  return lower, upper, initial, use, reqList
}

func fcfs(lower int, upper int, initial int, reqList []Cylinder)  {

  fmt.Printf("Seek algorithm: FCFS\n")
  fmt.Printf("\tLower cylinder: %5d\n", lower)
  fmt.Printf("\tUpper cylinder: %5d\n", upper)
  fmt.Printf("\tInit cylinder: %5d\n", initial)
  fmt.Printf("\tCylinder requests:\n")

  //print list of cylinders
  for i:=0; i<len(reqList); i++ {
    fmt.Printf("\t\tCylinder %5d\n", reqList[i].id)
  }

  totalRequests  := len(reqList)
  traversal := 0
  previousR := initial

  for i := 0; i < totalRequests; i++ {

    //get the current requested cylinder
    curr := reqList[i].id

    //check inbound else error 
    if ((curr > lower) && (curr < upper)) {

      //Display current cylinder under service
      fmt.Printf("Servicing %5d\n", curr)

      //calculate the traversal distance
      traversal += Abs(curr - previousR)

      //update the previous request to current request
      previousR = curr    

    } else {
      //generate error message 
      fmt.Printf("ERROR! Cylinder Request out of bounds!\n")
    }

  }

  //print traversal time
  fmt.Printf("FCFS traversal count = %d\n", traversal)
    
}
//this didnt work so i deleted it lmao 
func sstf(lower int, upper int, initial int, reqList []Cylinder){

}


func SCAN(lower int, upper int, initial int, reqList []Cylinder) {

  //print initial outputs
  fmt.Printf("Seek algorithm: SCAN\n")
  fmt.Printf("\tLower cylinder: %5d\n", lower)
  fmt.Printf("\tUpper cylinder: %5d\n", upper)
  fmt.Printf("\tInit cylinder: %5d\n", initial)
  fmt.Printf("\tCylinder requests:\n")

  for i:=0; i<len(reqList); i++ {
    fmt.Printf("\t\tCylinder %5d\n", reqList[i].id)
  }

  //inital request 
  var currentCylinder Cylinder
  currentCylinder.id = initial
  reqList = append(reqList, currentCylinder)

  //sort cylinders by id and get index of inital
  reqList, current_index := getStartIndex(selSort(reqList), initial)
  
  //deleted intial request and reset index
  reqList = deleteReq(reqList, currentCylinder)
  reqList,_ = getStartIndex(selSort(reqList), initial)

  traversal := 0
  start_processing := false
  cylinder_services_completed := 0
  previousR := initial
  totalRequests := len(reqList)


  for i := lower; i < upper; i++ {

    //case after inital 
    curr := reqList[current_index].id

   //only start reached inital 
    if(i == initial) {
      start_processing = true
    }

    if (curr > lower && curr < upper) {

      if(start_processing) {

        //if current cylinder has not been serviced then service it and calculate the traversal distance
        if(!reqList[current_index].complete && cylinder_services_completed != totalRequests) {

          fmt.Printf("Servicing %d\n", curr)

          traversal += Abs(curr - previousR)
          
          previousR = curr

          cylinder_services_completed++
          reqList[current_index].complete = true
        

          //end of list but not finished
          if(curr == reqList[totalRequests - 1].id && cylinder_services_completed != totalRequests) {

            curr = upper

            traversal += Abs(curr - previousR)

            previousR = curr

            //reverse direction from inital
            for i := initial-1; i > lower; i-- {

              curr = reqList[current_index].id

              //service the cylinders not serviced
              if(!reqList[current_index].complete && curr > lower) {  

                fmt.Printf("Servicing %d\n", curr)

                traversal += Abs(curr - previousR)

                previousR = curr
                
                cylinder_services_completed++
                reqList[current_index].complete = true

              }

              if((current_index > 0) && cylinder_services_completed != totalRequests) {
                current_index--
              } else if (cylinder_services_completed == totalRequests) {
                break
              } else {
                current_index = 0
              }

            }


          }

          if((current_index != len(reqList) - 1) && cylinder_services_completed != totalRequests) {
            current_index++
          } else if (cylinder_services_completed == totalRequests) {
            break
          } else {
            current_index = 0
          }

        } 

      }

    } else {
      //display error message
      fmt.Printf("ERROR! Cylinder Request out of bounds!\n")
    }
  }

  //print traversal time
  fmt.Printf("SCAN traversal count = %d\n", traversal)
}

func CSCAN(lower int, upper int, initial int, reqList []Cylinder) {

  fmt.Printf("Seek algorithm: C-SCAN\n")
  fmt.Printf("\tLower cylinder: %5d\n", lower)
  fmt.Printf("\tUpper cylinder: %5d\n", upper)
  fmt.Printf("\tInit cylinder: %5d\n", initial)
  fmt.Printf("\tCylinder requests:\n")

  for i:=0; i<len(reqList); i++ {
    fmt.Printf("\t\tCylinder %5d\n", reqList[i].id)
  }

  //inital request 
  var currentCylinder Cylinder
  currentCylinder.id = initial
  reqList = append(reqList, currentCylinder)

  //sort cylinders by id and get index of inital
  reqList, current_index := getStartIndex(selSort(reqList), initial)
  
  //deleted intial request and reset index
  reqList = deleteReq(reqList, currentCylinder)
  reqList,_ = getStartIndex(selSort(reqList), initial)

  traversal := 0
  start_processing := false
  cylinder_services_completed := 0
  previousR := initial
  totalRequests := len(reqList)


  for i := lower; i < upper; i++ {

    //case after inital 
    curr := reqList[current_index].id

   //only start reached inital 
    if(i == initial) {
      start_processing = true
    }

    if (curr > lower && curr < upper) {

      if(start_processing) {

        //if current cylinder has not been serviced then service it and calculate the traversal distance
        if(!reqList[current_index].complete && cylinder_services_completed != totalRequests) {

          fmt.Printf("Servicing %d\n", curr)

          traversal += Abs(curr - previousR)
          previousR = curr

          cylinder_services_completed++
          reqList[current_index].complete = true
        

          //end of list but not finished
          if(curr == reqList[totalRequests - 1].id && cylinder_services_completed != totalRequests) {

            curr = upper

            traversal += Abs(curr - previousR)
            previousR = curr

            //starting from 0
            curr = lower
            traversal += Abs(curr - previousR)
            previousR = curr

            //reverse direction from inital
            for i := lower; i < upper; i++ {

              curr = reqList[current_index].id

              //service the cylinders not serviced
              if(!reqList[current_index].complete && curr > lower) {  

                fmt.Printf("Servicing %d\n", curr)

                traversal += Abs(curr - previousR)
                previousR = curr
                
                cylinder_services_completed++
                reqList[current_index].complete = true

              }

              if((current_index != len(reqList) - 1) && cylinder_services_completed != totalRequests) {
                current_index++
              } else if (cylinder_services_completed == totalRequests) {
                break
              } else {
                current_index = 0
              }

            }


          }

          if((current_index != len(reqList) - 1) && cylinder_services_completed != totalRequests) {
            current_index++
          } else if (cylinder_services_completed == totalRequests) {
            break
          } else {
            current_index = 0
          }

        } 

      }

    } else {

      //display error message
      fmt.Printf("ERROR! Cylinder Request out of bounds!\n")

    }


  }

  //print traversal time
  fmt.Printf("C-SCAN traversal count = %d\n", traversal)

}

func LOOK(lower int, upper int, initial int, reqList []Cylinder) {

  fmt.Printf("Seek algorithm: LOOK\n")
  fmt.Printf("\tLower cylinder: %5d\n", lower)
  fmt.Printf("\tUpper cylinder: %5d\n", upper)
  fmt.Printf("\tInit cylinder: %5d\n", initial)
  fmt.Printf("\tCylinder requests:\n")

  for i:=0; i<len(reqList); i++ {
    fmt.Printf("\t\tCylinder %5d\n", reqList[i].id)
  }

  //inital request 
  var currentCylinder Cylinder
  currentCylinder.id = initial
  reqList = append(reqList, currentCylinder)

  //sort cylinders by id and get index of inital
  reqList, current_index := getStartIndex(selSort(reqList), initial)
  
  //deleted intial request and reset index
  reqList = deleteReq(reqList, currentCylinder)
  reqList,_ = getStartIndex(selSort(reqList), initial)

  traversal := 0
  start_processing := false
  cylinder_services_completed := 0
  previousR := initial
  totalRequests := len(reqList)


  for i := lower; i < upper; i++ {

    //case after inital 
    curr := reqList[current_index].id

   //only start reached inital 
    if(i == initial) {
      start_processing = true
    }

    if (curr > lower && curr < upper) {

      if(start_processing) {

        //if current cylinder has not been serviced then service it and calculate the traversal distance
        if(!reqList[current_index].complete && cylinder_services_completed != totalRequests) {

          fmt.Printf("Servicing %d\n", curr)

          traversal += Abs(curr - previousR)
          
          previousR = curr

          cylinder_services_completed++
          reqList[current_index].complete = true
        

          //end of list but not finished
          if(curr == reqList[totalRequests - 1].id && cylinder_services_completed != totalRequests) {

            //reverse direction from inital
            for i := initial-1; i > lower; i-- {

              curr = reqList[current_index].id

              //service the cylinders not serviced
              if(!reqList[current_index].complete && curr > lower) {  

                fmt.Printf("Servicing %d\n", curr)

                traversal += Abs(curr - previousR)

                previousR = curr
                
                cylinder_services_completed++
                reqList[current_index].complete = true

              }

              if((current_index > 0) && cylinder_services_completed != totalRequests) {
                current_index--
              } else if (cylinder_services_completed == totalRequests) {
                break
              } else {
                current_index = 0
              }

            }


          }

          if((current_index != len(reqList) - 1) && cylinder_services_completed != totalRequests) {
            current_index++
          } else if (cylinder_services_completed == totalRequests) {
            break
          } else {
            current_index = 0
          }

        } 

      }

    } else {

      //display error message
      fmt.Printf("ERROR! Cylinder Request out of bounds!\n")

    }


  }

  //print traversal time
  fmt.Printf("LOCK traversal count = %d\n", traversal)

}

func CLOOK(lower int, upper int, initial int, reqList []Cylinder) {

  //print initial outputs
  fmt.Printf("Seek algorithm: C-LOOK\n")
  fmt.Printf("\tLower cylinder: %5d\n", lower)
  fmt.Printf("\tUpper cylinder: %5d\n", upper)
  fmt.Printf("\tInit cylinder: %5d\n", initial)
  fmt.Printf("\tCylinder requests:\n")

  for i:=0; i<len(reqList); i++ {
    fmt.Printf("\t\tCylinder %5d\n", reqList[i].id)
  }

  //add the initially requested cylinder in the cylinder list
  var currentCylinder Cylinder
  currentCylinder.id = initial
  reqList = append(reqList, currentCylinder)

  //sort cylinders by id and get index of inital
  reqList, current_index := getStartIndex(selSort(reqList), initial)
  
  //deleted intial request and reset index
  reqList = deleteReq(reqList, currentCylinder)
  reqList,_ = getStartIndex(selSort(reqList), initial)

  traversal := 0
  start_processing := false
  cylinder_services_completed := 0
  previousR := initial
  totalRequests := len(reqList)


  for i := lower; i < upper; i++ {

    //case after inital 
    curr := reqList[current_index].id

   //only start reached inital 
    if(i == initial) {
      start_processing = true
    }

    if (curr > lower && curr < upper) {

      if(start_processing) {

        //if current cylinder has not been serviced then service it and calculate the traversal distance
        if(!reqList[current_index].complete && cylinder_services_completed != totalRequests) {

          fmt.Printf("Servicing %d\n", curr)

          traversal += Abs(curr - previousR)
          previousR = curr

          cylinder_services_completed++
          reqList[current_index].complete = true
        

          //end of list but not finished
          if(curr == reqList[totalRequests - 1].id && cylinder_services_completed != totalRequests) {

            //reverse direction from inital
            for i := lower; i < upper; i++ {

              curr = reqList[current_index].id

              //service the cylinders not serviced
              if(!reqList[current_index].complete && curr > lower) {  

                fmt.Printf("Servicing %d\n", curr)

                traversal += Abs(curr - previousR)
                previousR = curr
                
                cylinder_services_completed++
                reqList[current_index].complete = true

              }

              if((current_index != len(reqList) - 1) && cylinder_services_completed != totalRequests) {
                current_index++
              } else if (cylinder_services_completed == totalRequests) {
                break
              } else {
                current_index = 0
              }

            }


          }

          if((current_index != len(reqList) - 1) && cylinder_services_completed != totalRequests) {
            current_index++
          } else if (cylinder_services_completed == totalRequests) {
            break
          } else {
            current_index = 0
          }

        } 

      }

    } else {

      fmt.Printf("ERROR! Cylinder Request out of bounds!\n")

    }


  }

  //print traversal time
  fmt.Printf("C-LOCK traversal count = %d\n", traversal)

}

func Abs(num int) int {
  if num < 0 {
    return -num
  }
  return num
}

func selSort (cylinders []Cylinder) (c []Cylinder) {

  c = make([]Cylinder, 0, len(cylinders))
  c = cylinders
  
  for i := 0; i < (len(c) - 1); i++ {

      minIndex := i

      for j := i+1; j < len(c); j++ {

        if (c[j].id < c[minIndex].id) {

            minIndex = j;
        }
        
        
      }
      //swap
      temp := c[minIndex]
      c[minIndex] = c[i]
      c[i] = temp
    } 

  return c
}

func getStartIndex (cylinders []Cylinder, initial int) (c []Cylinder, pos int) {

  c = make([]Cylinder, 0, len(cylinders))
  c = cylinders
  
  c[0].index = 0;
  nextIndex := c[0].index
  
  for i := 1; i < len(c) ; i++ {
    nextIndex++
    c[i].index = nextIndex; 

    if(c[i].id == initial) {
      pos = c[i].index
    }
  } 

  return c, pos
}

//remove old request 
func deleteReq (listOfCylinderRequests []Cylinder, cylinderToBeDeleted Cylinder) (c []Cylinder) {

  c = make([]Cylinder, 0, len(listOfCylinderRequests))
  c = listOfCylinderRequests

  for i:=0; i < len(c); i++ {

    if(c[i].id == cylinderToBeDeleted.id) {
      c = append(c[:i], c[i+1:]...)
      break
    }
  }

  return c
}