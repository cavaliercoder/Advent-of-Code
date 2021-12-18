package day12

func CountPaths1(m map[string][]string) (pathCount int) {
	path := make([]string, 0, 64)
	checkFunc := func(path []string, e string) bool {
		if e == "start" {
			return false
		}
		if e == "end" {
			return true
		}
		if e[0] >= 'A' && e[0] <= 'Z' {
			return true
		}
		if !contains(path, e) {
			return true
		}
		return false
	}
	return dfs(m, path, "start", checkFunc)
}

func CountPaths2(m map[string][]string) (pathCount int) {
	path := make([]string, 0, 64)
	checkFunc := func(path []string, e string) bool {
		if e == "start" {
			return false
		}
		if e == "end" {
			return true
		}
		if e[0] >= 'A' && e[0] <= 'Z' {
			return true
		}
		if !contains(path, e) {
			return true
		}
		if !hasDuplicate(path) {
			return true
		}
		return false
	}
	return dfs(m, path, "start", checkFunc)
}

func dfs(
	m map[string][]string,
	path []string,
	v string,
	checkFunc func(path []string, cave string) bool,
) (pathCount int) {
	if v == "end" {
		return 1
	}
	path = append(path, v)
	for _, e := range m[v] {
		if checkFunc(path, e) {
			pathCount += dfs(m, path, e, checkFunc)
		}
	}
	path = path[:len(path)-1]
	return
}

func contains(a []string, v string) (ok bool) {
	for _, s := range a {
		if v == s {
			return true
		}
	}
	return false
}

func hasDuplicate(a []string) bool {
	m := make(map[string]struct{}, len(a))
	for _, s := range a {
		if len(s) > 0 && s[0] >= 'A' && s[0] <= 'Z' {
			continue
		}
		if _, ok := m[s]; ok {
			return true
		}
		m[s] = struct{}{}
	}
	return false
}
