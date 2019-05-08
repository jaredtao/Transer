package youdao

import (
	"reflect"
	"testing"

	"github.com/jaredtao/Transer/services/transer"
)

func TestTrans(t *testing.T) {
	type args struct {
		input *transer.TransInput
	}
	tests := []struct {
		name string
		args args
		want *transer.TransOutput
	}{
		{
			name: "youdao test 1",
			args: args{
				input: &transer.TransInput{
					ID:     "1bd659586c52ea1d",
					Secret: "5ZktXhHfLCpI0KnAdcxx4cPyGJwcVXaV",
					Query:  "Hello",
					To:     Zh,
				},
			},
			want: &transer.TransOutput{
				Result: "你好",
			},
		},
		// {
		// 	name: "youdao test 2",
		// 	args: args{
		// 		input: &transer.TransInput{
		// 			ID:     "1bd659586c52ea1d",
		// 			Secret: "5ZktXhHfLCpI0KnAdcxx4cPyGJwcVXaV",
		// 			Query:  "好",
		// 			To:     En,
		// 		},
		// 	},
		// 	want: &transer.TransOutput{
		// 		Result: "good",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trans(tt.args.input); !reflect.DeepEqual(got.Result, tt.want.Result) {
				t.Errorf("Trans() = %v, want %v", got.Result, tt.want.Result)
			}
		})
	}
}
