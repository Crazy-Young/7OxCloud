<h1 align="center" style="border-bottom: none;">7OxCloud</h1>
<h2 align="center" style="border-bottom: none;">基于推荐系统的高性能分布式Web短视频应用</h2>
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

#### 👀仓库地址：<https://github.com/Crazy-Young/7OxCloud>

#### 📚接口文档：<https://apifox.com/apidoc/shared-ddccc651-caf3-4d78-b296-eef873d9a6cd>

#### 🥽视频地址：<https://www.bilibili.com/video/BV1Cj411Y7TD>

#### 🏅架构设计：<https://github.com/Crazy-Young/7OxCloud/blob/master/docs/Architecture%20Design.docx>

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

## 推荐系统

推荐系统部分基于2022年WWW（CCF A)顶会的Paper [**"Filter-enhanced MLP is All You Need for Sequential Recommendation"**]()。本项目推荐系统代码部分为仓库主目录下的FMLP-RecSys文件夹

### Requirements

请提前下载安装python环境，具体见[Download Python | Python.org](https://www.python.org/downloads/)

项目需要的依赖及库已经形成于FMLP-RecSys/requirements.txt文件中，运行以下代码即可：

```shell
pip install -r FMLP-RecSys/requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
```

* Install Python, Pytorch(>=1.8).
* 如果你准备使用 GPU 计算,请安装 CUDA，并结合CUDA版本安装对应适配的Pytorch及相关库.具体可参考Pytorch官网：[Start Locally | PyTorch](https://pytorch.org/get-started/locally/#linux-pip)

### Usage

推荐系统部分被设计成一个服务端，用来与后端系统进行信息交互。因此在确保上述依赖和环境安装好的基础上，直接运行下述代码即可：

```shell
#如果想将该系统挂为后台应用，请参考使用linux screen命令
screen -S fmlp
screen -r fmlp

#注意运行server.py的主目录需要为FMLP-RecSys，否则可能会出现一些相对路径的问题
#切换到算法部分代码主目录
cd FMLP-RecSys

#将本机运行为服务端，请结合后端代码修改对应的监听端口
python server.py  
```

项目中main.py文件用于使用最新接受的数据来训练模型并生成模型文件，test.py文件用于根据后端发送过来的用户ID调用之前最新的模型以计算该用户最匹配的20个视频ID（为了方式与历史交互记录中的视频重合以及出现视频ID没有在任何人的交互记录中见过的情况，项目推荐结果生成之后会做一些筛选，因此总视频数可能会少于20）

后端系统会定期发送最新的交互记录记录（video-log.csv）并请求模型训练，代码将使用FMLP模型进行训练并保存本地模型文件；后端如若发送用户ID，代码将自动调用test.py文件的测试模型的接口进行该用户的top20匹配视频的推荐。

* 如果本地没有最新的模型文件，推荐将变为随机推荐。
* 用于训练的交互记录里（主要针对第一次训练），总视频ID总类需大于20，否则将会无法训练，继续使用随机推荐。
* 模型会自动筛掉用于训练的交互记录中交互数量小于3条的用户，因为模型至少需要分割成train、valid、test集才可进行训练，但模型会保存用户交互记录至本地的data/video-log-sorted-mapped.json文件，该文件将持续更新。
* 模型如果收到一个新用户的ID，将随机使用一位已记录过的其他用户的历史交互记录进行初始推荐，后续模型会根据该用户之后的交互进行更个性化更准确的推荐。

### Parameter

注意：需保证main.py和test.py中的模型参数一致！！

下述例举一些比较重要、影响模型较大的超参，可以根据个人需求和硬件情况进行个性化的更改：

* **hidden_size**：hidden size of model
* **num_hidden_layers**：number of filter-enhanced blocks
* **max_seq_length**：用于训练使用最长历史交互序列长度
* **epochs**：训练轮数
* **batch_size**：训练的batch大小
* **hidden_size**： hidden size of model，模型使用的向量的大小。
