package day02

import "fmt"

func PilotSub(commands []string) (depth, hpos int, err error) {
	for _, command := range commands {
		var magnitude int
		if n, _ := fmt.Sscanf(command, "up %d", &magnitude); n == 1 {
			depth -= magnitude
			continue
		}
		if n, _ := fmt.Sscanf(command, "down %d", &magnitude); n == 1 {
			depth += magnitude
			continue
		}
		if n, _ := fmt.Sscanf(command, "forward %d", &magnitude); n == 1 {
			hpos += magnitude
			continue
		}
		err = fmt.Errorf("unrecognised command: %s", command)
		return
	}
	return
}

func PilotSubWithAim(commands []string) (depth, hpos int, err error) {
	var aim int
	var magnitude int
	for _, command := range commands {
		if n, _ := fmt.Sscanf(command, "up %d", &magnitude); n == 1 {
			aim -= magnitude
			continue
		}
		if n, _ := fmt.Sscanf(command, "down %d", &magnitude); n == 1 {
			aim += magnitude
			continue
		}
		if n, _ := fmt.Sscanf(command, "forward %d", &magnitude); n == 1 {
			hpos += magnitude
			depth += aim * magnitude
			continue
		}
		err = fmt.Errorf("unrecognised command: %s", command)
		return
	}
	return
}
