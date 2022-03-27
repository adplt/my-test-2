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

func Test_sampleUseCase_GetWeightById(t *testing.T) {
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
			gotOut, gotHttpCode, gotCode, err := uc.GetWeightById(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeightById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetWeightById() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("GetWeightById() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("GetWeightById() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_AddWeight(t *testing.T) {
	type fields struct {
		restAPI rest_api.RestAPI
	}
	type args struct {
		in structs.WeightData
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
			gotOut, gotHttpCode, gotCode, err := uc.AddWeight(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AddWeight() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("AddWeight() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("AddWeight() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_UpdateWeight(t *testing.T) {
	type fields struct {
		restAPI rest_api.RestAPI
	}
	type args struct {
		in structs.WeightData
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
			gotOut, gotHttpCode, gotCode, err := uc.UpdateWeight(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("UpdateWeight() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("UpdateWeight() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("UpdateWeight() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
