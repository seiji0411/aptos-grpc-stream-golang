package grpcurl

const (
	TRANSACTION_TYPE_UNSPECIFIED      = "TRANSACTION_TYPE_UNSPECIFIED"
	TRANSACTION_TYPE_GENESIS          = "TRANSACTION_TYPE_GENESIS"
	TRANSACTION_TYPE_BLOCK_METADATA   = "TRANSACTION_TYPE_BLOCK_METADATA"
	TRANSACTION_TYPE_STATE_CHECKPOINT = "TRANSACTION_TYPE_STATE_CHECKPOINT"
	TRANSACTION_TYPE_USER             = "TRANSACTION_TYPE_USER"
	TRANSACTION_TYPE_UNRECOGNIZED     = "UNRECOGNIZED"
)

const (
	CHANGE_TYPE_UNSPECIFIED       = "TYPE_UNSPECIFIED"
	CHANGE_TYPE_DELETE_MODULE     = "TYPE_DELETE_MODULE"
	CHANGE_TYPE_DELETE_RESOURCE   = "TYPE_DELETE_RESOURCE"
	CHANGE_TYPE_DELETE_TABLE_ITEM = "TYPE_DELETE_TABLE_ITEM"
	CHANGE_TYPE_WRITE_MODULE      = "TYPE_WRITE_MODULE"
	CHANGE_TYPE_WRITE_RESOURCE    = "TYPE_WRITE_RESOURCE"
	CHANGE_TYPE_WRITE_TABLE_ITEM  = "TYPE_WRITE_TABLE_ITEM"
	CHANGE_TYPE_UNRECOGNIZED      = "UNRECOGNIZED"
)

const (
	TRANSACTION_PAYLOAD_TYPE_UNSPECIFIED            = "TYPE_UNSPECIFIED"
	TRANSACTION_PAYLOAD_TYPE_ENTRY_FUNCTION_PAYLOAD = "TYPE_ENTRY_FUNCTION_PAYLOAD"
	TRANSACTION_PAYLOAD_TYPE_SCRIPT_PAYLOAD         = "TYPE_SCRIPT_PAYLOAD"
	TRANSACTION_PAYLOAD_TYPE_MODULE_BUNDLE_PAYLOAD  = "TYPE_MODULE_BUNDLE_PAYLOAD"
	TRANSACTION_PAYLOAD_TYPE_WRITE_SET_PAYLOAD      = "TYPE_WRITE_SET_PAYLOAD"
	TRANSACTION_PAYLOAD_TYPE_MULTISIG_PAYLOAD       = "TYPE_MULTISIG_PAYLOAD"
	TRANSACTION_PAYLOAD_UNRECOGNIZED                = "UNRECOGNIZED"
)

type Transaction struct {
	Timestamp   Timestamp       `json:"timestamp"`
	Version     string          `json:"version"`
	Info        TransactionInfo `json:"info"`
	Epoch       string          `json:"epoch"`
	BlockHeight string          `json:"blockHeight"`
	TxType      string          `json:"type"`
	// BlockMetadata   BlockMetadataTransaction   `json:"blockMetadata"`
	// Genesis         GenesisTransaction         `json:"genesis"`
	// StateCheckpoint StateCheckpointTransaction `json:"stateCheckpoint"`
	User UserTransaction `json:"user"`
}
type Timestamp struct {
	/**
	 * Represents seconds of UTC time since Unix epoch
	 * 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	 * 9999-12-31T23:59:59Z inclusive.
	 */
	Seconds string `json:"seconds"`
	/**
	 * Non-negative fractions of a second at nanosecond resolution. Negative
	 * second values with fractions must still have non-negative nanos values
	 * that count forward in time. Must be from 0 to 999,999,999
	 * inclusive.
	 */
	Nanos uint64 `json:"nanos"`
}

type TransactionInfo struct {
	Hash                string           `json:"hash"`
	StateChangeHash     string           `json:"stateChangeHash"`
	EventRootHash       string           `json:"eventRootHash"`
	StateCheckpointHash string           `json:"stateCheckpointHash"`
	GasUsed             string           `json:"gasUsed"`
	Success             bool             `json:"success"`
	VmStatus            string           `json:"vmStatus"`
	AccumulatorRootHash string           `json:"accumulatorRootHash"`
	Changes             []WriteSetChange `json:"changes"`
}

type WriteSetChange struct {
	ChangeType string `json:"type"`
	// DeleteModule    DeleteModule    `json:"deleteModule"`
	// DeleteResource  DeleteResource  `json:"deleteResource"`
	// DeleteTableItem DeleteTableItem `json:"deleteTableItem"`
	// WriteModule     WriteModule     `json:"writeModule"`
	WriteResource  WriteResource  `json:"writeResource"`
	WriteTableItem WriteTableItem `json:"writeTableItem"`
}

//	type DeleteModule struct {
//		Address      string       `json:"address"`
//		StateKeyHash string       `json:"stateKeyHash"`
//		Module       MoveModuleId `json:"module"`
//	}
//
//	type DeleteResource struct {
//		Address      string        `json:"address"`
//		StateKeyHash string        `json:"stateKeyHash"`
//		ResourceType MoveStructTag `json:"type"`
//		TypeStr      string        `json:"typeStr"`
//	}
//
//	type DeleteTableItem struct {
//		StateKeyHash string          `json:"stateKeyHash"`
//		Handle       string          `json:"handle"`
//		Key          string          `json:"key"`
//		Data         DeleteTableData `json:"data"`
//	}
//
//	type DeleteTableData struct {
//		Key     string `json:"key"`
//		KeyType string `json:"keyType"`
//	}
//
//	type WriteModule struct {
//		Address      string             `json:"address"`
//		StateKeyHash string             `json:"stateKeyHash"`
//		Data         MoveModuleBytecode `json:"data"`
//	}
type WriteResource struct {
	Address      string        `json:"address"`
	StateKeyHash string        `json:"stateKeyHash"`
	ResourceType MoveStructTag `json:"type"`
	TypeStr      string        `json:"typeStr"`
	Data         string        `json:"data"`
}
type WriteTableData struct {
	Key       string `json:"key"`
	KeyType   string `json:"keyType"`
	Value     string `json:"value"`
	ValueType string `json:"valueType"`
}
type WriteTableItem struct {
	StateKeyHash string         `json:"stateKeyHash"`
	Handle       string         `json:"handle"`
	Key          string         `json:"key"`
	Data         WriteTableData `json:"data"`
}
type TransactionPayload struct {
	PayloadType string `json:"type"`
	// EntryFunctionPayload EntryFunctionPayload `json:"entryFunctionPayload"`
	// ScriptPayload        ScriptPayload        `json:"scriptPayload"`
	// ModuleBundlePayload  ModuleBundlePayload  `json:"moduleBundlePayload"`
	// WriteSetPayload      WriteSetPayload      `json:"writeSetPayload"`
	// MultisigPayload      MultisigPayload      `json:"multisigPayload"`
}

// type BlockMetadataTransaction struct {
// 	id?: string | undefined;
//     round?: bigint | undefined;
//     events?: Event[] | undefined;
//     previousBlockVotesBitvec?: Uint8Array | undefined;
//     proposer?: string | undefined;
//     failedProposerIndices?: number[] | undefined;
// }

// type GenesisTransaction struct {
// 	payload?: WriteSet | undefined;
//     events?: Event[] | undefined;
// }

// type StateCheckpointTransaction struct {
// }

type UserTransaction struct {
	Request UserTransactionRequest `json:"request"`
	Events  Event                  `json:"events"`
}

type MoveModuleId struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type MoveStructTag struct {
	Address string `json:"address"`
	Module  string `json:"module"`
	Name    string `json:"name"`
	// GenericTypeParams MoveType `json:"genericTypeParams"`
}

type UserTransactionRequest struct {
	Sender                  string             `json:"sender"`
	SequenceNumber          string             `json:"sequenceNumber"`
	MaxGasAmount            string             `json:"maxGasAmount"`
	GasUnitPrice            string             `json:"gasUnitPrice"`
	ExpirationTimestampSecs Timestamp          `json:"expirationTimestampSecs"`
	Payload                 TransactionPayload `json:"payload"`
	// Signature               Signature          `json:"signature"`
}

// type Signature struct {
//     type?: Signature_Type | undefined;
//     ed25519?: Ed25519Signature | undefined;
//     multiEd25519?: MultiEd25519Signature | undefined;
//     multiAgent?: MultiAgentSignature | undefined;
//     feePayer?: FeePayerSignature | undefined;
//     secp256k1Ecdsa?: Secp256k1ECDSASignature | undefined;
//     singleSender?: SingleSender | undefined;
// }

type Event struct {
	Key            EventKey `json:"key"`
	SequenceNumber string   `json:"sequenceNumber"`
	EventType      MoveType `json:"type"`
	TypeStr        string   `json:"typeStr"`
	Data           string   `json:"data"`
}

type EventKey struct {
	CreationNumber string `json:"creationNumber"`
	AccountAddress string `json:"accountAddress"`
}

type MoveType struct {
	MoveType string `json:"type"`
	// Vector                MoveType      `json:"vector"`
	MoveStruct            MoveStructTag `json:"struct"`
	GenericTypeParamIndex int           `json:"genericTypeParamIndex"`
	Reference             string        `json:"reference"`
	Unparsable            string        `json:"unparsable"`
}
