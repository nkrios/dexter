package engine_test

import (
	"github.com/coinbase/dexter/engine"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestInvestigatorCreationSetsName(t *testing.T) {
	assert := assert.New(t)

	name := "alice"

	investigator, _, err := engine.NewInvestigator(name, "password")
	assert.Nil(err)
	assert.Equal(name, investigator.Name)
}

func TestInvestigatorCanMakeString(t *testing.T) {
	assert := assert.New(t)

	investigator, _, err := engine.NewInvestigator("bob", "password")
	assert.Nil(err)
	_, err = investigator.String()
	assert.Nil(err)
}
