import { View, Text, Image, Input } from '@tarojs/components'
import PageContainer from '@shared/components/PageContainer'
import { Button } from '@nutui/nutui-react-taro'
import Taro from '@tarojs/taro'
import { useState } from 'react'
import './index.scss'

import logoImg from '@shared/assets/images/logo.png'
import cameraIcon from '../../assets/images/camera.png'
import scanIcon from '../../assets/images/scan.png'

export default function Index() {
  const [searchValue, setSearchValue] = useState('')

  const handleScan = () => {
    Taro.scanCode({
      success: (res) => {
        setSearchValue(res.result)
        Taro.showToast({
          title: '扫描成功',
          icon: 'success',
          duration: 2000
        })
      },
      fail: () => {
        Taro.showToast({
          title: '扫描失败',
          icon: 'error',
          duration: 2000
        })
      }
    })
  }

  const handleCamera = () => {
    Taro.chooseImage({
      count: 1,
      sizeType: ['compressed'],
      sourceType: ['camera'],
      success: () => {
        Taro.showToast({
          title: '拍照成功',
          icon: 'success',
          duration: 2000
        })
      },
      fail: () => {
        Taro.showToast({
          title: '拍照失败',
          icon: 'error',
          duration: 2000
        })
      }
    })
  }

  return (
    <PageContainer>
      <View className='px-4 py-6'>
        <Image src={logoImg} className='w-32 h-8 mb-6' />
        
        <View className='search-section mb-6'>
          <View className='search-box flex items-center bg-white rounded-2xl shadow-sm'>
            <View className='search-input-wrapper flex-1 flex items-center'>
              <View className='search-icon ml-4'>
                <View className='circle' />
                <View className='handle' />
              </View>
              <Input
                className='flex-1 px-3 py-3 text-base'
                placeholder='搜索食品、化妆品'
                placeholderClass='text-gray-400'
                value={searchValue}
                onInput={(e) => setSearchValue(e.detail.value)}
              />
            </View>
            <View className='search-actions flex items-center px-2 space-x-2'>
              <Image
                src={scanIcon}
                className='search-action-icon w-6 h-6 p-1'
                onClick={handleScan}
              />
              <Image
                src={cameraIcon}
                className='search-action-icon w-6 h-6 p-1'
                onClick={handleCamera}
              />
            </View>
          </View>
        </View>

        <View className='grid grid-cols-2 gap-4 mb-6'>
          <View className='feature-item bg-white rounded-lg p-4 flex items-center' onClick={handleCamera}>
            <View className='icon-bg rounded-full mr-3'>
              <Image src={cameraIcon} className='w-6 h-6' />
            </View>
            <View>
              <Text className='font-medium'>拍照识别</Text>
              <Text className='text-sm block'>拍照识别成分</Text>
            </View>
          </View>

          <View className='feature-item bg-white rounded-lg p-4 flex items-center' onClick={handleScan}>
            <View className='icon-bg rounded-full mr-3'>
              <Image src={scanIcon} className='w-6 h-6' />
            </View>
            <View>
              <Text className='font-medium'>扫码查询</Text>
              <Text className='text-sm block'>扫码查询成分</Text>
            </View>
          </View>
        </View>
      </View>
    </PageContainer>
  )
}
