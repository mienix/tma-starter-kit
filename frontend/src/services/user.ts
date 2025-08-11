import type { User, UserMeta } from '../types/user';
import { api } from '../utils/api';

export const userService = {
  async saveUserWithMeta(user: User, meta: UserMeta): Promise<void> {
    try {
      await api.post('/api/users', { ...user, meta });
    } catch (error) {
      console.error('Ошибка при создании пользователя и метаинформации:', error);
      throw error;
    }
  },

  async getUser(userId: number): Promise<User | null> {
    try {
      const response = await api.get<User>(`/api/users/${userId}`);
      return response;
    } catch (error) {
      console.error('Ошибка при получении пользователя:', error);
      return null;
    }
  }
};
