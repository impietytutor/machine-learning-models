package models

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetRandomSeed() int {
	return int(math.Round(math.Random()*100000))
}

func FloatSliceToString(floats []float64) string {
	var sb strings.Builder
	for i := range floats {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.FormatFloat(floats[i], 'f', 8, 64))
	}
	return sb.String()
}

func StringToFloatSlice(str string) ([]float64, error) {
	parts := strings.Split(str, ",")
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty string")
	}
	var floats []float64
	for _, part := range parts {
		f, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float: %w", err)
		}
		floats = append(floats, f)
	}
	return floats, nil
}

func GetLogger() *log.Logger {
	return log.New(log.Writer(), "", log.Ldate|log.Ltime)
}