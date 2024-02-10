package handler

type checkFuncSignature = func(string) bool

type StealerHandler struct {
    HandlerName   string
    StealerName   string
    CheckFunction checkFuncSignature
}
