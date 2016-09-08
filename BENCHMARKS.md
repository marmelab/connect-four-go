# Benchmarks results

```
docker run --rm --volume="`pwd`:/srv" -ti marmelab-go bash -c	"cd src/connectfour && go test -run=XXX -bench=."
PASS
BenchmarkNextBestMove	       1	1162464620 ns/op
BenchmarkGuessNextBoards	    5000	    627253 ns/op
ok  	connectfour	4.376s
``
