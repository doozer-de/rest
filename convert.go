package rest

import (
	"encoding/base64"
	"strconv"
)

func ConvertBool(value string) (bool, error) {
	if v, err := strconv.ParseBool(value); err == nil {
		return v, nil
	} else {
		return false, err
	}
}

func ConvertFloat32(value string) (float32, error) {
	if v, err := strconv.ParseFloat(value, 32); err == nil {
		return float32(v), nil
	} else {
		return 0.0, err
	}
}

func ConvertFloat64(value string) (float64, error) {
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		return float64(v), nil
	} else {
		return 0.0, err
	}
}

func ConvertString(value string) (string, error) {
	return value, nil
}

func ConvertInt(value string) (int, error) {
	if v, err := strconv.ParseInt(value, 10, 0); err == nil {
		return int(v), nil
	} else {
		return 0, err
	}
}

func ConvertInt8(value string) (int8, error) {
	if v, err := strconv.ParseInt(value, 10, 8); err == nil {
		return int8(v), nil
	} else {
		return 0, err
	}
}

func ConvertInt16(value string) (int16, error) {
	if v, err := strconv.ParseInt(value, 10, 16); err == nil {
		return int16(v), nil
	} else {
		return 0, err
	}
}

func ConvertInt32(value string) (int32, error) {
	if v, err := strconv.ParseInt(value, 10, 32); err == nil {
		return int32(v), nil
	} else {
		return 0, err
	}
}

func ConvertInt64(value string) (int64, error) {
	if v, err := strconv.ParseInt(value, 10, 64); err == nil {
		return int64(v), nil
	} else {
		return 0, err
	}
}

func ConvertUint(value string) (uint, error) {
	if v, err := strconv.ParseUint(value, 10, 0); err == nil {
		return uint(v), nil
	} else {
		return 0, err
	}
}

func ConvertUint8(value string) (uint8, error) {
	if v, err := strconv.ParseUint(value, 10, 8); err == nil {
		return uint8(v), nil
	} else {
		return 0, err
	}
}

func ConvertUint16(value string) (uint16, error) {
	if v, err := strconv.ParseUint(value, 10, 16); err == nil {
		return uint16(v), nil
	} else {
		return 0, err
	}
}

func ConvertUint32(value string) (uint32, error) {
	if v, err := strconv.ParseUint(value, 10, 32); err == nil {
		return uint32(v), nil
	} else {
		return 0, err
	}
}

func ConvertUint64(value string) (uint64, error) {
	if v, err := strconv.ParseUint(value, 10, 64); err == nil {
		return uint64(v), nil
	} else {
		return 0, err
	}
}

func ConvertBytes(value string) ([]byte, error) {
	if decoded, err := base64.URLEncoding.DecodeString(value); err == nil {
		return decoded, nil
	} else {
		return nil, err
	}
}
