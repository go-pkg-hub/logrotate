package logrotate

import (
    "regexp"
    "strconv"
)

func StringToSize(s string) int64 {
    re := regexp.MustCompile(`(?P<val>[\d.]+)(?P<unit>[a-z]*)`)
    matches := re.FindAllStringSubmatch(s, -1)
    for i := 0; i < len(matches); i++ {
        val, _ := strconv.ParseFloat(matches[i][1], 64)
        if val < 0 {
            val = 0
        }

        switch unit := matches[i][2]; unit {
        case "k", "kb":
            return int64(val * 1024)
        case "m", "mb":
            return int64(val * 1024 * 1024)
        case "g", "gb":
            return int64(val * 1024 * 1024 * 1024)
        default:
            return int64(val)
        }
    }

    return 0
}
