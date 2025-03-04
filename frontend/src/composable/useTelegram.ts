import { ref, onMounted, onUnmounted } from 'vue';
import type { WebAppUser } from '../types/telegram';
import type { UserMeta } from '../types/user'

export function useTelegram() {
  const isDevelopment = import.meta.env.DEV
  const webApp = window.Telegram?.WebApp || null
  const user = ref<WebAppUser | null>(null);
  const userMeta = ref<UserMeta | null>(null);

  const getTelegramUser = (): WebAppUser => {
    if (isDevelopment) {
      const devUser = {
        id: 12345,
        first_name: 'Dev',
        last_name: 'User',
        username: 'devuser',
        is_bot: false
      };
      user.value = devUser;
      return devUser;
    }

    if (!webApp?.initDataUnsafe?.user) {
      throw new Error('Telegram WebApp user data not available');
    }

    user.value = webApp.initDataUnsafe.user;
    return user.value;
  }

  const collectUserMeta = (userId: number): UserMeta => {

    let first_seen_at = localStorage.getItem('first_seen_at');
    if (!first_seen_at) {
      first_seen_at = new Date().toISOString();
      localStorage.setItem('first_seen_at', first_seen_at);
    }

    const launch_count = Number(localStorage.getItem('launch_count') || 0) + 1;
    localStorage.setItem('launch_count', launch_count.toString());

    const session_start = Number(localStorage.getItem('session_start') || Date.now());
    localStorage.setItem('session_start', session_start.toString());

    return {
      user_id: userId,
      language_code: window.Telegram?.WebApp?.initDataUnsafe?.user?.language_code ?? '',
      is_premium: window.Telegram?.WebApp?.initDataUnsafe?.user?.is_premium ?? false,
      is_expanded: window.Telegram?.WebApp?.isExpanded ?? false,
      viewport_height: window.Telegram?.WebApp?.viewportHeight ?? 0,
      viewport_stable_height: window.Telegram?.WebApp?.viewportStableHeight ?? 0,
      platform: window.Telegram?.WebApp?.platform ?? '',
      last_active_at: new Date().toISOString(),
      first_seen_at,
      launch_count,
      session_duration: Math.floor((Date.now() - session_start) / 1000),
      device_resolution: `${window.screen.width}x${window.screen.height}`,
      device_pixel_ratio: window.devicePixelRatio,
      browser_info: navigator.userAgent,
      is_active: true,
      is_background: document.hidden
    } as UserMeta;
  };

  const updateUserMeta = () => {
    if (user.value) {
      userMeta.value = collectUserMeta(user.value.id);
    }
  };

  const initTelegram = () => {
    if (!isDevelopment && webApp) {
      webApp.ready()
      webApp.expand()
    }

    if (webApp?.initDataUnsafe?.user) {
      user.value = webApp.initDataUnsafe.user;
      updateUserMeta();
    }
  };

  onMounted(() => {
    initTelegram();
    document.addEventListener('visibilitychange', updateUserMeta);
  });

  onUnmounted(() => {
    document.removeEventListener('visibilitychange', updateUserMeta);
  });

  return {
    user,
    userMeta,
    getTelegramUser,
  };
}
