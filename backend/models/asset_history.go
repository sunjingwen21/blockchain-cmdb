package models

import (
	"time"

	"gorm.io/gorm"
)

// AssetHistory records changes to an asset
type AssetHistory struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	AssetID     uint           `json:"asset_id" gorm:"index;not null"`
	Asset       *Asset         `json:"asset,omitempty" gorm:"foreignKey:AssetID"`
	Action      string         `json:"action" gorm:"size:20;not null"` // created, updated, deleted, transferred
	OldData     JSONMap        `json:"old_data" gorm:"type:jsonb"`
	NewData     JSONMap        `json:"new_data" gorm:"type:jsonb"`
	ChangedByID uint           `json:"changed_by_id" gorm:"index"`
	ChangedBy   *User          `json:"changed_by,omitempty" gorm:"foreignKey:ChangedByID"`
	ChangeReason string        `json:"change_reason" gorm:"size:500"`
	BlockchainTxID string     `json:"blockchain_tx_id" gorm:"size:66"`
	CreatedAt   time.Time      `json:"created_at"`
}

// TableName specifies the table name for AssetHistory
func (AssetHistory) TableName() string {
	return "asset_history"
}

// AssetHistoryResponse is used for API responses
type AssetHistoryResponse struct {
	ID             uint      `json:"id"`
	AssetID        uint      `json:"asset_id"`
	Action         string    `json:"action"`
	OldData        JSONMap   `json:"old_data"`
	NewData        JSONMap   `json:"new_data"`
	ChangedByID    uint      `json:"changed_by_id"`
	ChangedBy      *UserResponse `json:"changed_by,omitempty"`
	ChangeReason   string    `json:"change_reason"`
	BlockchainTxID string    `json:"blockchain_tx_id"`
	CreatedAt      time.Time `json:"created_at"`
}

// ToResponse converts AssetHistory to AssetHistoryResponse
func (h *AssetHistory) ToResponse() AssetHistoryResponse {
	resp := AssetHistoryResponse{
		ID:             h.ID,
		AssetID:        h.AssetID,
		Action:         h.Action,
		OldData:        h.OldData,
		NewData:        h.NewData,
		ChangedByID:    h.ChangedByID,
		ChangeReason:   h.ChangeReason,
		BlockchainTxID: h.BlockchainTxID,
		CreatedAt:      h.CreatedAt,
	}
	if h.ChangedBy != nil {
		changedByResp := h.ChangedBy.ToResponse()
		resp.ChangedBy = &changedByResp
	}
	return resp
}
