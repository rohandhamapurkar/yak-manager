To use this yak shop manager using cli or http
First build the go packages

```
$ go build ./cmd/cli
$ go build ./cmd/api
```

To use the cli to generate output

```
$ ./cli -h
Usage of C:\Users\rohan\Golang\upwork-assessment-1\cli.exe:
  -days int
        number of elapsed days, default value 0
  -file string
        full json file path

$ ./cli -file=C:\Users\rohan\Golang\upwork-assessment-1\input.json -days 13
Output for T = 13 :
In Stock:
        1104.480 liters of milk
        3 skins of wool
Herd:
        Betty-1 4.13 years old
        Betty-2 8.13 years old
        Betty-3 9.63 years old
```


To use the http api to generate output

```
$ ./api -h
Usage of C:\Users\rohan\Golang\upwork-assessment-1\api.exe:
  -file string
        full json file path

$ ./api -file=C:\Users\rohan\Golang\upwork-assessment-1\input.json
Running on :8000

# in another terminal execute
$ curl http://localhost:8000/yak-shop/herd/13
{"herd":[{"name":"Betty-1","age":4.13,"sex":"f","AgeLastShaved":4},{"name":"Betty-2","age":8.13,"sex":"f","AgeLastShaved":8},{"name":"Betty-3","age":9.63,"sex":"f","AgeLastShaved":9.5}]}

$ curl http://localhost:8000/yak-shop/stock/13
{"milk":1104.48,"skins":3}
```

For executing tests run the following command
```
$ go test -v ./pkg/
=== RUN   TestPrintStocks
--- PASS: TestPrintStocks (0.00s)
=== RUN   TestComputeStockAndHerdWith13DaysElapsed
--- PASS: TestComputeStockAndHerdWith13DaysElapsed (0.00s)
=== RUN   TestComputeStockAndHerdWith14DaysElapsed
--- PASS: TestComputeStockAndHerdWith14DaysElapsed (0.00s)
=== RUN   TestJsonParser
--- PASS: TestJsonParser (0.00s)
PASS
ok      upwork-assessment-1/pkg (cached)
```