<template>
    <!-- 用户页面 -->
    <el-container>
        <!-- 顶部用户信息 -->
        <header>
            <el-avatar :size="120" :src="userInfo.avatar">{{ userInfo.username || '加载中...' }}</el-avatar>
            <div class="userInfo">
                <h1 class="username">{{ userInfo.username || '加载中...' }}</h1>
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
                        <GenderIcon :gender="userInfo.gender"></GenderIcon>
                        <span>{{ userInfo.age }}岁</span>
                    </span>

                    <span class="location tag">
                        {{ userInfo.location }}</span>
                </p>
                <p class="signature" :title="userInfo.signature">{{ userInfo.signature }}</p>
            </div>
            <div class="tool" v-if="userInfo.uid">
                <el-button @click="follow" type="danger" :class="{ 'follow': userInfo.isFollow }">{{ userInfo.isFollow ?
                    '已关注' : '关注' }}</el-button>
            </div>
        </header>

        <!-- 视频列表 -->
        <main>
            <el-tabs v-model="activeChoice">
                <!-- 其他用户只能查看别人的作品 -->
                <el-tab-pane :disabled="!userInfo.uid" class="video-container"
                    :label="`作品 ${userInfo.uid ? userInfo.workCount : ''}`" name="work">
                    <VideoBox v-for="item in videoList" :key="item.vid" v-bind="item" :ratio="4 / 3"
                        :waterfall="false" ref="VideoBox">
                    </VideoBox>
                </el-tab-pane>
            </el-tabs>

            <DataStatus :isDataEnd="isDataEnd" :length="videoList.length" ref="loading" @getVideoList="getVideoList"></DataStatus>
        </main>

    </el-container>
</template>

<script>
import { FollowUser, UserInfo } from '@/api/user'
import { VideoPublishList } from '@/api/video'
import { mapState } from 'vuex'
import VideoBox from '@/components/VideoBox.vue'
import videoList from '@/mixins/videoList'
import GenderIcon from '@/icons/GenderIcon.vue'
import DataStatus from '@/components/DataStatus.vue'
export default {
    name: 'User',
    mixins: [videoList],
    components: {
    VideoBox,
    GenderIcon,
    DataStatus
},
    data() {
        return {
            activeChoice: 'work',
            userInfo: {},

        }
    },
    methods: {
        // 获取用户信息
        getUserInfo() {
            UserInfo(this.uid).then(res => {
                if (res.status === 200)
                    this.userInfo = res.data.data
                else this.$message.error(res.data.msg)
            })
        },
        // 关注用户
        follow() {
            // 发送关注请求
            FollowUser({
                userId: this.userInfo.uid,
                // 如果没有关注，那么type为1表示关注用户
                // 如果已经关注，那么type为2表示取消关注
                type: !this.userInfo.isFollow ? 1 : 2
            })
            // 更新关注状态
            this.userInfo.isFollow = !this.userInfo.isFollow
            // 更新关注数量
            this.userInfo.fanCount += this.userInfo.isFollow ? 1 : - 1
        },
        getVideoList() {
            return this.reqVideoList(() => VideoPublishList(this.uid, this.latestTime))
        }
    },
    computed: {
        // 当前用户ID
        uid() {
            return this.$route.params.uid
        },
        ...mapState({
            myInfo: state => state.user.userInfo
        })
    },
    mounted() {
        // 判断当前用户是否为自己，如果是，跳转到个人中心
        if (this.uid == this.myInfo.uid) {
            this.$router.replace({ name: 'Me' })
        }
        this.getUserInfo()
    },
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
                    background: @transparent-dark;
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

                &.follow {
                    background-color: @transparent-dark;
                }
            }

            .el-button--danger {
                background: @red;
                width: 78px;
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
</style>