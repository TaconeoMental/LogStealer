package handler

import (
    "plugin"
)


func Load(path string) (*StealerHandler, error) {
    p, err := plugin.Open(path)
    if err != nil {
        return &StealerHandler{}, err
    }

    h, err := p.Lookup("METADATA")
    if err != nil {
        return &StealerHandler{}, err
    }

    handler_s := h.(*StealerHandler)
    return handler_s, nil
}
