package sample

import (
	"project-name/app/adapter/rest_api"
	structs "project-name/app/shared/structs"
	"project-name/app/usecase/sample/response"
	"reflect"
	"testing"
	"time"
)

func Test_sampleUseCase_GetWeight_1(t *testing.T) {
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
			tt.args.in.Limit = 10
			tt.args.in.Offset = 0
			gotOut, gotHttpCode, gotCode, err := uc.GetWeight(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeight(1) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetWeight(1) gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("GetWeight(1) gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("GetWeight(1) gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_GetWeight_2(t *testing.T) {
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
			tt.args.in.Limit = 10
			tt.args.in.Offset = 10
			gotOut, gotHttpCode, gotCode, err := uc.GetWeight(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeight(2) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetWeight(2) gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("GetWeight(2) gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("GetWeight(2) gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_GetWeight_3(t *testing.T) {
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
			tt.args.in.Limit = 10
			tt.args.in.Offset = 20
			gotOut, gotHttpCode, gotCode, err := uc.GetWeight(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeight(3) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetWeight(3) gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("GetWeight(3) gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("GetWeight(3) gotCode = %v, want %v", gotCode, tt.wantCode)
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
		wantOut      response.WeightResponse
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
			*tt.args.in.WeightRecordId = "1"
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
			tt.args.in.Min = 50
			tt.args.in.Max = 55
			tt.args.in.Date = time.Now().Format("2006-01-02")
			tt.args.in.UserId = "Test Case"
			tt.args.in.Source = "Test Case"
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
			tt.args.in.Min = 50
			tt.args.in.Max = 65
			tt.args.in.Date = time.Now().Format("2006-01-02")
			tt.args.in.UserId = "Test Case"
			tt.args.in.Source = "Test Case"
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
