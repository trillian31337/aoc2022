package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "sort"
)

 func main() {
   // open file
   //f, err := os.Open("testinput")
   f, err := os.Open("input")
   if err != nil {
      log.Fatal(err)
   }
   // remember to close the file at the end of the program
   defer f.Close()

   // read csv values using csv.Reader
   fileReader := bufio.NewScanner(f)
   elfs := make([]int,1)
   elfindex := 0
   for fileReader.Scan() {
      // do something with read line
      if err := fileReader.Err(); err != nil {
         log.Fatal(err)
      }
      if fileReader.Text() == "" {
   	 elfindex += 1
	 elfs = append(elfs,0)
      }
      cal,_ := strconv.Atoi(fileReader.Text())
      //fmt.Println(cal)
      elfs[elfindex] += cal
   }
   fmt.Println(elfs)
   sort.Ints(elfs)
   fmt.Println("Calories:", elfs[len(elfs)-1]+elfs[len(elfs)-2]+elfs[len(elfs)-3])
};
 
