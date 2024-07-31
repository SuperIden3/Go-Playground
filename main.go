package main;

import (
  "fmt"
//  "math"
  "os"
  "strings"
  "bufio"
//  "log"
  "strconv"
);

func ask(question string) string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(question)
  os.Stdout.Sync()
  input, err := reader.ReadString('\n')
  if err != nil {
    panic(err)
  }
  input = strings.TrimSuffix(input, "\n")
  return input
}

func toBytes(str string) []byte {
    byteArr := make([]byte, 0, len(str))
    for _, r := range str {
        byteArr = append(byteArr, byte(r))
    }
    return byteArr
}

func main() {
  num, _ := strconv.ParseUint(ask("Put a positive number: "), 10, 8)
  fmt.Printf("Your number is %d!\n", num);
}
