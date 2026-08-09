package displayTools

func TruncateString(str string, num int, dots bool) string {

	if len(str) > num-3 && dots {
		return str[:num-3] + "..."
	} else if len(str) > num && !dots {
		return str[:num]
	}

	return str
}
