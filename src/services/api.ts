import Taro from '@tarojs/taro';

const BASE_URL = 'http://localhost:8080'; // 根据实际后端地址修改

export interface WechatUserInfo {
  nickName?: string;
  avatarUrl?: string;
}

export interface LoginResponse {
  data: {
    token: string;
    expire_at: string;
  };
  err_code: number;
  message: string;
  request_id: string;
}

class Api {
  async login(code: string, avatarUrl?: string, nickName?: string): Promise<LoginResponse> {
    const response = await Taro.request<LoginResponse>({
      url: `${BASE_URL}/wechat/login`,
      method: 'POST',
      data: { code, avatarUrl, nickName }
    });

    if (response.statusCode !== 200) {
      throw new Error(response.data?.toString() || '登录失败');
    }

    return response.data;
  }

  async updateUserInfo(openId: string, userInfo: WechatUserInfo) {
    const response = await Taro.request({
      url: `${BASE_URL}/user/wechat/userinfo`,
      method: 'POST',
      data: {
        open_id: openId,
        nickname: userInfo.nickName,
        avatar: userInfo.avatarUrl
      }
    });

    if (response.statusCode !== 200) {
      throw new Error(response.data?.toString() || '更新用户信息失败');
    }

    return response.data;
  }
}

export const api = new Api();
