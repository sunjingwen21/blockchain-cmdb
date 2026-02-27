package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents an Ethereum blockchain client
type Client struct {
	client  *ethclient.Client
	chainID *big.Int
}

// Config holds blockchain configuration
type Config struct {
	RPCURL         string
	PrivateKey     string
	ContractAddress string
}

// AssetRecord represents an asset on the blockchain
type AssetRecord struct {
	ID          string
	Name        string
	Owner       string
	Status      string
	CreatedAt   int64
	UpdatedAt   int64
	HistoryHash string
}

// NewClient creates a new blockchain client
func NewClient(cfg *Config) (*Client, error) {
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	log.Printf("Connected to Ethereum network (Chain ID: %d)", chainID)

	return &Client{
		client:  client,
		chainID: chainID,
	}, nil
}

// Close closes the blockchain connection
func (c *Client) Close() {
	c.client.Close()
}

// GetChainID returns the chain ID
func (c *Client) GetChainID() *big.Int {
	return c.chainID
}

// GetBalance gets the balance of an address
func (c *Client) GetBalance(address string) (*big.Int, error) {
	addr := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetLatestBlockNumber gets the latest block number
func (c *Client) GetLatestBlockNumber() (uint64, error) {
	header, err := c.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// RegisterAsset registers an asset on the blockchain (placeholder for smart contract)
func (c *Client) RegisterAsset(asset *AssetRecord) (string, error) {
	// This is a placeholder implementation
	// In production, this would interact with a smart contract
	log.Printf("Registering asset on blockchain: %s", asset.ID)
	
	// Generate a mock transaction hash
	txHash := crypto.Keccak256Hash([]byte(asset.ID + fmt.Sprintf("%d", asset.CreatedAt)))
	return txHash.Hex(), nil
}

// GetTransactionStatus gets the status of a transaction
func (c *Client) GetTransactionStatus(txHash string) (bool, error) {
	hash := common.HexToHash(txHash)
	_, isPending, err := c.client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return false, err
	}
	return !isPending, nil
}

// GenerateKeyPair generates a new Ethereum key pair
func GenerateKeyPair() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", fmt.Errorf("failed to cast public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return fmt.Sprintf("%x", privateKeyBytes), address, nil
}