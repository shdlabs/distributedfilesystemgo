package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestNewFileServer(t *testing.T) {
	assert := assert.New(t)

	fileServer := NewFileServer(FileServerOpts{})

	assert.IsType(new(FileServer), fileServer)
}

type ExampleTestSute struct {
	suite.Suite
	VarliableThatShouldStartAtFive int
}

func (suite *ExampleTestSute) SetupTest() {
	suite.VarliableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSute) TestExample() {
	assert.Equal(suite.T(), 5, suite.VarliableThatShouldStartAtFive)
}

func TestSuteRun(t *testing.T) {
	suite.Run(t, new(ExampleTestSute))
}
