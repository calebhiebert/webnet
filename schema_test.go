package web_net

import (
	"encoding/json"
	"fmt"
	"github.com/r3labs/diff/v2"
	"reflect"
	"testing"
)

func TestGenerateSchema(t *testing.T) {
	sch := GenerateSchema(GetJSONData())
	js, _ := json.Marshal(sch)

	fmt.Println(string(js))

	d, _ := diff.Diff(GetJSONData(), GetJSONData2())

	fmt.Println(d)
}

func TestGenerateSchema1(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		{
			name: "number",
			args: args{5},
			want: &Schema{
				Type: "integer",
			},
		},
		{
			name: "float",
			args: args{4.20},
			want: &Schema{
				Type: "float",
			},
		},
		{
			name: "string",
			args: args{"heyo"},
			want: &Schema{
				Type: "string",
			},
		},
		{
			name: "bool",
			args: args{true},
			want: &Schema{
				Type: "bool",
			},
		},
		{
			name: "object",
			args: args{struct {
				Foolish int
				Little  string
				Program *struct {
					Test int
					Ing  *float64
				}
			}{}},
			want: &Schema{
				Type: "object",
				Fields: []Schema{
					{
						Name: strptr("Foolish"),
						Type: "integer",
					},
					{
						Name: strptr("Little"),
						Type: "string",
					},
					{
						Name: strptr("Program"),
						Type: "object",
						Fields: []Schema{
							{
								Name: strptr("Test"),
								Type: "integer",
							},
							{
								Name: strptr("Ing"),
								Type: "float",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSchema(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
