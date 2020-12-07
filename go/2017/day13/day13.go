package day13

type Firewall map[int]int

func (fw Firewall) computeDamageAt(delay int, breakOnDamage bool) int {
	damage := 0
	/*
		This range operation is slow. We could speed things up by converting the
		firewall to a fixed array first, but the code is less elegant :)
	*/
	for d, r := range fw {
		d += delay
		if d%(r*2-2) == 0 {
			if breakOnDamage {
				return 1
			}
			damage += d * r
		}
	}
	return damage
}

// ComputeDamage is my solution for Part One and returns the amount of damage
// incurred by a packet traversing the firewall.
// The function returns in linear time O(n) where n is the number of scanners in
// the firewall.
func (fw Firewall) ComputeDamage() int {
	return fw.computeDamageAt(0, false)
}

// ComputeSafeDelay is my solution for Part Two and returns the first safe
// picosecond within which a packet might safely traverse the firewall without
// sustaining damage.
//
// This function returns in O(mn) where n is the number of scanners in the
// firewall, and m is picosecond delay required for safe traversal. It's super
// inefficient, but I can't image an improvement in time complexity - only
// constant time optimizations.
func (fw Firewall) ComputeSafeDelay() int {
	for i := 0; ; i++ {
		if fw.computeDamageAt(i, true) == 0 {
			return i
		}
	}
}
