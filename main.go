package main;

import (
  "fmt"
//  "math"
  "os"
  "strings"
  "bufio"
);

func ask(question string) (string, error) {
  var input string;
  var err error;
  reader := bufio.NewReader(os.Stdin);
  fmt.Print(question);
  input, err = reader.ReadString('\n');
  if err != nil {
    return "", err;
  }
  input = strings.TrimSuffix(input, "\n");
  return input, nil;
}

func main() {
  for i := 1; i < 10; i += 1 {
    fmt.Println(i);
  }
}
