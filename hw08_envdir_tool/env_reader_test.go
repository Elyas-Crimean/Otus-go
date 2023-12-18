package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	var commonEnv Environment

	t.Run("prepare", func(t *testing.T) {
		dirEnv, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		require.Positive(t, len(dirEnv))
		commonEnv = dirEnv
	})
	t.Run("quoted", func(t *testing.T) {
		v, ok := commonEnv["HELLO"]
		assert.True(t, ok)
		assert.Equal(t, EnvValue{"\"hello\"", false}, v)
	})
	t.Run("newline", func(t *testing.T) {
		v, ok := commonEnv["FOO"]
		assert.True(t, ok)
		assert.Equal(t, EnvValue{"   foo\nwith new line", false}, v)
	})
	t.Run("zero file", func(t *testing.T) {
		v, ok := commonEnv["UNSET"]
		assert.True(t, ok)
		assert.True(t, v.NeedRemove)
	})
}
