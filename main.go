package main;

import (
  "fmt"
//  "math"
  "os"
  "strings"
  "bufio"
//  "log"
);

func ask(question string) (string, error) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(question)
  os.Stdout.Sync()
  input, err := reader.ReadString('\n')
  if err != nil {
    return "", err
  }
  input = strings.TrimSuffix(input, "\n")
  return input, nil
}

func toBytes(str string) []byte {
    byteArr := make([]byte, 0, len(str))
    for _, r := range str {
        byteArr = append(byteArr, byte(r))
    }
    return byteArr
}

func main() {
  input, _ := ask("Type something: ")
  var arr []byte = toBytes(input)
  for i := 0; i < len(arr); i++ {
    fmt.Printf("\"%c\" = %d\n", input[i], arr[i]);
  }
}
