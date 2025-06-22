package constant

type CtxKey string

const(
	NewRelicTransactionCtx CtxKey = "newRelicTransaction"
)

const (
	QueueHigh   = "lender:queue:high"
	QueueMedium = "lender:queue:medium"
	QueueLow    =  "lender:queue:low"
)