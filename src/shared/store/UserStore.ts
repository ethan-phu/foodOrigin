import { makeAutoObservable } from 'mobx';
import Taro from '@tarojs/taro';
import { api, LoginResponse, WechatUserInfo } from '@/services/api';
import { ApiConfig } from '@/config';

interface UserInfoResponse {
  avatar: string;
  created: string;
  email: string;
  email_valid: boolean;
  id: number;
  mobile: string;
  mobile_valid: boolean;
  modified: string;
  name: string;
  nickname: string;
  online: boolean;
  open_id: string;
  sex: number;
  union_id: string;
}

class UserStore {
  userInfo: UserInfoResponse | null = null;
  isLoggedIn: boolean = false;
  token: string = '';

  constructor() {
    makeAutoObservable(this);
    // 从本地存储恢复 token
    this.token = Taro.getStorageSync('token') || '';
    this.isLoggedIn = !!this.token;
    // 从本地存储恢复用户信息
    const storedUserInfo = Taro.getStorageSync('userInfo');
    if (storedUserInfo) {
      this.setUserInfo(storedUserInfo);
    }
  }

  setToken(token: string) {
    this.token = token;
    // 保存 token 到本地存储
    Taro.setStorageSync('token', token);
  }

  setUserInfo(userInfo: UserInfoResponse) {
    this.userInfo = userInfo;
    this.isLoggedIn = true; // 设置用户信息时确保登录状态为 true
    // 保存用户信息到本地存储
    Taro.setStorageSync('userInfo', userInfo);
  }

  updateUserInfo(userInfo: WechatUserInfo) {
    if (!this.token) {
      throw new Error('Not logged in');
    }

    // 直接更新本地用户信息
    const updatedUserInfo = {
      ...this.userInfo,
      nickname: userInfo.nickName,
      avatar: userInfo.avatarUrl,
    } as UserInfoResponse;

    this.setUserInfo(updatedUserInfo);
  }

  async getUserProfile(): Promise<WechatUserInfo> {
    try {
      const { userInfo } = await Taro.getUserProfile({
        desc: '用于完善会员资料',
        lang: 'zh_CN'
      });

      // 更新用户信息到本地存储
      this.updateUserInfo(userInfo);
      return userInfo;
    } catch (error) {
      console.error('Get user profile failed:', error);
      throw error;
    }
  }

  async login() {
    try {
      const { code } = await Taro.login();
      const response = await api.login(code);
      
      if (response.err_code === 0) {
        const { token } = response.data;
        this.setToken(token);
        // 等待获取用户信息完成
        const userInfo = await this.fetchUserInfo();
        // 确保设置登录状态
        this.setUserInfo(userInfo);
      } else {
        throw new Error(response.message || '登录失败');
      }
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  }

  async fetchUserInfo() {
    if (!this.token) {
      throw new Error('Not logged in');
    }

    try {
      const response = await Taro.request({
        url: `${ApiConfig.base_url}/v1/user`,
        method: 'GET',
        header: {
          'Authorization': `Bearer ${this.token}`
        }
      });

      if (response.statusCode === 200 && response.data.err_code === 0) {
        this.setUserInfo(response.data.data);
        return response.data.data;
      } else {
        throw new Error(response.data.message || '获取用户信息失败');
      }
    } catch (error) {
      console.error('Fetch user info failed:', error);
      throw error;
    }
  }

  logout() {
    this.token = '';
    this.userInfo = null;
    this.isLoggedIn = false;
    // 清除本地存储
    Taro.removeStorageSync('token');
    Taro.removeStorageSync('userInfo');
  }
}

export const userStore = new UserStore();
