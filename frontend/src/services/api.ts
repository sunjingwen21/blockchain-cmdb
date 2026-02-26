import axios, { AxiosInstance, AxiosResponse } from 'axios';
import {
  User,
  UserCreateRequest,
  UserUpdateRequest,
  Asset,
  AssetCreateRequest,
  AssetUpdateRequest,
  PaginatedResponse,
  PaginationParams,
  HealthResponse,
  ConfigResponse,
  AssetHistoryResponse,
} from '../types';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

class ApiService {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: `${API_BASE_URL}/api/v1`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Add request interceptor for auth token
    this.client.interceptors.request.use((config) => {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });

    // Add response interceptor for error handling
    this.client.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          localStorage.removeItem('token');
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }

  // Health check
  async healthCheck(): Promise<HealthResponse> {
    const response: AxiosResponse<HealthResponse> = await axios.get(
      `${API_BASE_URL}/health`
    );
    return response.data;
  }

  // Config
  async getConfig(): Promise<ConfigResponse> {
    const response: AxiosResponse<ConfigResponse> = await this.client.get('/config');
    return response.data;
  }

  // Users
  async getUsers(params?: PaginationParams): Promise<PaginatedResponse<User>> {
    const response: AxiosResponse<PaginatedResponse<User>> = await this.client.get(
      '/users',
      { params }
    );
    return response.data;
  }

  async getUser(id: number): Promise<User> {
    const response: AxiosResponse<User> = await this.client.get(`/users/${id}`);
    return response.data;
  }

  async createUser(data: UserCreateRequest): Promise<User> {
    const response: AxiosResponse<User> = await this.client.post('/users', data);
    return response.data;
  }

  async updateUser(id: number, data: UserUpdateRequest): Promise<User> {
    const response: AxiosResponse<User> = await this.client.put(`/users/${id}`, data);
    return response.data;
  }

  async deleteUser(id: number): Promise<void> {
    await this.client.delete(`/users/${id}`);
  }

  // Assets
  async getAssets(
    params?: PaginationParams & { type?: string; status?: string }
  ): Promise<PaginatedResponse<Asset>> {
    const response: AxiosResponse<PaginatedResponse<Asset>> = await this.client.get(
      '/assets',
      { params }
    );
    return response.data;
  }

  async getAsset(id: number): Promise<Asset> {
    const response: AxiosResponse<Asset> = await this.client.get(`/assets/${id}`);
    return response.data;
  }

  async createAsset(data: AssetCreateRequest): Promise<Asset> {
    const response: AxiosResponse<Asset> = await this.client.post('/assets', data);
    return response.data;
  }

  async updateAsset(id: number, data: AssetUpdateRequest): Promise<Asset> {
    const response: AxiosResponse<Asset> = await this.client.put(`/assets/${id}`, data);
    return response.data;
  }

  async deleteAsset(id: number): Promise<void> {
    await this.client.delete(`/assets/${id}`);
  }

  async getAssetHistory(id: number): Promise<AssetHistoryResponse> {
    const response: AxiosResponse<AssetHistoryResponse> = await this.client.get(
      `/assets/${id}/history`
    );
    return response.data;
  }
}

export const apiService = new ApiService();
export default apiService;
