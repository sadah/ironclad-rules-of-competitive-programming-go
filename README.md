# ironclad-rules-of-competitive-programming-go

This repository is personal practices of [競技プログラミングの鉄則 アルゴリズム力と思考力を高める77の技術](https://www.amazon.co.jp/dp/483997750X).

The original support site is here.

* https://github.com/E869120/kyopro-tessoku

The tasks in AtCoder are here.

* https://atcoder.jp/contests/tessoku-book

## Require

atcoder-tools

* https://kyuridenamida.github.io/atcoder-tools/
* https://github.com/kyuridenamida/atcoder-tools/issues/263


Generated codes using following command.

```sh
atcoder-tools gen tessoku-book --lang=go
```


## Usage

### Make

```
Usage:
        make build DIR=<build_directory>        Run test scripts (Specify the directory).
        make test DIR=<test_directory>          Run test scripts (Specify the directory).
        make submit DIR=<submit_directory>      Run test scripts (Specify the directory).
```

```
❯ make build DIR=./src/001
Build directory: ./src/001
[Main Program]
compile command:  go build -o main main.go
Compiling... (command: `go build -o main main.go`)
```

```
❯ make test DIR=./src/001
make: Circular test <- test dependency dropped.
Test directory: ./src/001
# in_1.txt ... PASSED 1 ms
# in_2.txt ... PASSED 1 ms
# in_3.txt ... PASSED 1 ms
Passed all test cases!!!
```

### Manual

Compile

```
atcoder-tools compile
```

Test

```
atcoder-tools test
```

Submit

```
atcoder-tools submit -u
```
