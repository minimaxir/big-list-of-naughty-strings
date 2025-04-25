package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	in := strings.Repeat("#", 200)
	exp := "############################################################################\n"
	exp += "############################################################################\n"
	exp += "################################################"
	act := split(in, 76)
	assert.Equal(t, exp, act)
}
