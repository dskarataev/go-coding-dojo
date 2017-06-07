package solution

import (
	"testing"

	check "gopkg.in/check.v1"
)

type coffeequeueTestSuite struct{}

var _ = check.Suite(&coffeequeueTestSuite{})

func TestCoffeeQueue(t *testing.T) {
	check.TestingT(t)
}
