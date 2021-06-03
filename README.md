# Golang study record


## Execution and compilation
Described in hello.go

### run the source file
To run the source directly, enter the following command.

```bash:
$ go run sample.go
```

### compile and run

```bash:
$ go build sample.go
$ ./sample
```


## Print
Described in print.go

### fmt.Print

no line breaks, no space, no format.

### fmt.Println

with line breaks, with space, no format

### fmt.Printf

no line breaks, no space, with format

#### format types

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


## Variable
Described in variable.go

### var
var name type

```go:
var box int
box = 123
```

can be omitted as follows.

```go:
var box int = 123
```

if the type is clear

```go:
var box = 123
```

if you write without using var, it will be as follows.

```go:
box := 123
box1, box2 := 123, "string"
```

but, this doesn't work

```go:
box int := 123
```

### constant
const name

```go:
const BOX = 123
```

Usually the type is omitted.


##
