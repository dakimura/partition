package ab

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
)

type GroupID = string

type Test struct {
	groups      map[GroupID]Group
	modAllocate map[uint32]GroupID
}

type Group struct {
	ID GroupID
	// 0~100
	TargetPercentage int
}

func NewTest(groups []Group) (*Test, error) {
	percentageSum := 0
	for _, group := range groups {
		percentageSum += group.TargetPercentage
	}
	//if percentageSum > 100 {
	//	return nil, errors.New(
	//		fmt.Sprintf("Sum of TargetPercentage for each Group should be 100 or under. got=%d", percentageSum),
	//	)
	//}
	if percentageSum != 100 {
		return nil, errors.New(
			fmt.Sprintf("Sum of TargetPercentage for each Group should be 100. got=%d", percentageSum),
		)
	}

	cnt := uint32(0)
	groups2 := map[GroupID]Group{}
	groupMod := map[uint32]string{}
	for i, group := range groups {
		group := group
		groups[i] = group
		for k := 0; k < group.TargetPercentage; k++ {
			groupMod[cnt] = group.ID
			cnt++
		}
	}
	return &Test{
		groups:      groups2,
		modAllocate: groupMod,
	}, nil
}

func (t *Test) GetGroup(targetID string) GroupID {
	b := []byte(targetID)
	// https://stackoverflow.com/questions/13521252/a-b-test-partition-function
	md5sum := md5.Sum(b)

	mod := binary.BigEndian.Uint32(md5sum[12:16]) % 100

	return t.modAllocate[mod]
}
