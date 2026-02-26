// Type definitions for Blockchain CMDB

export interface User {
  id: number;
  username: string;
  email: string;
  role: 'user' | 'admin';
  is_active: boolean;
  created_at: string;
}

export interface UserCreateRequest {
  username: string;
  email: string;
  password: string;
  role?: 'user' | 'admin';
}

export interface UserUpdateRequest {
  username?: string;
  email?: string;
  role?: 'user' | 'admin';
  is_active?: boolean;
}

export type AssetType = 'hardware' | 'software' | 'license' | 'network' | 'other';
export type AssetStatus = 'active' | 'inactive' | 'maintenance' | 'retired';

export interface Asset {
  id: number;
  asset_id: string;
  name: string;
  description?: string;
  type: AssetType;
  status: AssetStatus;
  owner_id: number;
  owner?: User;
  location?: string;
  purchase_date?: string;
  warranty_expiry?: string;
  cost: number;
  serial_number?: string;
  blockchain_hash?: string;
  blockchain_tx_id?: string;
  metadata?: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}

export interface AssetCreateRequest {
  name: string;
  description?: string;
  type: AssetType;
  location?: string;
  purchase_date?: string;
  warranty_expiry?: string;
  cost?: number;
  serial_number?: string;
  metadata?: Record<string, unknown>;
}

export interface AssetUpdateRequest {
  name?: string;
  description?: string;
  status?: AssetStatus;
  location?: string;
  warranty_expiry?: string;
  cost?: number;
  metadata?: Record<string, unknown>;
}

export interface PaginationParams {
  page?: number;
  limit?: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  pagination: {
    page: number;
    limit: number;
    total: number;
  };
}

export interface HealthResponse {
  status: string;
  version: string;
  service: string;
  mode: string;
}

export interface ConfigResponse {
  server: {
    port: string;
    mode: string;
  };
  database: {
    host: string;
    port: number;
    name: string;
  };
  blockchain: {
    chain_id: number;
    rpc_url: string;
  };
}

export interface BlockchainTx {
  tx_hash: string;
  event: string;
  timestamp: string;
  data: Record<string, unknown>;
}

export interface AssetHistoryResponse {
  asset_id: string;
  transactions: BlockchainTx[];
}
