package mongo

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name   string
	Age    int
	Thenga int `tt:"manga"`
}

func TestStructToMap(t *testing.T) {
	type args struct {
		in        interface{}
		tag       string
		withValue bool
	}
	tests := []struct {
		name    string
		args    args
		wantM   map[string]interface{}
		wantErr bool
	}{
		{"valid struct",
			args{TestStruct{"sarath", 1, 1}, "", true},
			map[string]interface{}{"Name": "sarath", "Age": 1, "Thenga": 1},
			false,
		},
		{"valid struct with empty field",
			args{&TestStruct{"sarath", 0, 1}, "", true},
			map[string]interface{}{"Name": "sarath", "Thenga": 1},
			false,
		},
		{"pointer to a valid struct with empty field",
			args{&TestStruct{"sarath", 0, 1}, "", true},
			map[string]interface{}{"Name": "sarath", "Thenga": 1},
			false,
		},

		{"valid struct with empty field and with value set to false",
			args{TestStruct{"sarath", 0, 1}, "", false},
			map[string]interface{}{"Name": "sarath", "Age": 0, "Thenga": 1},
			false,
		},
		{"valid struct with tag passed ",
			args{TestStruct{"sarath", 0, 1}, "tt", false},
			map[string]interface{}{"Name": "sarath", "Age": 0, "manga": 1},
			false,
		},
		{"in valid struct - not a strcut",
			args{"Just a string", "tt", false},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := StructToMap(tt.args.in, tt.args.tag, tt.args.withValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("StructToMap() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
