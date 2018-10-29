package meshinfo

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewLookupHandler(t *testing.T) {
	tests := []struct {
		name string
		want *LookupHandler
	}{
		{"new", &LookupHandler{
			internalIDCounter: 1,
			lookup:            map[reflect.Type]MeshInfo{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLookupHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLookupHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupHandler_AddInformation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMesh := NewMockMeshInfo(mockCtrl)
	h := NewLookupHandler()
	herr := NewLookupHandler()
	herr.internalIDCounter = maxInternalID
	type args struct {
		info MeshInfo
	}
	tests := []struct {
		name               string
		h                  *LookupHandler
		args               args
		wantErr            bool
		expectedInternalID uint64
	}{
		{"1", h, args{mockMesh}, false, 1},
		{"2", h, args{mockMesh}, false, 2},
		{"3", h, args{mockMesh}, false, 3},
		{"max", herr, args{mockMesh}, true, maxInternalID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.info.(*MockMeshInfo).EXPECT().InfoType().Return(reflect.TypeOf(""))
			tt.args.info.(*MockMeshInfo).EXPECT().setInternalID(tt.expectedInternalID)
			if err := tt.h.AddInformation(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("LookupHandler.AddInformation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLookupHandler_InfoTypes(t *testing.T) {
	tests := []struct {
		name string
		h    *LookupHandler
		want []reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.InfoTypes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupHandler.InfoTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupHandler_AddFace(t *testing.T) {
	type args struct {
		newFaceCount uint32
	}
	tests := []struct {
		name    string
		h       *LookupHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddFace(tt.args.newFaceCount); (err != nil) != tt.wantErr {
				t.Errorf("LookupHandler.AddFace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLookupHandler_GetInformationByType(t *testing.T) {
	type args struct {
		infoType reflect.Type
	}
	tests := []struct {
		name  string
		h     *LookupHandler
		args  args
		want  MeshInfo
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.h.GetInformationByType(tt.args.infoType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupHandler.GetInformationByType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LookupHandler.GetInformationByType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLookupHandler_GetInformationCount(t *testing.T) {
	tests := []struct {
		name string
		h    *LookupHandler
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.GetInformationCount(); got != tt.want {
				t.Errorf("LookupHandler.GetInformationCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupHandler_AddInfoFromTable(t *testing.T) {
	type args struct {
		otherHandler     Handler
		currentFaceCount uint32
	}
	tests := []struct {
		name    string
		h       *LookupHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddInfoFromTable(tt.args.otherHandler, tt.args.currentFaceCount); (err != nil) != tt.wantErr {
				t.Errorf("LookupHandler.AddInfoFromTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLookupHandler_CloneFaceInfosFrom(t *testing.T) {
	type args struct {
		faceIndex      uint32
		otherHandler   Handler
		otherFaceIndex uint32
	}
	tests := []struct {
		name string
		h    *LookupHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.CloneFaceInfosFrom(tt.args.faceIndex, tt.args.otherHandler, tt.args.otherFaceIndex)
		})
	}
}

func TestLookupHandler_ResetFaceInformation(t *testing.T) {
	type args struct {
		faceIndex uint32
	}
	tests := []struct {
		name string
		h    *LookupHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.ResetFaceInformation(tt.args.faceIndex)
		})
	}
}

func TestLookupHandler_RemoveInformation(t *testing.T) {
	type args struct {
		infoType reflect.Type
	}
	tests := []struct {
		name string
		h    *LookupHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.RemoveInformation(tt.args.infoType)
		})
	}
}

func TestLookupHandler_PermuteNodeInformation(t *testing.T) {
	type args struct {
		faceIndex  uint32
		nodeIndex1 uint32
		nodeIndex2 uint32
		nodeIndex3 uint32
	}
	tests := []struct {
		name string
		h    *LookupHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.PermuteNodeInformation(tt.args.faceIndex, tt.args.nodeIndex1, tt.args.nodeIndex2, tt.args.nodeIndex3)
		})
	}
}
