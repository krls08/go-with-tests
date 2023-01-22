package main

type colorType int32

const (
	colorType_COLOR_UNKNOWN colorType = 0
	colorType_COLOR_RED     colorType = 1
	colorType_COLOR_BLUE    colorType = 2

// MissionType_CONTINGENCY_LoG MissionType = 7
)

// Enum value maps for MissionType.
var (
	colorType_name = map[int32]string{
		0: "COLOR_UNKNOWN",
		1: "COLOR_RED",
		2: "COLOR_BLUE",
	}
	colorType_value = map[string]int32{
		"COLOR_UNKNOWN": 0,
		"COLOR_RED":     1,
		"COLOR_BLUE":    2,
	}
)

type me struct {
	color colorType
}

func main() {
	color := colorType_COLOR_BLUE

	m := me{}
	m.color = color

}
