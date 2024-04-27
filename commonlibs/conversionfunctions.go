package commonlibs

// ConversionAnyToInt преобразовывает любое числовое значение в int или
// возвращает 0 если было получено не число
func ConversionAnyToInt(i interface{}) int {
	switch customType := i.(type) {
	case int:
		return customType

	case int16:
		return int(customType)

	case int32:
		return int(customType)

	case int64:
		return int(customType)

	case uint:
		return int(customType)

	case uint16:
		return int(customType)

	case uint32:
		return int(customType)

	case uint64:
		return int(customType)

	case float32:
		return int(customType)

	case float64:
		return int(customType)

	}

	return 0
}
