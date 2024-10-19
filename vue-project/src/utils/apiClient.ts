import axios from 'axios';
import type { AxiosInstance, AxiosError } from 'axios';


const apiClient: AxiosInstance = axios.create({
  baseURL: process.env.NODE_ENV === 'production' 
    ? process.env.VUE_APP_API_URL 
    : '/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const handleApiError = (error: unknown): never => {
  if (axios.isAxiosError(error)) {
    const axiosError = error as AxiosError<{ message: string }>;
    if (axiosError.response) {
      throw new Error(axiosError.response.data.message || 'An error occurred during the API request');
    }
  }
  throw new Error('An unexpected error occurred');
};

export default apiClient;