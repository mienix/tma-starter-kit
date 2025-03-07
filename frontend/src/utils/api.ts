export const api = {
  async get<T>(url: string, params?: Record<string, string | number>): Promise<T> {
    try {
      const query = params ? `?${new URLSearchParams(Object.entries(params).map(([key, value]) => [key, String(value)])).toString()}` : '';
      const response = await fetch(`${url}${query}`, {
        method: 'GET',
        headers: getHeaders(),
      });

      return handleResponse<T>(response);
    } catch (error) {
      console.error(`API GET ${url} error:`, error);
      throw error;
    }
  },

  async post<T>(url: string, body: unknown): Promise<T> {
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify(body),
      });

      return handleResponse<T>(response);
    } catch (error) {
      console.error(`API POST ${url} error:`, error);
      throw error;
    }
  }
};

function getHeaders(): HeadersInit {
  const headers: { [key: string]: string } = {
    'Content-Type': 'application/json',
  };
  const initData: string = window.Telegram?.WebApp?.initData || '';
  if (initData) {
    headers['X-Telegram-Init-Data'] = initData;
  }
  return headers;
}

async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const errorData = await response.json().catch(() => null);
    throw new Error(errorData?.message || `HTTP Error ${response.status}`);
  }

  const contentType = response.headers.get("content-type");
  if (!contentType || !contentType.includes("application/json")) {
    return null as T;
  }

  return response.json();
}
