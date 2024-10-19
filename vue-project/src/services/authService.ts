import apiClient, { handleApiError } from '@/utils/apiClient';

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface User {
  id: number;
  name: string;
  email: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export const authService = {
  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    try {
      console.log('Sending login request...');
      const response = await apiClient.post<LoginResponse>('/login', credentials);
      
      console.log('Response received:', response);
      console.log('Response data:', response.data);
      
      if (!response.data) {
        throw new Error('No data received from server');
      }
      
      if (!response.data.token || !response.data.user) {
        throw new Error('Invalid response format');
      }
      
      return response.data;
    } catch (error) {
      console.error('Login error:', error);
      throw handleApiError(error);
    }
  },

  async logout(): Promise<void> {
    try {
      await apiClient.post('/logout');
    } catch (error) {
      throw handleApiError(error);
    }
  },
};

export default authService;