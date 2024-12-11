import Taro from '@tarojs/taro'
import { API_BASE_URL } from '@/constants/config'

export interface LoginResponse {
  token: string
  expires_in: number
}

/**
 * 微信登录
 */
export const login = async (): Promise<LoginResponse> => {
  try {
    // 获取微信登录凭证
    const { code } = await Taro.login()
    
    // 调用后端登录接口
    const res = await Taro.request({
      url: `${API_BASE_URL}/api/wechat/login`,
      method: 'POST',
      data: { code }
    })

    // 保存 token
    if (res.statusCode === 200 && res.data.token) {
      Taro.setStorageSync('token', res.data.token)
      return res.data
    }

    throw new Error(res.data.error || '登录失败')
  } catch (error) {
    console.error('登录失败:', error)
    throw error
  }
}

/**
 * 检查登录状态
 */
export const checkLogin = (): boolean => {
  try {
    const token = Taro.getStorageSync('token')
    return !!token
  } catch (error) {
    return false
  }
}

/**
 * 退出登录
 */
export const logout = (): void => {
  try {
    Taro.removeStorageSync('token')
  } catch (error) {
    console.error('退出登录失败:', error)
  }
}
