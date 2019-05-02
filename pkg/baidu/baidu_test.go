package baidu

import (
	"reflect"
	"testing"

	"github.com/wentaojia2014/Transer/pkg/transer"
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
			name: "baidu test",
			args: args{
				input: &transer.TransInput{
					ID:     "20190502000293463",
					Secret: "0d2RvCho9XZNEO5GCGNs",
					Query:  "Hello",
					To:     Zh,
				},
			},
			want: &transer.TransOutput{
				Result: "你好",
			},
		},
		{
			name: "baidu test",
			args: args{
				input: &transer.TransInput{
					ID:     "20190502000293463",
					Secret: "0d2RvCho9XZNEO5GCGNs",
					Query:  "世界",
					To:     En,
				},
			},
			want: &transer.TransOutput{
				Result: "world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trans(tt.args.input); !reflect.DeepEqual(got.Result, tt.want.Result) {
				t.Errorf("Trans() = %v, want %v", got.Result, tt.want.Result)
			}
		})
	}
}
