package benchmarkingandprofiling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var benchText = `In publishing and graphic design is available to why this.`

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(benchText)
		require.Equal(b, 10, len(tokens))
	}
}

/** RUNNING FROM TERMINAL
$ go test -bench . -run NONE
$ go test -bench . -run NONE -cpuprofile cpu.prof
$ go tool pprof -http=:8080 cpu.prof (brew install graphviz)
**/