# GOLANG TEST DRIVEN DEVELOPMENT DOCUMENTATION
This documentation will guide you to understand TDD with go step-by-step


### SUB-TOPIC
- Tl-TD
- Test file
- Test function
- Example
- Benchmark


## TL-TD
to run tdd in your app, you need to know about this in the first place

##### Convention Name
- file test name end with "_test.go"
- Put the file in the same package as the one being tested
- function name must be "TestXxx" where Xxx does not start with a lowercase letter

##### Command
- `$ go test` //normal test in current dir
- `$ go test -v` //test function in current dir one-by-one
- `$ go test -run <func_name>` //test specific function
- `$ go test <package_name>` //test specific package


## Test file

first step, clone this project in your local directory and make sure you already install go

`$ git clone https://gitlab.com/ibrohhm/go-tdd-doc.git`

in this project, I made Add function and some test function, check this project is running well out by typing `$ go test` in your console,
if you want to test your test function one-by-one, just type `$ go test -v` in your console (v for verbose)


## Test function

In general `$ go test -run <func_name>` fill `<func_name>` with test function name that you want to test, example

run `$ go test -run TestAddTwoNumbers` if you want to test just TestAddTwoNumbers function 


## Example
the word "Example" doesn't mean example, but it's a parts of sub TDD with go. I still don't get it why this features was made, I'l tell you what it is anyway

in TDD go, Example function check the output from function (so far, I only know Example function just check output from fmt.Println). 

go doc say that:
- Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run.
- The comment prefix "Unordered output:" is like "Output:", but matches any line order, and 
- Example functions without output comments are compiled but not executed.

just try example function in folder example in this project, and run `$ go test -v` hope you'll understand

## Benchmark

Benchmark function are executed by the `$ go test` command when its -bench flag us provided,
bencmark function must run the taget cobe b.N times and benchmark function convention name "BenchmarkXxx"

go to "benchmark" folder and run command `$ go test -bench`, you'll see output in console like this

```
BenchmarkHello-8   	20000000	        77.3 ns/op
BenchmarkPerm-8    	 5000000	       259 ns/op
```

means that the loop n BenchmarkHello ran 20000000 times at a speed of 77.3 ns per loop, likewise BenchmarkPerm 

## Other

- `go test -cover` //to see test coverage
- `go test -coverprofile=coverage.out; go tool cover -html=coverage.out` //this will open a html in browser with showing which lines of code are covered in green (and not covered with red) (sample image below)
- `go test -v -short` // variable `testing.Short()` and `testing.Verbose()` has true value, mostly short flag used for skip function with long test

## Notes:
- `t.Error()` = `t.Log()`+`t.Fail()` //t.Fail() fails the current test but continues with rest of the tests
- `t.Fatal()` = `t.Log()`+`t.FailNow()` // t.Fatal() and t.FailNow() fails & stops ALL the tests


source:
- https://golang.org/pkg/testing/
- https://github.com/gunjan5/go-test-driven-development
