<!-- netease-music-box-go 说明：-->

## 🖥 Go版本使用说明

1. 在你的 Profile 仓库（如 Viper373/Viper373）配置 secrets：
   - USER_ID: 网易云用户ID
   - USER_TOKEN: 网易云用户Token
2. 在 .github/workflows/ 目录下添加 schedule.yml（见本项目示例）
3. 在 README.md 需要插入内容的位置添加：
   <!-- netease-music-box start -->
   <!-- netease-music-box end -->
4. 每天自动更新你的听歌排行到 README

---

<table align="center">
<tr>
<td>

<!-- netease-music-box start -->
#### <a href="https://gist.github.com/0de3f9fc7f3078a800f738e25eccea54" target="_blank">🎵 我最近一周的听歌排行</a>
```text
🥇 这，就是爱 - 张杰			5次    
🥈 My Type - T...			4次    
🥉 Off The Hoo...			4次    
🏅 平凡的一天 - 毛不易			3次    
🏅 千秋令 - 银临/KB...			3次    
```

<!-- netease-music-box end -->

</td>
</tr>
</table>

<p align="center">
  <h2 align="center">Netease Music Box</h2>
  <p align="center">将你最近一周的网易云音乐的听歌记录更新到 Gist</p>
</p>

---

> 📌✨ 更多像这样的 Pinned Gist 项目请访问：https://github.com/matchai/awesome-pinned-gists

## 🖥 使用

### 🎒 前置工作

1. 创建一个公开的 Github Gist (https://gist.github.com)

2. 创建一个 GitHub Token，需要勾选 `gist` 权限，复制生成的 Token (https://github.com/settings/tokens/new)

3. 获取网易云音乐用户 ID (https://music.163.com)

    - ID 为个人主页页面（`https://music.163.com/#/user/home?id=xxx`），`id` 后紧跟的那串数字

    ![USER_ID](https://github.com/llnancy/netease-music-box/blob/master/assets/user_id.png)

4. 获取网易云音乐用户 Token

    - 在登录态下打开浏览器开发者工具，查看 Cookie，获取 `key` 为 `MUSIC_U` 的 `value`

    ![USER_TOKEN](https://github.com/llnancy/netease-music-box/blob/master/assets/user_token.png)

### 🚀 安装

1. Fork 这个仓库

2. 进入 Fork 后的仓库，启用 Github Actions

3. 编辑 `.github/workflows/schedule.yml` 文件中的环境变量：

    - **GIST_ID**: ID 是新建 Gist 的 `url` 后缀: `https://gist.github.com/llnancy/`**`475826c54f1a5cd2037aa96c604043c0`**

    - **USER_ID**: 网易云音乐用户 ID

4. 在项目的 `Settings > Secrets` 中创建两个变量 `GH_TOKEN` 和 `USER_TOKEN`，分别为 Github Token 和 网易云音乐用户 Token

5. [在个人资料中嵌入 Gist](https://docs.github.com/en/github/setting-up-and-managing-your-github-profile/pinning-items-to-your-profile)

6. 如果需要写入到某个 `markdown` 文件，请在对应文件需要写入的地方添加以下注释

```text
<!-- netease-music-box start -->
<!-- netease-music-box end -->
```

## 🤔 工作原理

- 基于 [NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) 获取听歌记录
- 基于 Github API 更新 Gist
- 使用 Github Actions 自动更新 Gist

## 📄 开源协议

本项目使用 [MIT](./LICENSE) 协议
