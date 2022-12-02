package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "container/ring"
)

// return p2 shape
func strategy(r *ring.Ring, p1 string, p2 string) string {
   if p2 == "Y" { // draw
      return p1
   }
   for {
      if r.Value == p1 {
         break
      }
      r = r.Next()
   }
   if p2 == "X" { // p2 lose
      return r.Prev().Value.(string)
   } else { // p2 win
      return r.Next().Value.(string)
   }
   return ""
}     


func main() {
   // open file
   //f, err := os.Open("testinput")
   f, err := os.Open("input")
   if err != nil {
      log.Fatal(err)
   }
   // remember to close the file at the end of the program
   defer f.Close()

   // use ring list to represent shapes and winning order
   r := ring.New(3)
   r.Value = "A" // Rock
   r = r.Next()
   r.Value = "B" // Paper
   r = r.Next()
   r.Value = "C" // Scissors
   r = r.Next()

   shape := map[string]int{"A":1,"B":2,"C":3}
    
   score_p1 := 0
   score_p2 := 0

   fileReader := bufio.NewScanner(f)
   for fileReader.Scan() {
      if err := fileReader.Err(); err != nil {
         log.Fatal(err)
      }
      round := strings.Split(fileReader.Text()," ")

      p1 := round[0]
      p2strategy := round[1]
      p2 := strategy(r,p1,p2strategy)

      // scoring for winning
      if p2strategy == "X" { // p2 lose
         score_p1 += 6
      } else if p2strategy == "Y" { // draw
         score_p1 += 3
         score_p2 += 3
      } else if p2strategy == "Z" { // p1 lose
         score_p2 += 6
      }
      // scoring for shape
      score_p1 += shape[p1]
      score_p2 += shape[p2]

   }
   fmt.Println("Score: p1:", score_p1,"p2:",score_p2)
};
 
