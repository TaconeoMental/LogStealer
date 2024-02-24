package stealerlog

type StealerLogHandler struct {
    HandlerName string
    StealerName string
    RuleSet     []Rule
}

func testRule(rule Rule, path string) bool {
    rule.setBaseDir(path)
    if !rule.CheckRoot() && !rule.IsOptional() {
        return false
    }
    // Rule either passed or is optional
    return true
}

func (slh *StealerLogHandler) Test(path string) bool {
    for _, rule := range slh.RuleSet {
        if !testRule(rule, path) {
            return false
        }
    }
    return true
}

