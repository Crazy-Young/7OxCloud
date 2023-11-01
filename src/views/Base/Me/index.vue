<template>
    <el-container>
        <header>
            <el-avatar :size="120" :src="userInfo.avatar">{{ userInfo.username || '未登录' }}</el-avatar>
            <div class="userInfo">
                <h1 class="username">{{ userInfo.username || '未登录' }}</h1>
                <p class="list">
                    <span>关注</span><i>{{ userInfo.uid ? userInfo.followCount : '-' }}</i>
                    <el-divider direction="vertical"></el-divider>
                    <span>粉丝</span><i>{{ userInfo.uid ? userInfo.fanCount : '-' }}</i>
                    <el-divider direction="vertical"></el-divider>
                    <span>获赞</span><i>{{ userInfo.uid ? userInfo.likedCount : '-' }}</i>
                </p>
                <p class="other" v-if="userInfo.uid">
                    <span class="account">账号ID: {{ userInfo.uid }}</span>
                    <span class="gender tag">
                        <svg width="12" v-if="userInfo.gender === 1" height="12" fill="none"
                            xmlns="http://www.w3.org/2000/svg" viewBox="0 0 12 12" style="margin-right: 4px;">
                            <path fill-rule="evenodd" clip-rule="evenodd"
                                d="M8 1.25a.75.75 0 000 1.5h1.09L7.54 4.298a.757.757 0 00-.058.066 4 4 0 10.968 1.112.752.752 0 00.15-.117L10.25 3.71V5a.75.75 0 001.5 0V2a.75.75 0 00-.75-.75H8zM5 10a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
                                fill="#168EF9"></path>
                        </svg>
                        <svg width="12" v-else height="12" fill="none" xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 12 12" style="margin-right: 4px;">
                            <mask id="woman_svg__a" maskUnits="userSpaceOnUse" x="-2" y="-2" width="16" height="16"
                                style="mask-type: alpha;">
                                <path fill="#C4C4C4" d="M-2-2h16v16H-2z"></path>
                            </mask>
                            <g mask="url(#woman_svg__a)" stroke="#F5588E" stroke-width="1.5" stroke-linecap="round"
                                stroke-linejoin="round">
                                <circle cx="7.2" cy="4.896" r="3.25"></circle>
                                <path d="M1.617 10.511l3.115-3.115M1.904 7.396l2.828 2.829"></path>
                            </g>
                        </svg>
                        {{ userInfo.age }}岁</span>

                    <span class="location tag">
                        {{ userInfo.location }}</span>
                </p>
                <p class="signature" :title="userInfo.signature">{{ userInfo.signature }}</p>
            </div>
            <div class="tool" v-if="userInfo.uid">
                <el-button @click="isChange = true">编辑资料</el-button>
                <el-button @click="logout" type="danger">退出登录</el-button>
            </div>
        </header>

        <main>
            <el-tabs v-model="activeChoice">
                <el-tab-pane :disabled="!userInfo.uid" class="video-container"
                    :label="`作品 ${userInfo.uid ? userInfo.workCount : ''}`" name="work">
                    <VideoBox v-for="item in videoList" :key="item.vid" v-bind="item" :ratio="4/3" :waterfall="false"></VideoBox>
                </el-tab-pane>
                <el-tab-pane :disabled="!userInfo.uid" class="video-container" :label="`喜欢 ${userInfo.likeCount || ''}`"
                    name="like">
                </el-tab-pane>
                <el-tab-pane :disabled="!userInfo.uid" class="video-container" :label="`收藏 ${userInfo.collectCount || ''}`"
                    name="collect">
                </el-tab-pane>
                <el-tab-pane :disabled="!userInfo.uid" class="video-container"
                    :label="`观看历史 ${userInfo.historyCount || ''}`" name="history">
                </el-tab-pane>
            </el-tabs>

            <el-empty description="点击右上角按钮进行登录" v-if="!userInfo.uid" :image-size="180">
            </el-empty>
        </main>

        <ChangeInfo v-if="isChange" @close="isChange = false"></ChangeInfo>
    </el-container>
</template>

<script>
import { mapState } from 'vuex'
import ChangeInfo from './ChangeInfo'
import { VideoPublishList } from '@/api/video'
import VideoBox from '@/components/VideoBox.vue'
export default {
    name: 'Me',
    components: {
        ChangeInfo,
        VideoBox
    },
    computed: {
        ...mapState({
            userInfo: state => state.user.userInfo
        })
    },
    data() {
        return {
            activeChoice: '',
            isChange: false,
            videoList: [],
        }
    },
    methods: {
        getData() {
            console.log(this.activeChoice);
        },
        logout() {
            this.$store.dispatch('UserLogout')
        },
        async getVideoList() {
            const res = await VideoPublishList()
            if (res.status !== 200) {
                this.$message.error(res.data.msg)
                return
            }
            this.videoList = res.data.data.videos
        }
    },
    watch: {
        activeChoice(newVal, oldVal) {
            if (newVal === oldVal) return
            this.getData();
        },
        'userInfo.uid': {
            immediate: true,
            handler(newVal, oldVal) {
                if (newVal) {
                    this.activeChoice = 'work'
                } else {
                    this.activeChoice = ''
                }
            }
        }
    },
    created() {
        this.getVideoList()
    }
}
</script>

<style lang="less" scoped>
.el-container {
    padding: 0 40px;
    flex-direction: column;
    min-width: 720px;
    max-width: 1280px;
    margin: 8px auto;

    header {
        display: flex;
        align-items: center;
        min-width: max-content;

        .el-avatar {
            font-size: 28px;
        }

        .userInfo {
            margin-left: 32px;
            min-height: 120px;
            display: flex;
            flex-direction: column;
            justify-content: center;
            margin-right: auto;

            .username {
                color: @white;
                margin-bottom: 16px;
                font-size: 20px;
                font-weight: 500;
                line-height: 28px;
                max-width: 300px;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                margin: 0;
            }

            .list {
                display: flex;
                align-items: center;
                margin-top: 4px;
                width: 100%;
                white-space: nowrap;

                span {
                    font-size: 14px;
                    line-height: 22px;
                    margin-right: 6px;
                }

                i {
                    font-size: 16px;
                    line-height: 24px;
                    color: @white;
                    font-style: normal;
                }

                .el-divider {
                    background-color: @gray-3;
                    transform: scaleX(.5);
                    margin: 0 16px;
                }
            }

            .other {
                align-items: center;
                display: flex;
                height: 20px;
                margin-top: 12px;
                width: 100%;

                .account {
                    font-size: 12px;
                    line-height: 20px;
                    margin-right: 20px;
                }

                .tag {
                    background: @nav-bg-hover;
                    border-radius: 4px;
                    color: @white;
                    display: flex;
                    align-items: center;
                    font-size: 12px;
                    height: 20px;
                    line-height: 20px;
                    margin-right: 4px;
                    padding: 0 8px;
                }
            }

            .signature {
                font-size: 12px;
                line-height: 20px;
                color: @gray-1;
                margin-top: 6px;
                margin-right: 10px;
                white-space: nowrap;
                max-width: 300px;
                overflow: hidden;
                text-overflow: ellipsis;
                cursor: pointer;
            }

        }

        .tool {
            align-self: flex-end;

            .el-button {
                background: @gray-3;
                border: none;
                color: @white;
                font-size: 13px;
                font-weight: 400;
            }

            .el-button--danger {
                background: @red;
            }
        }

    }

    main {

        .el-tabs {
            margin: 11px 0;
            box-sizing: content-box;
            padding-top: 16px;

            /deep/ .el-tabs__item {
                cursor: pointer;
                color: @gray-2;
                font-size: 18px;
                height: unset;
                line-height: 3;

                &:hover {
                    color: @gray-2;
                }

                &.is-active {
                    color: @white;
                }
            }

            /deep/ .el-tabs__nav-wrap {
                &::after {
                    background-color: @gray-3;
                    transform: scaleY(.5);
                    height: 1px;
                }
            }

            /deep/ .el-tabs__active-bar {
                display: none;
                background-color: @white;
            }

            /deep/ .el-tabs__content {
                .video-container {
                    display: grid;
                    grid-template-columns: repeat(6, 1fr);
                    grid-gap: 16px;

                    @media screen and (max-width: 1280px) {
                        grid-template-columns: repeat(3, 1fr);
                    }

                }
            }
        }

    }
}
</style>