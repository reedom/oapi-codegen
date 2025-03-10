package externalref

import (
	"testing"

	"github.com/stretchr/testify/require"

	packageA "github.com/reedom/oapi-codegen/internal/test/externalref/packageA"
	packageB "github.com/reedom/oapi-codegen/internal/test/externalref/packageB"
)

func TestParameters(t *testing.T) {
	b := &packageB.ObjectB{}
	_ = Container{
		ObjectA: &packageA.ObjectA{ObjectB: b},
		ObjectB: b,
	}
}

func TestGetSwagger(t *testing.T) {
	_, err := packageB.GetSwagger()
	require.Nil(t, err)

	_, err = packageA.GetSwagger()
	require.Nil(t, err)

	_, err = GetSwagger()
	require.Nil(t, err)
}
