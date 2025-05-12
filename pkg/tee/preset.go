package tee

// Preset KeyContexts for TEE-derived secrets across the system.

// Ethereum accounts used for contract deployment and authorized interaction.
var (
	// Used to deploy and interact with on-chain contracts as the trusted owner.
	KeyCtxTEEContractOwner = NewKeyContext(
		KEYKIND_ETH_ACC,
		"tee-eth-account",
		"contract owner",
	)
)

// Symmetric encryption keys used for data protection.
var (
	// Used to encrypt user-uploaded diagnostic reports before IPFS storage.
	KeyCtxUploadReportEncrypt = NewKeyContext(
		KEYKIND_ENC_KEY,
		"upload-report-key",
		"encrypt user test report",
	)
)
