package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("test is valid", func(t *testing.T) {
		expectedEnv := Environment{
			"ADDED": EnvValue{Value: "added env", NeedRemove: false},
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
		}

		env, err := ReadDir("testdata/env")

		fmt.Println(env)

		require.NoError(t, err)
		require.Equal(t, expectedEnv, env)
	})
}
