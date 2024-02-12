package stealerlog

type checkFuncSignature = func(string) bool

type StealerLogHandler struct {
    HandlerName   string
    StealerName   string
    CheckFunction checkFuncSignature
}

