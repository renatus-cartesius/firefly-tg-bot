package firefly

import "testing"

func TestAccountsResponse_String(t *testing.T) {
	type fields struct {
		Data []AccountsData
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &AccountsResponse{
				Data: tt.fields.Data,
			}
			if gotRes := ar.String(); gotRes != tt.wantRes {
				t.Errorf("AccountsResponse.String() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
