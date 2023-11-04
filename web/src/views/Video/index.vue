<template>
    <div class="video" @dblclick="$router.go(-1)">
        <!-- 视频播放器: 含有操作控制和评论 -->
        <PlayerWithMethod v-bind="video" height="100%" v-if="video.playUrl" isUser autoplay>
            <template>
                <!-- 左上角关闭 -->
                <div class="back" @click="$router.go(-1)">
                    <CloseIcon></CloseIcon>
                </div>
            </template>
            <template #title>
                <!-- 左下角标题 -->
                <div class="xgplayer-title">
                    <span class="xgplayer-title-location">
                        <i class="el-icon-location-outline"></i>
                        <strong>{{ video.author.location }}</strong>
                    </span>
                    <p class="xgplayer-title-user">
                        <span class="username">
                            <router-link :to="`/user/${video.author.uid}`">
                                @{{ video.author.username }}
                            </router-link>
                        </span>
                        <span class="and">·</span>
                        <span class="time">{{ formatTime(video.publishTime) }}</span>
                    </p>
                    <h3 class="xgplayer-title-text">{{ video.description }}</h3>
                    <p class="xgplayer-title-topics">
                        <Topic v-for="topic in video.topics" :key="topic.id" v-bind="topic"></Topic>
                    </p>
                </div>
            </template>
        </PlayerWithMethod>
        <!-- 视频加载前Loading -->
        <Loading v-else></Loading>
    </div>
</template>

<script>
import { VideoInfo, ViewVideo } from '@/api/video';
import CloseIcon from '@/icons/CloseIcon.vue';
import formatTime from '@/utils/formatTime';
import Topic from '@/components/Topic.vue';
import PlayerWithMethod from '@/components/PlayerWithMethod.vue';
import Loading from '@/components/Loading.vue';
export default {
    name: 'Video',
    data() {
        return {
            // 视频信息
            video: {}
        }
    },
    components: {
        CloseIcon,
        Topic,
        PlayerWithMethod,
        Loading
    },
    computed: {
        // 视频id
        vid() {
            return this.$route.params.vid
        }
    },
    methods: {
        formatTime,
        // 获取视频信息
        async getVideoInfo() {
            const res = await VideoInfo(this.vid)
            if (res.status === 200) {
                // 获取视频信息成功
                this.video = res.data.data
            } else {
                // 获取视频信息失败
                this.$message.error(res.data.msg)
            }
        },
    },
    mounted() {
        this.getVideoInfo()
        // 记录用户的浏览历史
        ViewVideo(this.vid)
    }
}
</script>

<style lang="less" scoped>
.video {
    position: relative;
    width: 100vw;
    height: 100vh;
    display: flex;

    .back {
        position: absolute;
        top: 50px;
        left: 50px;
        width: 50px;
        height: 50px;
        border-radius: 50%;
        background-color: @transparent-darker;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all .2s;
        cursor: pointer;

        svg {
            fill: @gray-2;
        }

        &:hover {
            transform: scale(1.2);
            background-color: @gray-1;
        }
    }

    .loading {
        position: absolute;
        inset: 0;
        height: fit-content;
        width: fit-content;
        color: @white;
        margin: auto;
    }

    .xgplayer-title {
        position: absolute;
        bottom: 48px;
        color: @white;
        width: 100%;
        padding: 20px 100px 20px 20px;

        .xgplayer-title-location {
            font-size: 14px;
            border-radius: 4px;
            display: flex;
            align-items: center;
            overflow: hidden;
            width: fit-content;
            height: 24px;

            i {
                height: 24px;
                width: 24px;
                background-color: @red;
                display: flex;
                justify-content: center;
                align-items: center;
            }

            strong {
                display: inline-block;
                height: 100%;
                line-height: 24px;
                padding: 0 8px;
                background-color: @transparent-dark;
            }
        }

        .xgplayer-title-user {
            font-size: 20px;
            display: flex;
            align-items: center;
            margin-top: 8px;
            white-space: nowrap;

            .username {

                a {
                    color: @white;
                    max-width: calc(100% - 100px);
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }
            }

            .and {
                margin: 0 5px;
            }

            .time {
                font-size: 12px;
                white-space: nowrap;
            }
        }

        .xgplayer-title-text {
            font-size: 18px;
            line-height: 1.5;
            font-weight: 100;
            margin-top: 6px;
        }

        .xgplayer-title-topics {
            font-size: 14px;
            display: flex;
            flex-wrap: wrap;
            margin-top: 2px;

            .topic-item {
                color: @white;
            }
        }
    }

}
</style>