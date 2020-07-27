# Partition

parition/ab package contains an implementation to partition your users to some groups for A/B testing.

## Installation

~~~~
go get "github.com/dakimura/partition/ab"
~~~~

Or you can manually git clone the repository to
`$(go env GOPATH)/src/github.com/dakimura/partition`.

## Usage

```
package example

import "github.com/dakimura/partition/ab"

func example() {
	abTest, _ := ab.NewTest(
		[]ab.Group{
			{ID: "default", TargetPercentage: 50,},
			{ID: "GroupA", TargetPercentage: 25,},
			{ID: "GroupB", TargetPercentage: 25,},
		},
	)

	userID := "Kimura Takuya"
	group := abTest.GetGroup(userID)
	switch group {
	case "GroupA":
		print("behavior 1")
	case "GroupB":
		print("behavior 2")
	default:
		print("default behavior")
	}
}
```
