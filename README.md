# 食知源

一个基于 Taro 3.x 开发的微信小程序，帮助用户快速查询食品和化妆品的成分信息。

## 项目简介

食知源是一款专注于食品和化妆品成分查询的小程序。通过扫描条形码或拍照识别的方式，用户可以快速了解产品的成分信息，帮助做出更明智的消费选择。

## 技术栈

- Taro 3.x
- React
- NutUI React
- TypeScript
- SCSS
- TailwindCSS

## 功能特性

- 🔍 搜索功能：支持文本搜索食品和化妆品信息
- 📸 拍照识别：通过拍照快速识别产品成分
- 📱 扫码查询：扫描商品条形码获取详细信息
- 💡 用户友好：简洁直观的界面设计
- 🎨 响应式设计：完美适配各种屏幕尺寸

## 快速开始

### 环境要求

- Node.js 14.0+
- 微信开发者工具

### 安装依赖

```bash
# 使用 npm
npm install

# 或使用 yarn
yarn
```

### 开发

```bash
# 开发模式
npm run dev:weapp

# 生产构建
npm run build:weapp
```

### 项目结构

```
src/
├── app.config.ts        # 应用配置
├── app.scss            # 全局样式
├── app.tsx             # 应用入口
├── index.html          # HTML 模板
├── pages/             # 页面文件夹
│   ├── home/          # 首页
│   └── my/            # 我的页面
└── shared/            # 共享资源
    ├── components/    # 公共组件
    └── assets/        # 静态资源
```

## 开发指南

### 新增页面

1. 在 `src/pages` 下创建新的页面文件夹
2. 在 `app.config.ts` 中添加页面路径
3. 实现页面组件和样式

### 样式开发

- 使用 SCSS 进行样式开发
- 遵循 BEM 命名规范
- 优先使用 TailwindCSS 类名
- 特殊样式使用 SCSS 编写

## 贡献指南

1. Fork 本仓库
2. 创建您的特性分支 (git checkout -b feature/AmazingFeature)
3. 提交您的更改 (git commit -m 'Add some AmazingFeature')
4. 推送到分支 (git push origin feature/AmazingFeature)
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详细信息

## 联系方式

如有任何问题或建议，欢迎提出 Issue 或 Pull Request。

---

🌟 如果这个项目对您有帮助，欢迎点个 star 支持一下！
