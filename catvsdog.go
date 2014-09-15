// Author: Errol Burger

package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	testCases, cats, dogs, numOfVoters int
	answers                            []int
)

type Voter struct {
	likes                  string
	votedFor, votedAgainst int
	collision              bool
	checked                bool
}

type Voters struct {
	all []Voter
}

var voters = new(Voters)

// Gameshow data
func assignData() (int, int, int) {
	// cats, dogs, numOfVoters
	var c, d, n int

	fmt.Scanf("%d %d %d", &c, &d, &n)

	// Check values against gameshow limitations
	switch {
	case c > 100 || d > 100 || c < 1 || d < 1:
		os.Exit(0)
	case n > 500 || n < 0:
		os.Exit(0)
	}
	return c, d, n
}

func runGameshow() {
	// Provide test cases
	var t int
	fmt.Scanf("%d", &t)
	testCases = t

	if testCases > 100 || testCases < 0 {
		os.Exit(0)
	} else {
		for testCases > 0 {
			// Clear voters from previous block
			voters.all = nil
			runCase()
			testCases--
		}
	}
}

func castVotes() {
	// Create new voters and append to voter block
	for numOfVoters > 0 {
		var vote1, vote2 string

		fmt.Scanf("%s %s", &vote1, &vote2)

		vLikes := vote1[:1]
		vFor, _ := strconv.Atoi(vote1[1:len(vote1)])
		vAgainst, _ := strconv.Atoi(vote2[1:len(vote1)])

		v := Voter{
			likes: vLikes,
		}
		v.votedFor, v.votedAgainst = vFor, vAgainst
		voters.all = append(voters.all, v)
		numOfVoters--
	}
}

func maximumViewers(voters *Voters) int {
	collisions := 0

	// Loop through voter block, once for every voter, using i as a starting
	// point. Flags are used to check whether a collision (apposing votes) took
	// place, flagging primary voter as collision true. To avoid double checks
	// on inner vote, a second flag (checked) is used for the inner loop.
	for i, voter := range voters.all {
		if voter.collision == false {
			for x := i; x < len(voters.all); x++ {
				if voter.votedFor == voters.all[x].votedAgainst &&
					voter.likes != voters.all[x].likes &&
					voters.all[i].collision != true &&
					voters.all[x].checked != true {

					voters.all[i].collision = true
					voters.all[x].checked = true
					collisions++
				} else if voter.votedAgainst == voters.all[x].votedFor &&
					voter.likes != voters.all[x].likes &&
					voters.all[i].collision != true &&
					voters.all[x].checked != true {

					voters.all[i].collision = true
					voters.all[x].checked = true
					collisions++
				}
			}
		}
	}
	// Maximum possible number of satisfied voters for the show
	// Answer:  total voters minus the number of unique collisions
	return (len(voters.all) - collisions)

}

// Run test case
func runCase() {
	cats, dogs, numOfVoters = assignData()
	castVotes()
	answers = append(answers, maximumViewers(voters))
}

func main() {
	runGameshow()

	// print out answers
	for _, answer := range answers {
		fmt.Println(answer)
	}
}
