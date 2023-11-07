<h1 align="center" style="border-bottom: none;">7OxCloud</h1>
<div class="labels" align="center">
    <a href="https://github.com/Crazy-Young/7OxCloud/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/Crazy-Young/7OxCloud?style=flat-square" alt="license">
    </a>
    <a href="#">
      <img src="https://img.shields.io/github/stars/Crazy-Young/7OxCloud?style=flat-square" alt="GitHub stars">
    </a>
    <a href="https://github.com/Crazy-Young/7OxCloud/network">
      <img src="https://img.shields.io/github/forks/Crazy-Young/7OxCloud?style=flat-square" alt="GitHub forks">
    </a>
    <span>
      <img src="https://img.shields.io/github/languages/top/Crazy-Young/7OxCloud" alt="language">
    </span>
    <span>
      <img src="https://img.shields.io/github/last-commit/Crazy-Young/7OxCloud" alt="last commit">
    </span>
   <span>
      <img src="https://komarev.com/ghpvc/?username=7OxCloud&label=Views&color=0e75b6&style=flat" alt="访问量统计" />
    </span>
</div>

####  👀仓库地址：<https://github.com/Crazy-Young/7OxCloud>

#### 📚文档地址：<>

#### 🥽视频地址：<https://www.bilibili.com/video/BV1Cj411Y7TD>

#### 🎶在线演示：<http://152.136.156.247>

由于阿里短信的限制，这里我们提供演示账号：

```
手机号：17383836194
密码：test123456
```

由于演示服务器配置的限制，推荐系统相关代码并未部署到服务器上。

# 程序运行说明

## 后端

本项目建议在linux系统下运行。首先需要保证系统上有`golang`环境：

```shell
#克隆仓库
git clone https://github.com/Crazy-Young/7OxCloud.git

#安装依赖
go mod tidy
```

第三方依赖：

- MySQL
- Redis
- Consul
- Nacos
- RabbitMQ
- Docker
- Nginx
- 七牛云对象存储
- 阿里云短信服务
- ......

建议使用docker来安装和部署相应依赖，教程可见：

<https://blog.csdn.net/m0_63230155/article/details/134090405?spm=1001.2014.3001.5501>

后端混合使用`viper`和`nacos`进行配置管理。需要分别编写api网关和微服务部分的`config.yaml`，这个`config.yaml`用于加载`nacos`的配置,例如:

```yaml
host: '127.0.0.1'
port: 8848
namespace: '6a11f263-64a0-4ca5-8b69-1259cd14dc87'
user: 'nacos'
password: 'nacos'
dataId: '7OxCloud-Api.json'
group: 'dev'
```

然后编写网关部分的`nacos`配置：

```json
{
    "api": {
        "name": "7OxCloud-Api",
        "host": "10.252.79.1",
        "tags": ["7OxCloud","Api"]
    },
    "service": {
        "user": "UserService",
        "video": "VideoService",
        "interaction": "InteractionService",
        "social": "SocialService"
    },
    "jwt": {
        "signingKey": "7OxCloudByPalp1tate",
        "expiration": 60
    },
    "sms": {
        "accessKeyId": "",
        "accessKeySecret": "",
        "signName": "",
        "templateCode": ""
    },
    "redis": {
        "host": "127.0.0.1",
        "port": 6379,
        "database": 0,
        "expiration": 120
    },
    "consul": {
        "host": "127.0.0.1",
        "port": 8500
    },
    "qiniuyun": {
        "accessKey": "",
        "secretKey": "",
        "bucket": "",
        "domain": ""
    }
}
```

由于使用的是分布式微服务架构，并引入了分布式配置中心，配置较为复杂，上述的两个操作只演示了网关部分，微服务部分的操作类似，在此不做赘述。

安装好相应的依赖并配置好后，只需要简单的命令即可启动项目：

```shell
#进入项目根目录
cd 7OxCloud

#提高脚本执行权限
sudo chmod 777 run.sh

#执行脚本
./run.sh
```

注意到该脚本也会自动启动nginx，确保其已经部署完成。

通过上述的操作，用户微服务、视频微服务、社交微服务、互动微服务、定时任务、两个api网关（nginx做反向代理）均会启动。当然每个微服务，网关可以同时开启多个，来提高整个系统的并发量。

## 前端

运行本项目需安装`node`，无`node`环境无法运行开发环境。如需安装`node`，请在官网安装稳定版，官网链接：[Node.js (nodejs.org)](https://nodejs.org/en)。

运行项目`node`版本须在 v16.0+。

### 开发环境

若需要运行**开发环境**，请在项目根目录下打开命令行使用以下指令，运行成功后项目运行在本地8080端口：

```shell
# 切换到web目录
$ cd ./web
# 安装依赖 也可以使用 npm i
$ npm install
# 运行项目
$ npm run serve
```

如需配置本地代理地址为本地网址，只需修改`vue.config.js`中`proxy`中`/api`配置的`target`为本地后端服务地址即可。

如需配置生产环境地址，可修改`.env.production`文件中的`VUE_APP_API_BASE_URL`，此时打包文件请求接口基地址均为设定后端服务地址。

### 生产环境

**生产环境**请直接运行`dist`文件夹下的内容，也可以直接双击打开`index.html`文件（不推荐）。

请在项目根目录下使用以下指令（如使用指令时遇到选择，请选择 *y* ）：

```shell
# 切换到web目录
$ cd ./web/dist
# 若有node环境，可执行：
$ npx serve
# 若有python环境，可执行：
$ python -m http.server
```