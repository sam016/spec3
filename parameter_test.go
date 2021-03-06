package spec3

import (
	"reflect"
	"testing"
)

func TestParameters_Get(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForParameters()}, args{"skipParam"}, &Parameter{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForParameters()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Parameters{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameters.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameters_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantParameter *Parameter
		wantOK        bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForParameters()}, args{"limitParam"}, &Parameter{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForParameters()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Parameters{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantParameter) {
				t.Errorf("Parameters.GetOK() got = %v, want %v", got, tt.wantParameter)
			}
			if got1 != tt.wantOK {
				t.Errorf("Parameters.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestParameters_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Parameter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForParameters()}, args{"getParam", &Parameter{Description: "Getting Parameters"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForParameters()}, args{"limitParam", &Parameter{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForParameters()}, args{"", &Parameter{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Parameters{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("Parameters.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("Parameters.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("Parameters.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestParameters_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Parameter) error
	}
	type foundParameter struct {
		parameter *Parameter
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundParameter
		wantErr          error
	}{
		{
			"Should iterate 4 items for Parameters fixture",
			fields{buildOrderMapForParameters()},
			map[string]*foundParameter{
				"skipParam":  &foundParameter{&Parameter{Description: "default parameter"}, false},
				"limitParam": &foundParameter{&Parameter{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in Parameters",
			fields{},
			map[string]*foundParameter{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Parameters{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotParameter *Parameter) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotParameter) {
						t.Fatalf("Parameters.ForEach() for key = %s val = %v, want = %v", key, gotParameter, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("Parameters.ForEach() for key = %s val = %v, want = %v", key, gotParameter, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("Parameters.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("Parameters.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestParameters_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for Parameters fixture", fields{buildOrderMapForParameters()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in Parameters", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Parameters{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("Parameters.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForParameters() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForParameters() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &Parameter{Description: "default parameter"},
			"limitParam": &Parameter{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildParametersFixture() Parameters {
	m := Parameters{
		data: buildOrderMapForParameters(),
	}

	return m
}
