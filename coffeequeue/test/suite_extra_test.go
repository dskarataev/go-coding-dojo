package test

import (
	"testing"

	check "gopkg.in/check.v1"
)

type extraTestSuite struct{}

var _ = check.Suite(&extraTestSuite{})

func TestExtra(t *testing.T) {
	check.TestingT(t)
}
