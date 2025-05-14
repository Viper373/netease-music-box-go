<!-- netease-music-box-go 说明：-->

## 🖥 Go版本使用说明

1. 在你的 Profile 仓库（如 Viper373/Viper373）配置 secrets（仓库 → Settings → Secrets and variables → Actions → New repository secret）：
   - `USER_ID`：网易云用户ID
   - `USER_TOKEN`：网易云用户Token（即 MUSIC_U）
   - `GH_TOKEN`：GitHub Token（需 gist 权限）
2. 在 `.github/workflows/` 目录下添加 `schedule.yml`（可参考本项目示例，注意要传递上述环境变量）
3. 在 `README.md` 需要插入内容的位置添加如下注释块：

```
<!-- netease-music-box start -->
<!-- netease-music-box end -->
```

4. 每天自动更新你的听歌排行到 README 或 Gist，无需手动操作

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
3. 编辑 `.github/workflows/schedule.yml` 文件，确保如下环境变量已传递：
    - `USER_ID`：网易云音乐用户 ID
    - `USER_TOKEN`：网易云音乐用户 Token
    - `GH_TOKEN`：GitHub Token
    - 其它可选变量：`GIST_FILENAME`、`UPDATE_OPTION`、`MARKDOWN_FILE`、`MARKDOWN_START_TAG`、`MARKDOWN_END_TAG`
4. 在项目的 `Settings > Secrets and variables > Actions` 中创建上述变量
5. 在 README 或其他 markdown 文件需要写入的地方添加如下注释块：

```
<!-- netease-music-box start -->
<!-- netease-music-box end -->
```

## 🤔 工作原理

- 基于 [NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) 获取听歌记录
- 基于 Github API 更新 Gist 或 Markdown
- 使用 Github Actions 自动更新

## ⚠️ 注意事项

- 若 workflow 运行失败，请检查环境变量是否配置完整，建议在 `main.go` 中增加环境变量校验，缺失时给出友好提示。
- 注释块必须完整且唯一，否则内容无法正确替换。

## 📄 开源协议

本项目使用 [MIT](./LICENSE) 协议
