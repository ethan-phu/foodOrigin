import React from "react";
import { View } from "@tarojs/components";
import { NavBar } from "@nutui/nutui-react-taro";
import Taro from "@tarojs/taro";
import './index.scss';

interface IProps {
  className?: string;
  children: React.ReactNode;
  title?: string;
  showBack?: boolean;
}

const PageContainer = (props: IProps): JSX.Element => {
  const { className = "", children, title, showBack = false } = props;

  const handleBack = () => {
    Taro.navigateBack();
  };

  return (
    <View className={`page-container ${className}`}>
      {(title || showBack) && (
        <NavBar
          title={title}
          leftShow={showBack}
          onClickBack={handleBack}
          fixed
        />
      )}
      <View className={`page-content ${showBack ? 'pt-12' : ''}`}>
        {children}
      </View>
    </View>
  );
};

export default PageContainer;
