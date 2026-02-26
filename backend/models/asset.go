package models

import (
	"time"

	"gorm.io/gorm"
)

// AssetType represents the type of asset
type AssetType string

const (
	AssetTypeHardware AssetType = "hardware"
	AssetTypeSoftware AssetType = "software"
	AssetTypeLicense  AssetType = "license"
	AssetTypeNetwork  AssetType = "network"
	AssetTypeOther    AssetType = "other"
)

// AssetStatus represents the status of an asset
type AssetStatus string

const (
	AssetStatusActive    AssetStatus = "active"
	AssetStatusInactive  AssetStatus = "inactive"
	AssetStatusMaintenance AssetStatus = "maintenance"
	AssetStatusRetired   AssetStatus = "retired"
)

// Asset represents an asset in the CMDB
type Asset struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	AssetID         string         `json:"asset_id" gorm:"uniqueIndex;size:50;not null"` // Human-readable ID
	Name            string         `json:"name" gorm:"size:200;not null"`
	Description     string         `json:"description" gorm:"type:text"`
	Type            AssetType      `json:"type" gorm:"size:20;not null"`
	Status          AssetStatus    `json:"status" gorm:"size:20;default:'active'"`
	OwnerID         uint           `json:"owner_id" gorm:"index"`
	Owner           *User          `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Location        string         `json:"location" gorm:"size:200"`
	PurchaseDate    *time.Time     `json:"purchase_date,omitempty"`
	WarrantyExpiry  *time.Time     `json:"warranty_expiry,omitempty"`
	Cost            float64        `json:"cost" gorm:"type:decimal(15,2);default:0"`
	SerialNumber    string         `json:"serial_number" gorm:"size:100"`
	BlockchainHash  string         `json:"blockchain_hash" gorm:"size:66;index"` // Ethereum tx hash
	BlockchainTxID  string         `json:"blockchain_tx_id" gorm:"size:66"`
	Metadata        JSONMap        `json:"metadata" gorm:"type:jsonb"` // Flexible JSON metadata
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Asset
func (Asset) TableName() string {
	return "assets"
}

// BeforeCreate hook is called before creating a new asset
func (a *Asset) BeforeCreate(tx *gorm.DB) error {
	if a.Status == "" {
		a.Status = AssetStatusActive
	}
	if a.AssetID == "" {
		// Generate asset ID (e.g., AST-20240226-00001)
		a.AssetID = generateAssetID()
	}
	return nil
}

// JSONMap is a helper type for JSON fields
type JSONMap map[string]interface{}

// AssetResponse is used for API responses
type AssetResponse struct {
	ID             uint      `json:"id"`
	AssetID        string    `json:"asset_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	OwnerID        uint      `json:"owner_id"`
	Owner          *UserResponse `json:"owner,omitempty"`
	Location       string    `json:"location"`
	PurchaseDate   *time.Time `json:"purchase_date,omitempty"`
	WarrantyExpiry *time.Time `json:"warranty_expiry,omitempty"`
	Cost           float64   `json:"cost"`
	SerialNumber   string    `json:"serial_number"`
	BlockchainHash string    `json:"blockchain_hash"`
	BlockchainTxID string    `json:"blockchain_tx_id"`
	Metadata       JSONMap   `json:"metadata"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ToResponse converts Asset to AssetResponse
func (a *Asset) ToResponse() AssetResponse {
	resp := AssetResponse{
		ID:             a.ID,
		AssetID:        a.AssetID,
		Name:           a.Name,
		Description:    a.Description,
		Type:           string(a.Type),
		Status:         string(a.Status),
		OwnerID:        a.OwnerID,
		Location:       a.Location,
		PurchaseDate:   a.PurchaseDate,
		WarrantyExpiry: a.WarrantyExpiry,
		Cost:           a.Cost,
		SerialNumber:   a.SerialNumber,
		BlockchainHash: a.BlockchainHash,
		BlockchainTxID: a.BlockchainTxID,
		Metadata:       a.Metadata,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}
	if a.Owner != nil {
		ownerResp := a.Owner.ToResponse()
		resp.Owner = &ownerResp
	}
	return resp
}

// generateAssetID generates a unique asset ID
func generateAssetID() string {
	return "AST-" + time.Now().Format("20060102") + "-XXXXX"
}
