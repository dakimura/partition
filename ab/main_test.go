package ab

import (
	"testing"
)

func TestTest_GetGroup(t1 *testing.T) {
	tests := []struct {
		name     string
		groups   []Group
		targetID string
		want     GroupID
	}{
		{
			name: "normal",
			groups: []Group{
				{ID: "Group0", TargetPercentage: 100},
			},
			targetID: "Kato Rosa",
			want:     "Group0",
		},
		{
			name: "normal",
			groups: []Group{
				{ID: "Group0", TargetPercentage: 30},
				{ID: "Group1", TargetPercentage: 30},
				{ID: "Group2", TargetPercentage: 40},
			},
			targetID: "foobar", // I already know this user is grouped to Group2 by a manual test...
			want:     "Group2",
		},
	}
	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			test, _ := NewTest(tt.groups)
			if got := test.GetGroup(tt.targetID); got != tt.want {
				t1.Errorf("GetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTest(t *testing.T) {
	tests := []struct {
		name    string
		groups  []Group
		wantErr bool
	}{
		{
			name: "Success",
			groups: []Group{
				{
					ID:               "Group1",
					TargetPercentage: 50,
				},
				{
					ID:               "Group2",
					TargetPercentage: 50,
				},
			},
			wantErr: false,
		},
		{
			name: "Only 40% is the target scope",
			groups: []Group{
				{
					ID:               "Group1",
					TargetPercentage: 20,
				},
				{
					ID:               "Group2",
					TargetPercentage: 20,
				},
			},
			wantErr: true,
		},
		{
			name: "Sum of TargetPercentage is over 100",
			groups: []Group{
				{
					ID:               "Group1",
					TargetPercentage: 60,
				},
				{
					ID:               "Group2",
					TargetPercentage: 60,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTest(tt.groups)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
