import { ref } from 'vue'
import { userService } from '../services/user'
import type { User, UserMeta } from '../types/user'
import { useTelegram } from './useTelegram'

export function useUser() {
  const { getTelegramUser, userMeta } = useTelegram()
  const currentUser = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const initializeUser = async (): Promise<void> => {
    loading.value = true
    error.value = null

    try {
      const telegramUser = getTelegramUser()
      currentUser.value = telegramUser

      await userService.saveUserWithMeta({
        id: telegramUser.id,
        first_name: telegramUser.first_name,
        last_name: telegramUser.last_name,
        username: telegramUser.username
      } as User,
        userMeta.value as UserMeta)
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Failed to initialize user'
      console.error('Failed to initialize user:', e)
    } finally {
      loading.value = false
    }
  }

  return {
    currentUser,
    loading,
    error,
    initializeUser
  }
}
