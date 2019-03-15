package protocol

import (
	"reflect"
	"testing"
)

func TestPublicKeyCredentialRequestOptions_GetAllowedCredentialIDs(t *testing.T) {
	type fields struct {
		Challenge          Challenge
		Timeout            int
		RelyingPartyID     string
		AllowedCredentials []CredentialDescriptor
		UserVerification   UserVerificationRequirement
		Extensions         AuthenticationExtensions
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PublicKeyCredentialRequestOptions{
				Challenge:          tt.fields.Challenge,
				Timeout:            tt.fields.Timeout,
				RelyingPartyID:     tt.fields.RelyingPartyID,
				AllowedCredentials: tt.fields.AllowedCredentials,
				UserVerification:   tt.fields.UserVerification,
				Extensions:         tt.fields.Extensions,
			}
			if got := a.GetAllowedCredentialIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicKeyCredentialRequestOptions.GetAllowedCredentialIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}