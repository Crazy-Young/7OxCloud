<template>
    <el-container>
        <el-header height="68px">
            <div class="logo" @click="navigateTo('/')">
                <Logo></Logo>
            </div>
            <div class="user">
                <div class="tools">
                    <div class="upload" @click="navigateTo('/upload')">
                        <svg width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg" class="lG5TmBA_"
                            viewBox="0 0 24 24">
                            <path fill-rule="evenodd" clip-rule="evenodd"
                                d="M11.349 5.17h-.033c-1.068 0-1.915 0-2.598.054-.698.056-1.29.172-1.832.441a4.75 4.75 0 00-2.14 2.14c-.269.542-.386 1.134-.441 1.832-.055.683-.055 1.53-.055 2.599v.064c0 1.069 0 1.916.055 2.599.055.698.172 1.29.441 1.831a4.75 4.75 0 002.14 2.14c.542.27 1.134.386 1.832.442.683.055 1.53.055 2.598.055H12.684c1.068 0 1.915 0 2.598-.055.698-.056 1.29-.172 1.832-.442a4.75 4.75 0 002.14-2.14c.269-.541.386-1.133.441-1.831.055-.683.055-1.53.055-2.599v-.064c0-1.069 0-1.916-.055-2.599-.055-.698-.172-1.29-.441-1.832a4.75 4.75 0 00-2.14-2.14c-.542-.269-1.134-.385-1.832-.441-.683-.055-1.53-.055-2.598-.055h-1.335zM7.554 7.008c.299-.149.676-.241 1.284-.29.616-.05 1.403-.05 2.51-.05h1.303c1.108 0 1.895 0 2.511.05.608.049.985.141 1.284.29a3.25 3.25 0 011.464 1.464c.15.3.241.676.29 1.284.05.616.05 1.403.05 2.51 0 1.109 0 1.896-.05 2.512-.049.608-.14.985-.29 1.284a3.25 3.25 0 01-1.464 1.464c-.299.149-.676.241-1.284.29-.616.05-1.403.05-2.51.05h-1.303c-1.108 0-1.895 0-2.511-.05-.608-.049-.985-.141-1.284-.29a3.25 3.25 0 01-1.464-1.464c-.15-.3-.242-.676-.29-1.284-.05-.616-.05-1.403-.05-2.511s0-1.895.05-2.511c.048-.608.14-.985.29-1.284a3.25 3.25 0 011.464-1.464zm3.696 8.259v-2.25H9v-1.5h2.25v-2.25h1.5v2.25H15v1.5h-2.25v2.25h-1.5z">
                            </path>
                        </svg>
                        <span class="text">投稿</span>
                    </div>
                </div>
                <el-avatar alt="logo" @click.native="navigateTo('/me')" v-if="userInfo.avatar" :src="userInfo.avatar">{{
                    userInfo.username }}</el-avatar>
                <button type="button" v-else class="button" @click="toggleLogin"><svg width="18" height="18" fill="none"
                        xmlns="http://www.w3.org/2000/svg" class="" viewBox="0 0 18 18" id="svg_icon_avatar">
                        <rect width="18" height="18" rx="9" fill="#fff" fill-opacity="0.3"></rect>
                        <path
                            d="M2.244 14.946A8.978 8.978 0 019 11.893a8.978 8.978 0 016.756 3.053A8.978 8.978 0 019 18a8.978 8.978 0 01-6.756-3.054zM12.214 7.393a3.214 3.214 0 11-6.428 0 3.214 3.214 0 016.428 0z"
                            fill="#fff"></path>
                    </svg>
                    <p class="login">登录</p>
                </button>
            </div>
        </el-header>

        <el-container>
            <el-aside width="160px">
                <el-menu :default-active="$route.path" router>
                    <el-menu-item :index="item.path" v-for="item in routes" :key="item.name">
                        <div class="icon"
                            :style="{ '--index': currentRouteName === item.name ? item.meta.index * 2 + 1 : item.meta.index * 2 }">
                        </div>
                        <span slot="title">{{ item.meta.title }}</span>
                    </el-menu-item>
                </el-menu>
            </el-aside>

            <el-main>
                <router-view class="main" :key="$route.path"></router-view>
            </el-main>
        </el-container>

        <Login v-if="isLogin" @close="toggleLogin"></Login>
    </el-container>
</template>

<script>
import Logo from "@/components/Logo";
import Login from "@/components/Login";
import { mapState } from "vuex";
import { VideoCategoryList } from "@/api/video";
export default {
    name: "Base",
    components: {
        Logo,
        Login
    },
    data() {
        return {
            isLogin: false,
            categories: [],
        }
    },
    created() {
        VideoCategoryList().then(res => {
            if (res.status === 200) {
                const { categories } = res.data.data
                if (categories)
                    this.categories = categories
            }
        })
    },
    computed: {
        routes() {
            const indexMap = {
                "其他": 16,
                "生活": 11,
                "教育": 8,
                "体育": 12
            }

            return [
                ...this.$router.getRoutes().filter((item) => item.meta.isNav),
                ...this.categories.reverse().map(_ => ({
                    path: '/channel/' + _.id,
                    name: _.name,
                    meta: {
                        title: _.name,
                        index: indexMap[_.name]
                    }
                }))]
        },
        currentRouteName() {
            return this.$route.name
        },
        ...mapState({
            userInfo: state => state.user.userInfo
        })
    },
    methods: {
        toggleLogin() {
            this.isLogin = !this.isLogin
        },
        navigateTo(path) {
            this.$router.push({
                path
            })
        }
    }
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
        white-space: nowrap;

        .logo,
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
                            path {
                                fill: @gray-1;
                            }
                        }

                        .text {
                            color: @gray-1;
                        }
                    }

                    svg {
                        path {
                            fill: @gray-2;
                        }
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

    .el-aside {

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
                    background-color: @nav-bg-hover;
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

    .el-container {

        .el-main {
            height: calc(100vh - 68px);
            overflow-y: overlay;
            padding: 0;

            &::-webkit-scrollbar {
                display: none;
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
