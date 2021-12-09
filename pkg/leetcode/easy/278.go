package easy

import (
	"log"
	"sort"
)

/*
	278. First Bad Version
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var isBadVersion func(version int) bool

func firstBadVersion(n int) int {

	log.Println(sort.Search(10, func(i int) bool {
		return i == 11
	}))

	return sort.Search(n, isBadVersion)
}
