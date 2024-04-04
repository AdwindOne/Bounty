# Bounty

Bounty 是一个基于 [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) 和 [Ignite CLI](https://ignite.com/cli) 开发的去中心化应用（dApp），允许用户发布、领取并在完成后获得赏金的任务管理平台。

## 快速开始

以下步骤将指导你如何在本地环境安装、运行并测试Bounty。

### 环境要求

确保你的开发环境中已安装：

- [Git](https://git-scm.com/)
- [Go](https://golang.org/doc/install) 版本 1.21 或更高
- [Ignite CLI](https://ignite.com/cli)

### 安装指南

克隆仓库并准备环境：

```bash
git clone https://github.com/AdwindOne/Bounty.git
cd bounty
```

构建项目并启动本地节点：
```bash
ignite chain build
ignite chain serve
```

### 基本操作
发布任务
发布一个新的赏金任务：
```bash
bountyd tx bounty create-bounty [title] [description] [reward] --from=[your_key_name]
```

领取任务
领取一个现有的赏金任务：
```bash
bountyd tx bounty claim-bounty [bountyId] --from=[hacker_key_name]
```

完成任务
一旦任务完成，更新任务状态：
```bash
bountyd tx bounty complete-bounty [claimId] --from=[your_key_name]
```


### 如何贡献
我们欢迎并鼓励社区成员以各种形式贡献，无论是提出建议、报告问题或直接贡献代码。请通过创建issue或拉取请求来提交贡献。

### 许可
该项目采用 MIT 许可证。更多信息请参见 LICENSE 文件。

### 致谢
感谢 Cosmos SDK 和 Ignite CLI 团队提供的工具和库。
感谢所有直接或间接贡献了代码、文档和思想的社区成员。