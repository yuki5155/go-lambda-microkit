import axios from 'axios';
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
        const response = await apiClient.get<LoginResponse>('/hello/', {
          timeout: 10000, // 10 seconds timeout
        });
        
        console.log('Response received:', response);
        console.log('Response headers:', response.headers);
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
        if (axios.isAxiosError(error)) {
          if (error.response) {
            console.error('Error data:', error.response.data);
            console.error('Error status:', error.response.status);
            console.error('Error headers:', error.response.headers);
          } else if (error.request) {
            console.error('Error request:', error.request);
          } else {
            console.error('Error message:', error.message);
          }
          console.error('Error config:', error.config);
        }
        throw handleApiError(error);
      }
    },

  async logout(): Promise<void> {
    try {
      await apiClient.post('/auth/logout');
    } catch (error) {
      throw handleApiError(error);
    }
  },
};

export default authService;