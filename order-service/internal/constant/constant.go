package constant

type CtxKey string

const(
	NewRelicTransactionCtx CtxKey = "newRelicTransaction"
)

const (
	QueueHigh   = "order:queue:high"
	QueueMedium = "order:queue:medium"
	QueueLow    = "order:queue:low"
)