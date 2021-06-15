// Package proverb outputs the full text of a proverbial rhyme.
package proverb

import "fmt"

// Proverb generate the relevant proverb.
func Proverb(rhyme []string) []string {
	var ret []string
	for i := range rhyme {
		if i > 0 {
			ret = append(ret, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
		}
		if i == len(rhyme)-1 {
			ret = append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		}
	}
	return ret
}
