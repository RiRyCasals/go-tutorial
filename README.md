# Golang study record

## Execution and compilation
Described in hello.go

### Run the source file
To run the source directly, enter the following command.

```
$ go run sample.go
```

### Compile and run

```
$ go build sample.go
$ ./sample
```
## print
Described in print.go

### fmt.Print

no line breaks, no space, no format.

### fmt.Println

with line breaks, with space, no format

### fmt.Printf

no line breaks, no space, with format

#### List of format types

| format | description |
| :----- | :---------- |
| %v     | default |
| %#v    | go language notation |
| %t     | boolean |
| %d     | decimal |
| %s     | string |
| %c     | charactor |
| %f, %F | |
| %e, %E | |
| %g     | automatic selection of %f or %e |
| %b     | binary |
| %o, %O | octal |
| %x, %X | hexadecimal |
| %U     | unicode |
| %p     | pointer |


