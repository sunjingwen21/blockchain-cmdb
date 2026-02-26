package models

import (
	"testing"
	"time"
)

func TestAssetToResponse(t *testing.T) {
	now := time.Now()
	purchaseDate := now.AddDate(-1, 0, 0)
	warrantyDate := now.AddDate(2, 0, 0)

	owner := &User{
		ID:       1,
		Username: "owner",
		Email:    "owner@example.com",
		Role:     "admin",
	}

	asset := Asset{
		ID:             1,
		AssetID:        "AST-20240226-00001",
		Name:           "Test Server",
		Description:    "Production server",
		Type:           AssetTypeHardware,
		Status:         AssetStatusActive,
		OwnerID:        owner.ID,
		Owner:          owner,
		Location:       "Data Center A",
		PurchaseDate:   &purchaseDate,
		WarrantyExpiry: &warrantyDate,
		Cost:           5000.00,
		SerialNumber:   "SN123456",
		BlockchainHash: "0xabc123...",
		Metadata: JSONMap{
			"cpu":    "Intel Xeon",
			"memory": "64GB",
		},
		CreatedAt: now,
	}

	response := asset.ToResponse()

	if response.ID != asset.ID {
		t.Errorf("Expected ID %d, got %d", asset.ID, response.ID)
	}

	if response.AssetID != asset.AssetID {
		t.Errorf("Expected AssetID '%s', got '%s'", asset.AssetID, response.AssetID)
	}

	if response.Name != asset.Name {
		t.Errorf("Expected Name '%s', got '%s'", asset.Name, response.Name)
	}

	if response.Type != string(asset.Type) {
		t.Errorf("Expected Type '%s', got '%s'", asset.Type, response.Type)
	}

	if response.Status != string(asset.Status) {
		t.Errorf("Expected Status '%s', got '%s'", asset.Status, response.Status)
	}

	if response.Owner == nil {
		t.Error("Expected Owner to be included in response")
	} else if response.Owner.ID != owner.ID {
		t.Errorf("Expected Owner ID %d, got %d", owner.ID, response.Owner.ID)
	}

	if response.Cost != asset.Cost {
		t.Errorf("Expected Cost %.2f, got %.2f", asset.Cost, response.Cost)
	}

	if response.Metadata["cpu"] != "Intel Xeon" {
		t.Error("Expected metadata to include CPU info")
	}
}

func TestAssetTableName(t *testing.T) {
	asset := Asset{}
	if asset.TableName() != "assets" {
		t.Errorf("Expected table name 'assets', got '%s'", asset.TableName())
	}
}

func TestAssetTypes(t *testing.T) {
	tests := []struct {
		assetType AssetType
		expected  string
	}{
		{AssetTypeHardware, "hardware"},
		{AssetTypeSoftware, "software"},
		{AssetTypeLicense, "license"},
		{AssetTypeNetwork, "network"},
		{AssetTypeOther, "other"},
	}

	for _, tt := range tests {
		if string(tt.assetType) != tt.expected {
			t.Errorf("Expected asset type '%s', got '%s'", tt.expected, tt.assetType)
		}
	}
}

func TestAssetStatuses(t *testing.T) {
	tests := []struct {
		status   AssetStatus
		expected string
	}{
		{AssetStatusActive, "active"},
		{AssetStatusInactive, "inactive"},
		{AssetStatusMaintenance, "maintenance"},
		{AssetStatusRetired, "retired"},
	}

	for _, tt := range tests {
		if string(tt.status) != tt.expected {
			t.Errorf("Expected status '%s', got '%s'", tt.expected, tt.status)
		}
	}
}

func TestAssetToResponseWithoutOwner(t *testing.T) {
	asset := Asset{
		ID:      1,
		AssetID: "AST-20240226-00002",
		Name:    "Test Asset",
		Type:    AssetTypeSoftware,
		Status:  AssetStatusActive,
		Owner:   nil,
	}

	response := asset.ToResponse()

	if response.Owner != nil {
		t.Error("Expected Owner to be nil when not loaded")
	}
}
