/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/10/8 18:18
 */

package tools

import (
	"reflect"
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	type args struct {
		str    string
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "t1",
			args: args{str: "2019-10-08 18:17:42", format: "Y-m-d H:i:s"},
			want: time.Now(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToTime(tt.args.str, tt.args.format)

			t.Log(time.Now().Unix())
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeFormat(t *testing.T) {
	type args struct {
		format string
		time   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{format: "Y-m-d H:i:s", time: time.Now()},
			want: "2019-10-08 18:23:22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeFormat(tt.args.format, tt.args.time); got != tt.want {
				t.Errorf("TimeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
