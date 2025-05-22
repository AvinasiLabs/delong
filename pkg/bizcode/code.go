package bizcode

type Code string

const SUCCESS Code = "SUCCESS"

const (
	BAD_REQUEST           Code = "BAD_REQUEST"           // Invalid request parameters
	UNAUTHORIZED          Code = "UNAUTHORIZED"          // Authentication required
	FORBIDDEN             Code = "FORBIDDEN"             // Permission denied
	NOT_FOUND             Code = "NOT_FOUND"             // Resource not found
	TOO_MANY_REQUESTS     Code = "TOO_MANY_REQUESTS"     // Rate limit exceeded
	INTERNAL_SERVER_ERROR Code = "INTERNAL_SERVER_ERROR" // Unhandled server error
)

const (
	DB_WRITE_FAIL Code = "DB_WRITE_FAIL"
	DB_READ_FAIL  Code = "DB_READ_FAIL"
	DB_DUPLICATE  Code = "DB_DUPLICATE_ENTRY"

	IPFS_UPLOAD_FAIL  Code = "IPFS_UPLOAD_FAIL"
	MINIO_UPLOAD_FAIL Code = "MINIO_UPLOAD_FAIL"
	MYSQL_WRITE_FAIL  Code = "MYSQL_WRITE_FAIL"
	MYSQL_READ_FAIL   Code = "MYSQL_READ_FAIL"
)

const (
	ETHEREUM_CALL_FAIL  Code = "ETHEREUM_CALL_FAIL"
	RECEIPT_QUERY_FAIL  Code = "RECEIPT_RECEIVE_FAIL"
	BLOCK_QUERY_FAIL    Code = "BLOCK_QUERY_FAIL"
	BLOCK_NOT_CONFIRMED Code = "BLOCK_NOT_CONFIRMED"
)

const (
	AlGO_LINK_INVALID Code = "AlGO_LINK_INVALID"
)

const (
	FILE_READ_FAIL   Code = "FILE_READ_FAIL"
	FILE_UPLOAD_FAIL Code = "FILE_UPLOAD_FAIL"
)

const (
	REPORT_ALREADY_EXIST    Code = "REPORT_ALREADY_EXIST"
	REPORT_ANALYZE_FAIL     Code = "REPORT_ANALYZE_FAIL"
	REPORT_DESERIALIZE_FAIL Code = "REPORT_DESERIALIZE_FAIL"
)

const (
	KEY_DERIVE_FAIL Code = "KEY_DERIVE_FAIL"
)
