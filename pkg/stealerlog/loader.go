package stealerlog

import "plugin"


func Load(path string) (*StealerLogHandler, error) {
    p, err := plugin.Open(path)
    if err != nil {
        return &StealerLogHandler{}, err
    }

    h, err := p.Lookup("HANDLER")
    if err != nil {
        return &StealerLogHandler{}, err
    }

    handler_s := h.(*StealerLogHandler)
    return handler_s, nil
}
