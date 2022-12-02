package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
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

   score_p1 := 0
   score_p2 := 0

   t := map[string]string{"A":"Rock","B":"Paper","C":"Scissors","X":"Rock","Y":"Paper","Z":"Scissors"}
   shape := map[string]int{"Rock":1,"Paper":2,"Scissors":3}

   // read csv values using csv.Reader
   fileReader := bufio.NewScanner(f)
   for fileReader.Scan() {
      // do something with read line
      if err := fileReader.Err(); err != nil {
         log.Fatal(err)
      }
      round := strings.Split(fileReader.Text()," ")

      p1 := t[round[0]]
      p2 := t[round[1]]
      switch {
         case p1 == p2:
            // draw
            score_p1 += 3
            score_p2 += 3
         case (p1 == "Rock" && p2 == "Scissors") || (p1 == "Scissors" && p2 == "Paper") || (p1 == "Paper" && p2 == "Rock"):
            score_p1 += 6
         case (p2 == "Rock" && p1 == "Scissors") || (p2 == "Scissors" && p1 == "Paper") || (p2 == "Paper" && p1 == "Rock"):
             score_p2 += 6
      }
      score_p1 += shape[p1]
      score_p2 += shape[p2]
   }
   fmt.Println("Score: p1:", score_p1," p2:",score_p2)
};
 
