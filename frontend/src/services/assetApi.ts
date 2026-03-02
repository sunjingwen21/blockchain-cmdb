import { authApi } from './authApi';

export interface Asset {
  id: number;
  name: string;
  type: string;
  status: 'active' | 'inactive' | 'maintenance' | 'retired';
  owner_id: number;
  owner_name?: string;
  blockchain_address?: string;
  metadata?: Record<string, any>;
  created_at: string;
  updated_at: string;
}

export interface CreateAssetRequest {
  name: string;
  type: string;
  status?: string;
  owner_id?: number;
  blockchain_address?: string;
  metadata?: Record<string, any>;
}

export interface UpdateAssetRequest {
  name?: string;
  type?: string;
  status?: string;
  owner_id?: number;
  blockchain_address?: string;
  metadata?: Record<string, any>;
}

export interface AssetListResponse {
  assets: Asset[];
  total: number;
  page: number;
  page_size: number;
}

export interface AssetFilters {
  status?: string;
  type?: string;
  search?: string;
  page?: number;
  page_size?: number;
}

export const assetService = {
  getAssets: async (filters: AssetFilters = {}): Promise<AssetListResponse> => {
    const params = new URLSearchParams();
    if (filters.status) params.append('status', filters.status);
    if (filters.type) params.append('type', filters.type);
    if (filters.search) params.append('search', filters.search);
    if (filters.page) params.append('page', filters.page.toString());
    if (filters.page_size) params.append('page_size', filters.page_size.toString());

    const response = await authApi.get(`/assets?${params.toString()}`);
    return response.data;
  },

  getAsset: async (id: number): Promise<Asset> => {
    const response = await authApi.get(`/assets/${id}`);
    return response.data;
  },

  createAsset: async (data: CreateAssetRequest): Promise<Asset> => {
    const response = await authApi.post('/assets', data);
    return response.data;
  },

  updateAsset: async (id: number, data: UpdateAssetRequest): Promise<Asset> => {
    const response = await authApi.put(`/assets/${id}`, data);
    return response.data;
  },

  deleteAsset: async (id: number): Promise<void> => {
    await authApi.delete(`/assets/${id}`);
  },

  getAssetTypes: async (): Promise<string[]> => {
    const response = await authApi.get('/assets/types');
    return response.data;
  },

  getAssetStats: async (): Promise<Record<string, number>> => {
    const response = await authApi.get('/assets/stats');
    return response.data;
  },
};
