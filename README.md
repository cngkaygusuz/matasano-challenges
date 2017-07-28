# matasano-challenges
My solutions to the problems posted at http://cryptopals.com/


## Where are the solutions ?
The code is not organized around the challenges itself. Some are in tests, some are solved through dedicated solvers,
with my primary concern being maximising code re-usage as much as possible. If you want to just see the parts where
the challenge is actually tackled, I marked the points-of-interest by a comment like

> // Matasano <challenge_number>

You can search the code to find them. My favorite way to do so is by the program `ag`

```shell
ag Matasano
```


## Running the Code
Following instructions are written for unixy operating systems that has golang toolchain set up.

Steps for installing golang is available [here.](https://golang.org/doc/install)

If the solution is implemented through a test, cd to the directory the file is present, and run `go test -v`

```bash
cd repeating-otp
go test -v
```

If the solution is implemented by a stand-alone program, cd to the directory the source file is present, and run `go run`

```
cd cmd/ch8
go run ch8
```


#AN IMPORTANT NOTE
This code is NOT cryptographically safe and it will not provide you security you are trying to establish. You are
most welcome inspecting the code for educational purposes and playing with it, but do not, >> DO NOT << use this code
in an environment that is even remotely trying to be secure. Always use well-tested and trusted public libraries for
this purpose.