package sample

import (
	"project-name/app/adapter/rest_api"
	structs "project-name/app/shared/structs"
	"project-name/app/usecase/sample/response"
	"reflect"
	"testing"
)

func Test_sampleUseCase_GetWeight(t *testing.T) {
	type fields struct {
		restAPI rest_api.RestAPI
	}
	type args struct {
		in structs.GetWeight
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOut      []response.WeightResponse
		wantHttpCode int
		wantCode     string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &sampleUseCase{
				restAPI: tt.fields.restAPI,
			}
			gotOut, gotHttpCode, gotCode, err := uc.GetWeight(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetWeight() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("GetWeight() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("GetWeight() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
