package bios

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpsUnmarshal(t *testing.T) {
	in := `{"op":"system.setcode","label":"Test","data":{"contract_name_ref": "mama"}}`
	var opType OperationType
	require.NoError(t, json.Unmarshal([]byte(in), &opType))
	assert.Equal(t, "mama", opType.Data.(*OpSetCode).ContractNameRef)
	// fmt.Printf("woah: %T %#v\n", opType.Data, opType.Data)
}
