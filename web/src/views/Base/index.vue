<template>
    <el-container>
        <!-- 顶部导航栏 -->
        <el-header height="68px">
            <!-- 左上角的logo -->
            <div class="logo" @click="navigateTo('/')">
                <Logo></Logo>
            </div>
            <!-- 顶部搜索区域 -->
            <el-input v-model="search" @keyup.enter.native="navigateTo(`/search?keyword=${search}`)">
                <template #append>
                    <button @click="navigateTo(`/search?keyword=${search}`)">
                        <SearchIcon></SearchIcon>
                        <span>搜索</span>
                    </button>
                </template>
            </el-input>
            <!-- 右侧用户登录区域 -->
            <div class="user">
                <div class="tools">
                    <div class="upload" @click="navigateTo('/upload')">
                        <PublishIcon></PublishIcon>
                        <span class="text">投稿</span>
                    </div>
                </div>
                <el-avatar alt="logo" @click.native="navigateTo('/me')" v-if="userInfo.avatar" :src="userInfo.avatar">{{
                    userInfo.username }}</el-avatar>
                <button type="button" v-else class="button" @click="handleLogin">
                    <UserIcon></UserIcon>
                    <p class="login">登录</p>
                </button>
            </div>
        </el-header>

        <!-- 主体部分 -->
        <el-container>
            <!-- 左侧导航栏 -->
            <el-aside>
                <!-- 配置路由 -->
                <el-menu :default-active="$route.path" router>
                    <el-menu-item :index="item.path" v-for="item in routes" :key="item.name">
                        <div class="icon"
                            :style="{ '--index': currentRoute === item.path ? item.meta.index * 2 + 1 : item.meta.index * 2 }">
                        </div>
                        <span slot="title">{{ item.meta.title }}</span>
                    </el-menu-item>
                </el-menu>
            </el-aside>

            <!-- 右侧内容区 -->
            <el-main>
                <router-view class="main" :key="$route.fullPath"></router-view>
            </el-main>
        </el-container>

    </el-container>
</template>

<script>
import Logo from "@/icons/Logo";
import { mapState } from "vuex";
import VideoBox from "@/components/VideoBox";
import SearchIcon from "@/icons/SearchIcon.vue";
import PublishIcon from "@/icons/PublishIcon.vue"
import UserIcon from "@/icons/UserIcon.vue"

// 映射类别对应的nav_bg中icon的位置
const indexMap = {
    "其他": 16,
    "生活": 11,
    "教育": 8,
    "体育": 12
}

export default {
    name: "Base",
    components: {
        Logo,
        VideoBox,
        SearchIcon,
        PublishIcon,
        UserIcon
    },
    data() {
        return {
            search: '',
        }
    },
    computed: {
        // 页面路由
        routes() {
            // 获取所有的路由，并添加四个类别的频道路由
            return [
                ...this.$router.getRoutes().filter((item) => item.meta.isNav),
                ...this.categoryList.reverse().map(_ => ({
                    path: '/channel/' + _.id,
                    name: _.name,
                    meta: {
                        title: _.name,
                        index: indexMap[_.name]
                    }
                }))
            ]
        },
        // 当前路由
        currentRoute() {
            return this.$route.path
        },
        ...mapState({
            userInfo: state => state.user.userInfo,
            categoryList: state => state.category.categoryList
        })
    },
    methods: {
        navigateTo(path) {
            this.$router.push({
                path
            })
        },
        handleLogin() {
            this.$bus.$emit("toggleLogin")
        }
    },
};
</script>

<style lang="less" scoped>
.el-container {

    .el-header {
        z-index: 99;
        line-height: 68px;
        padding: 0 32px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        white-space: nowrap;
        min-width: 540px;

        .el-input {
            width: 420px;
            margin: 0 20px 0 auto;
            border: unset;
            display: flex;
            align-items: center;
            height: 40px;

            /deep/ .el-input__inner {
                background-color: @gray-4;
                border: unset;
                color: @gray-1;
            }

            /deep/ .el-input-group__append {
                height: 100%;
                border: unset;
                width: fit-content;
                padding: 0;
                height: 100%;
                background-color: @gray-4;
                display: flex;

                button {
                    height: 100%;
                    all: unset;
                    line-height: 1;
                    padding: 0 20px 0 0;
                    display: flex;
                    font-size: 16px;
                    justify-content: center;
                    align-items: center;
                    gap: 5px;

                    &:hover {
                        color: @gray-1;

                        svg {
                            path {
                                fill: @gray-1;
                            }
                        }
                    }
                }
            }
        }

        .logo {
            cursor: pointer;
            display: flex;
            align-items: center;
            margin-right: 20px;
        }

        .user {
            display: flex;
            align-items: center;
            gap: 10px;
            cursor: pointer;

            .tools {
                display: flex;
                align-items: center;
                margin: 0 10px;

                &>div {
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                    line-height: 1;

                    &:hover {
                        svg {
                            fill: @gray-1;
                        }

                        .text {
                            color: @gray-1;
                        }
                    }

                    svg {
                        fill: @gray-2;
                    }

                    .text {
                        font-size: 12px;
                        line-height: 20px;
                    }
                }
            }

            .button {
                border-radius: 12px;
                height: 38px;
                padding: 6px 16px;
                width: 88px;
                box-sizing: border-box;
                align-items: center;
                justify-content: space-between;
                background-color: @red;
                font-weight: 500;
                border: unset;
                color: @white;
                display: flex;
                font-size: 15px;
                margin: 0;

                &:hover {
                    filter: brightness(0.9);
                }

                &:active {
                    filter: brightness(0.8);
                }

                .login {
                    color: @white;
                    font-size: 15px;
                    font-weight: 400;
                    margin-left: 4px;
                    font-family: "PingFang SC", "Helvetica Neue", Helvetica, "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
                }
            }
        }
    }



    .el-container {

        .el-aside {
            width: 160px !important;
            overflow: auto;

            .el-menu {
                height: 100%;
                background-color: transparent;
                border-right: unset;
                padding: 0 16px;
                display: flex;
                flex-direction: column;
                --index: 0;

                .el-menu-item {
                    width: 100%;
                    padding: 8px 16px !important;
                    border-radius: 12px;
                    display: flex;
                    line-height: unset;
                    height: unset;
                    align-items: center;
                    color: @font;
                    gap: 12px;

                    &:not(:last-child) {
                        margin-bottom: 3px;
                    }


                    &:focus,
                    &:active {
                        background-color: unset;
                    }

                    &.is-active,
                    &:hover {
                        background-color: @transparent-dark;
                        color: @white;

                        .icon {
                            opacity: 1;
                        }
                    }

                    .icon {
                        width: 24px;
                        height: 24px;
                        background: url('../../assets/nav.png') no-repeat;
                        background-size: auto 24px;
                        background-position-x: calc(var(--index) * -24px);
                        opacity: .5;
                        background-position-y: 0;
                    }
                }
            }
        }

        .el-main {
            height: calc(100vh - 68px);
            overflow-y: overlay;
            overflow-x: hidden;
            padding: 0;

            &::-webkit-scrollbar {
                width: 8px;
            }

            &::-webkit-scrollbar-thumb {
                background-color: @gray-3;
                border-radius: 4px;
            }
        }

        &>div {
            height: 100%;
            width: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
        }
    }
}
</style>
