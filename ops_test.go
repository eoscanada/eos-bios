package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpsUnmarshal(t *testing.T) {
	in := `{"op":"system.setcode","label":"Test","data":{}}`
	var opType OperationType
	require.NoError(t, json.Unmarshal([]byte(in), &opType))
	assert.Equal(t, opType.Data, "mama")
	fmt.Printf("woah: %T %#v\n", opType.Data, opType.Data)
}
