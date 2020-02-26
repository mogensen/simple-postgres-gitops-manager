package main

import (
	"reflect"
	"testing"
)

func Test_findDatabasesToCreate(t *testing.T) {
	type args struct {
		desiredState *State
		currentState *State
	}
	tests := []struct {
		name string
		args args
		want []Database
	}{
		{
			name: "Can handle no desired state",
			args: args{
				desiredState: &State{},
				currentState: &State{},
			},
			want: []Database{},
		},
		{
			name: "Can create new database",
			args: args{
				desiredState: &State{
					[]Database{
						{Name: "new"},
					},
				},
				currentState: &State{},
			},
			want: []Database{
				{Name: "new"},
			},
		},
		{
			name: "Can create extra database in current",
			args: args{
				desiredState: &State{
					[]Database{
						{Name: "new"},
					},
				},
				currentState: &State{
					[]Database{
						{Name: "old"},
					},
				},
			},
			want: []Database{
				{Name: "new"},
			},
		},
		{
			name: "Can create extra database in current",
			args: args{
				desiredState: &State{
					[]Database{
						{Name: "new"},
						{Name: "old"},
					},
				},
				currentState: &State{
					[]Database{
						{Name: "old"},
					},
				},
			},
			want: []Database{
				{Name: "new"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDatabasesToCreate(tt.args.desiredState, tt.args.currentState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDatabasesToCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
