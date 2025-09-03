package errorhandling

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	//configs
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}
	defer file.Close()

	cfg := &Config{}
	return cfg, nil
}

// PANIC FUNCTION
// used to signal an error that can't be handled
// we can use defer and recover to catch a panic
func triggerPanic() {
	panic("unrecoverable error")
}

func main() {
	vals := []int{1, 2, 3}
	v := vals[5] // this will cause a runtime panic: index out of range
	fmt.Println(v)
}
