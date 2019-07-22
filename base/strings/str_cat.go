package base

import (
	"fmt"
	"strconv"
	"strings"
)

func StrCat(parts ...interface{}) string {
	var builder strings.Builder

	for _, part := range parts {
		switch part.(type) {
		case bool: // Note this is different from C++, where bool == 0 or 1.
			builder.WriteString(strconv.FormatBool(part.(bool)))
		case float32:
			appendFloat(&builder, float64(part.(float32)))
		case float64:
			appendFloat(&builder, part.(float64))
		case byte:
			builder.WriteByte(part.(byte))
		case rune:
			builder.WriteRune(part.(rune))
		case int:
			appendInt(&builder, int64(part.(int)))
		case int8:
			appendInt(&builder, int64(part.(int8)))
		case int16:
			appendInt(&builder, int64(part.(int16)))
		// int32 is handled as rune.
		case int64:
			appendInt(&builder, int64(part.(int64)))
		case uint:
			appendUint(&builder, uint64(part.(uint)))
		// uint8 is handled as byte.
		case uint16:
			appendUint(&builder, uint64(part.(uint16)))
		case uint32:
			appendUint(&builder, uint64(part.(uint32)))
		case uintptr:
			appendUint(&builder, uint64(part.(uintptr)))
		case uint64:
			appendUint(&builder, part.(uint64))
		case []byte:
			builder.Write(part.([]byte))
		case string:
			builder.WriteString(part.(string))
		default:
			fmt.Printf("Param %T[%v] is not supported\n", part, part)
		}
	}
	return builder.String()
}

func appendFloat(builder *strings.Builder, d float64) {
	builder.WriteString(strconv.FormatFloat(d, 'g', 6, 64))
}

func appendInt(builder *strings.Builder, i int64) {
	builder.WriteString(strconv.FormatInt(i, 10))
}

func appendUint(builder *strings.Builder, ui uint64) {
	builder.WriteString(strconv.FormatUint(ui, 10))
}
