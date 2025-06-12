//go:build integration
// +build integration

package api

import (
	"bytes"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

const TEST_MEMEBR_WALLET = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"

func setCommitteeMember(t *testing.T, wallet string, isApproved bool) {
	body := types.SetCommitteeMemberReq{
		MemberWallet: wallet,
		IsApproved:   &isApproved,
	}
	jsonBody, _ := json.Marshal(body)

	_ = assertPostSuccessAndWaitConfirm(t, "/committee", bytes.NewBuffer(jsonBody), "application/json", 20*time.Second)
}

func TestCommitteeMemberCreate(t *testing.T) {
	setCommitteeMember(t, TEST_MEMEBR_WALLET, true)
}

func TestCommitteeMembersList(t *testing.T) {
	u, _ := url.Parse(TEST_BASE_URL + "/committee")
	q := u.Query()
	q.Set("page", "1")
	q.Set("page_size", "10")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	apiResp := responser.Response{}
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	t.Logf("List response: %v", apiResp)

	if apiResp.Code != bizcode.SUCCESS {
		t.Errorf("Expected CODE SUCCESS, got %v", apiResp.Code)
	}
}

func TestCommitteeMemberTake(t *testing.T) {
	is_approved := false
	body := types.SetCommitteeMemberReq{
		MemberWallet: TEST_MEMEBR_WALLET,
		IsApproved:   &is_approved,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal body: %v", err)
	}

	msg := assertPostSuccessAndWaitConfirm(t, "/committee", bytes.NewBuffer(jsonBody), "application/json", 1*time.Minute)
	t.Logf("Received msg: %v", string(msg))

	wsResp := responser.ResponseRaw{}
	err = json.Unmarshal(msg, &wsResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	if wsResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected CODE SUCCESS, got %v", wsResp.Code)
	}

	var tx models.BlockchainTransaction
	err = json.Unmarshal(wsResp.Data, &tx)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	memberId := tx.EntityID
	getResp, err := http.Get(TEST_BASE_URL + "/committee/" + strconv.Itoa(int(memberId)))
	if err != nil {
		t.Fatalf("Failed to GET committee member: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("GET /committee/{id} failed, status: %d", getResp.StatusCode)
	}

	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatalf("Failed to read GET body: %v", err)
	}

	getRespJson := responser.ResponseRaw{}
	err = json.Unmarshal(getBody, &getRespJson)
	if err != nil {
		t.Fatalf("Failed to unmarshal GET response: %v", err)
	}
	if getRespJson.Code != bizcode.SUCCESS {
		t.Fatalf("GET expected SUCCESS, got %v", getRespJson.Code)
	}

	var member models.CommitteeMember
	err = json.Unmarshal(getRespJson.Data, &member)
	if err != nil {
		t.Fatalf("Failed to decode committee member: %v", err)
	}

	if member.ID != int(memberId) {
		t.Errorf("Expected ID %d, got %d", memberId, member.ID)
	}
	if member.MemberWallet != TEST_MEMEBR_WALLET {
		t.Errorf("Expected wallet %s, got %s", TEST_MEMEBR_WALLET, member.MemberWallet)
	}
}
