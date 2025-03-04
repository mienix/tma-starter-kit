export interface User {
  id: number;
  first_name: string;
  last_name?: string;
  username?: string;
}

export interface UserMeta {
  user_id: number;
  is_bot?: boolean;
  language_code?: string;
  is_premium?: boolean;
  is_active?: boolean;
  is_background?: boolean;
  is_expanded?: boolean;
  viewport_height?: number;
  viewport_stable_height?: number;
  platform?: string;
  last_active_at: string;
  first_seen_at: string;
  session_duration: number;
  launch_count: number;
  device_resolution?: string;
  device_pixel_ratio?: number;
  browser_info?: string;
}
