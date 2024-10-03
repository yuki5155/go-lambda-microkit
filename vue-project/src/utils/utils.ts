export const setAuthToken = (token: string): void => {
    localStorage.setItem('authToken', token);
  };
  
  export const getAuthToken = (): string | null => {
    return localStorage.getItem('authToken');
  };
  
  export const removeAuthToken = (): void => {
    localStorage.removeItem('authToken');
  };
  
  export const setUserData = (user: object): void => {
    localStorage.setItem('userData', JSON.stringify(user));
  };
  
  export const getUserData = (): object | null => {
    const userData = localStorage.getItem('userData');
    return userData ? JSON.parse(userData) : null;
  };
  
  export const removeUserData = (): void => {
    localStorage.removeItem('userData');
  };
  
  // 必要に応じて他のユーティリティ関数を追加