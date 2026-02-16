## [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)

This tutorial introduces the basics of fuzzing in Go. With fuzzing, random data is run against your test in an attempt to find vulnerabilities or crash-causing inputs. Some examples of vulnerabilities that can be found by fuzzing are SQL injection, buffer overflow, denial of service and cross-site scripting attacks.

We will write a fuzz test for a simple function, run the go command, and debug and fix issues in the code.

For help with terminology throughout this tutorial, see the [Go Fuzzing glossary](https://go.dev/doc/security/fuzz/#glossary).

We will progress through the following sections:

- Create a folder for your code.
- Add code to test.
- Add a unit test.
- Add a fuzz test.
- Fix two bugs.
- Explore additional resources.

Note: Go fuzzing currently supports a subset of built-in types, listed in the [Go Fuzzing docs](https://go.dev/doc/security/fuzz/#requirements), with support for more built-in types to be added in the future.

### Add a fuzz test

The unit test has limitations, namely that each input must be added to the test by the developer. One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.

Run FuzzReverse with fuzzing, to see if any randomly generated string inputs will cause a failure. This is executed using go test with a new flag, -fuzz, set to the parameter Fuzz. Copy the command below.

```
go test -fuzz=Fuzz ./utils
```

```
go test -fuzz=FuzzReverse ./utils
```

A failure occurred while fuzzing, and the input that caused the problem is written to a seed corpus file that will be run the next time go test is called, even without the -fuzz flag. To view the input that caused the failure, open the corpus file written to the testdata/fuzz/FuzzReverse directory in a text editor. Your seed corpus file may contain a different string, but the format will be the same.

```
go test fuzz v1 ./utils
```