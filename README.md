# Crunch-go

crunch-go is go runner for crunch wordlist generator. just builds and runs the binary with given args.

## Examples

### Generate From Charset

You can use any charset from lst file.
`hex-lower, hex-upper, numeric, symbols14, symbols14-space, symbols-all, symbols-all-space, ualpha, ualpha-space, ualpha-numeric, ualpha-numeric-space, ualpha-numeric-symbol14, ualpha-numeric-symbol14-space, ualpha-numeric-all, ualpha-numeric-all-space, lalpha, lalpha-space, lalpha-numeric, lalpha-numeric-space, lalpha-numeric-symbol14, lalpha-numeric-symbol14-space, lalpha-numeric-all, lalpha-numeric-all-space, mixalpha, mixalpha-space, mixalpha-numeric, mixalpha-numeric-space, mixalpha-numeric-symbol14, mixalpha-numeric-symbol14-space, mixalpha-numeric-all, mixalpha-numeric-all-space`

```go
package main

import (
 "io/ioutil"

 "github.com/umutbasal/crunch-go"
)

func main() {
 b, err := crunch.GenerateFromCharset(5, 5, "ualpha")
 if err != nil {
  panic(err)
 }

 // write to file
 err = ioutil.WriteFile("crunch.txt", b, 0644)
 if err != nil {
  panic(err)
 }
}
```

### Custom args

```go
func GenerateFromCharset(start, end int, charset string) ([]byte, error) {
 out := fmt.Sprintf("%s/%s", tmpDir, "tmp.txt")
 params := fmt.Sprintf("%d %d -f %s/charset.lst %s -o %s", start, end, tmpDir, charset, out)
 err := crunch.Run(params)
 if err != nil {
  return nil, err
 }

 b, err := ioutil.ReadFile(out)
 if err != nil {
  return nil, err
 }
 return b, nil
}
```
