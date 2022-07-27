package checker

import (
	"mtt/tools/mocks"
	"strings"
	"testing"
)

func TestCheckTokenOk(t *testing.T) {
	var (
		token        = "test_token"
		scope        = "test_scope"
		expectedResp = "client_id: test_client_id\nclient_type: 13\nexpires_in: 1296000\nuser_id: 42\nusername: test_username\n"
	)

	mockConnection := &mocks.MockTCPConnection{Success: true}
	checker := Checker{Connection: mockConnection}

	resp, err := checker.CheckToken(token, scope)

	if err != nil {
		t.Fatalf("invalid token checking. %s", err.Error())
	}

	if strings.Compare(expectedResp, resp) != 0 {
		t.Fatalf("responses not equal. expected: %s, but actual: %s", expectedResp, resp)
	}
}

func TestCheckTokenError(t *testing.T) {
	var (
		token        = "test_token"
		scope        = "test_scope"
		expectedResp = "error: CUBE_OAUTH2_ERR_BAD_SCOPE\nmessage: bad scope\n"
	)

	mockConnection := &mocks.MockTCPConnection{Success: false}
	checker := Checker{Connection: mockConnection}

	resp, err := checker.CheckToken(token, scope)

	if err != nil {
		t.Fatalf("invalid token checking. %s", err.Error())
	}

	if strings.Compare(expectedResp, resp) != 0 {
		t.Fatalf("responses not equal. expected: %s, but actual: %s", expectedResp, resp)
	}
}
