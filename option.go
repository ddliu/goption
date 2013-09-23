package goption

import (
    "fmt"
    "strconv"
)

type OptionError struct {
    message string
}

func (this *OptionError) Error() string {
    return this.message
}

type Option struct {
    data map[string] interface{}
}

func NewOption(data map[string]interface{}) *Option {
    return &Option{data}
}

func (this *Option) Merge(that *Option) {
    this.MergeMap(that.Map())
}

func (this *Option) MergeMap(that map[string]interface{}) {
    for k, v := range that {
        this.data[k] = v
    }
}

func (this *Option) MergeAs(that *Option) *Option {
    return this.MergeMapAs(that.Map())
}

func (this *Option) MergeMapAs(that map[string]interface{}) *Option {
    result := NewOption(this.data)

    result.MergeMap(that)

    return result
}

func (this *Option) Map() map[string]interface{} {
    return this.data
}

func (this *Option) Set(key string, value interface{}) *Option {
    this.data[key] = value

    return this
}

func (this *Option) Get(key string) (interface{}, bool) {
    v, ok := this.data[key]

    return v, ok
}

func (this *Option) MustGet(key string) interface{} {
    v, ok := this.data[key]

    if !ok {
        panic(fmt.Sprintf("Option key does not exist %s", key))
    }

    return v
}

func (this *Option) GetBool(key string) (bool, bool) {
    v, ok := this.data[key]

    if !ok {
        return false, false
    }

    vBool, ok := v.(bool)
    if !ok {
        return false, false
    }

    return vBool, true
}

func (this *Option) MustGetBool(key string) bool {
    v, ok := this.GetBool(key)

    if !ok {
        panic(fmt.Sprintf("MustGetBool error: %s", key))
    }

    return v
}

func (this *Option) GetInt(key string) (int, bool) {
    v, ok := this.GetInt64(key)

    return int(v), ok
}

func (this *Option) MustGetInt(key string) int {
    return int(this.MustGetInt64(key))
}

func (this *Option) GetInt32(key string) (int32, bool) {
    v, ok := this.GetInt64(key)

    return int32(v), ok
}

func (this *Option) MustGetInt32(key string) int32 {
    return int32(this.MustGetInt64(key))
}

func (this *Option) GetInt64(key string) (int64, bool) {
    v, ok := this.data[key]

    if !ok {
        return 0, false
    }

    return toInt(v)
}

func (this *Option) MustGetInt64(key string) int64 {
    v, ok := this.GetInt64(key)

    if !ok {
        getKeyError(key)
    }

    return v
}

func (this *Option) GetFloat32(key string) (float32, bool) {
    v, ok := this.GetFloat64(key)

    return float32(v), ok
}

func (this *Option) MustGetFloat32(key string) float32 {
    return float32(this.MustGetFloat64(key))
}

func (this *Option) GetFloat64(key string) (float64, bool) {
    v, ok := this.data[key]

    if !ok {
        return 0, false
    }

    return toFloat(v)
}

func (this *Option) MustGetFloat64(key string) float64 {
    v, ok := this.GetFloat64(key)

    if !ok {
        getKeyError(key)
    }

    return v
}

func (this *Option) GetString(key string) (string, bool) {
    v, ok := this.data[key]

    if !ok {
        return "", false
    }

    vString, ok := v.(string)

    return vString, ok
}

func (this *Option) MustGetString(key string) string {
    v, ok := this.GetString(key)

    if !ok {
        getKeyError(key)
    }

    return v
}

func toInt(v interface{}) (int64, bool) {
    switch v := v.(type) {
    case int:
        return int64(v), true
    case int32:
        return int64(v), true
    case int64:
        return v, true
    case float32:
        return int64(v), true
    case float64:
        return int64(v), true
    case string:
        result, err := strconv.ParseInt(v, 10, 0)
        if err != nil {
            return 0, false
        }
        return result, true
    }

    return 0, false
}

func toFloat(v interface{}) (float64, bool) {
    switch v := v.(type) {
    case int:
        return float64(v), true
    case int32:
        return float64(v), true
    case int64:
        return float64(v), true
    case float32:
        return float64(v), true
    case float64:
        return v, true
    case string:
        result, err := strconv.ParseFloat(v, 64)
        if err != nil {
            return 0, false
        }
        return result, true
    }

    return 0, false
}

func getKeyError(key string) {
    panic(fmt.Sprintf("Get key error: %s", key))
}