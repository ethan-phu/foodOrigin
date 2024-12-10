# taro3-react-mobx-tailwind-template

基于 Taro3 的 React 多端项目模版，默认支持编译成微信小程序，其他小程序自行测试

## 技术栈

- Taro
- React Hooks
- Mobx
- NutUI

## 运行环境

- Node >=16.14.2
- Pnpm 8+

## 支持特性

- 🏠 基于 Taro3.6.23
- 📦 支持 React
- 🐑 CSS 预处理器( SCSS )
- 🎨 TailwindCSS - 世界上最流行，生态最好的原子化 CSS 框架
- 🥣 完全使用 TypeScript 开发
- 🔛 request 拦截器封装
- 🔥 单元测试(jest+enzyme)
- 🧬 iconfont 支持(替换为自己的 iconfont 文件)
- 🌩️ 使用多核心及缓存提升编译速度
- 💰 更多特性持续迭代中...

## 安装和运行

### 环境准备
- Node.js >=16.14.2
- Pnpm 8+
- 微信开发者工具（用于运行微信小程序）

### 安装依赖
```bash
# 使用pnpm安装依赖
pnpm install
```

### 运行项目

```bash
# 运行微信小程序（开发环境）
pnpm dev:weapp

# 运行H5版本（开发环境）
pnpm dev:h5

# 运行其他平台
# 支持 weapp/swan/alipay/tt/h5/rn/qq/jd/quickapp
pnpm dev:<platform>
```

### 项目打包

```bash
# 打包微信小程序
pnpm build:weapp

# 打包H5版本
pnpm build:h5

# 打包其他平台
# 支持 weapp/swan/alipay/tt/h5/rn/qq/jd/quickapp
pnpm build:<platform>
```

### 运行测试
```bash
pnpm test
```

## 项目结构
```
├── config                 // 项目编译配置目录
├── src                    // 源码目录
├── test                   // 测试文件目录
├── .editorconfig         // 编辑器配置
├── .eslintrc.js         // ESLint配置
├── babel.config.js      // Babel配置
├── tsconfig.json        // TypeScript配置
├── project.config.json  // 小程序项目配置
└── package.json         // 项目依赖配置
```

## 注意事项
- 首次运行微信小程序时，需要使用微信开发者工具打开项目的dist目录
- 开发时请确保已经安装了对应平台的开发工具
- 如需修改项目配置，可以编辑config目录下的相关文件
