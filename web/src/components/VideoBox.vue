<template>
    <article class="video">
        <div class="video-content" :style="{
            paddingTop: `${ratio * 100}%`
        }" @mouseenter="handleMouseEnter" @mouseleave="handleMouseLeave">
            <XgPlayer :playUrl="playUrl" :coverUrl="coverUrl"
                :ignores="['fullscreen', 'definition', 'i18n', 'cssfullscreen', 'playbackrate']" muted :vid="vid"
                class="video-content-xg" ref="video" ctrlpos="bottom">
                <template>
                    <div class="video-like-count">
                        <LikeHollowIcon></LikeHollowIcon>
                        {{ likeCount }}
                    </div>
                    <div class="video-delete" @click="handleDelete" v-if="isDelete">
                        <i class="el-icon-delete"></i>
                    </div>
                </template>
            </XgPlayer>
        </div>
        <div class="video-info">
            <div class="video-info__header" :class="{ 'fix-height': !waterfall }">
                <span class="desc">
                    <router-link class="title" :to="`/video/${vid}`">{{ description }}</router-link>
                </span>
                <Topic v-for="item in topics" :key="item.id" v-bind="item" class="topic"></Topic>
            </div>
            <div class="video-info__footer" v-if="author">
                <span>
                    <router-link class="author" :to="{ name: 'User', params: { uid: author.uid } }">
                        @ <span>{{ author.username }}</span>
                    </router-link>
                    <span> · </span>
                    <span v-if="author.isFollow" class="follow">已关注</span>
                    <span class="time" v-else>{{ formatTime(publishTime) }}</span>
                </span>
            </div>
        </div>
    </article>
</template>

<script>
import XgPlayer from './XgPlayer.vue';
import formatTime from '@/utils/formatTime';
import Topic from "@/components/Topic.vue"
import LikeHollowIcon from "@/icons/LikeHollowIcon"
import { mapState } from 'vuex';
import { VideoDelete } from '@/api/video';
export default {
    name: 'VideoBox',
    data() {
        return {
            timer: null
        }
    },
    components: {
        XgPlayer,
        Topic,
        LikeHollowIcon
    },
    computed: {
        ...mapState({
            userInfo: state => state.user.userInfo
        })
    },
    props: {
        vid: {
            type: String | Number,
        },
        publishTime: {
            type: Number,
        },
        playUrl: {
            type: String,
            default: () => ''
        },
        coverUrl: {
            type: String,
            default: () => ''
        },
        description: {
            type: String,
            default: '视频标题'
        },
        author: {
            type: Object,
            default: () => { }
        },
        ratio: {
            type: Number,
            default: 3 / 4
        },
        topics: {
            type: Array,
            default: () => []
        },
        waterfall: {
            type: Boolean,
            default: true
        },
        likeCount: {
            type: Number,
            default: 0
        },
        isDelete: {
            type: Boolean,
            default: false
        }
    },
    methods: {
        formatTime,
        play() {
            this.$refs.video.play();
        },
        pause() {
            this.$refs.video.pause();
        },
        handleMouseEnter() {
            this.timer = setTimeout(() => {
                this.play();
            }, 100)
        },
        handleMouseLeave() {
            clearTimeout(this.timer);
            this.pause();
        },
        handleDelete() {
            VideoDelete(this.vid).then(res => {
                if (res.status === 200) {
                    this.$message.success(res?.data?.msg);
                    this.$emit('deleteVideo', this.vid);
                }else {
                    this.$message.error(res?.data?.msg);
                }
            })
        }
    },
    beforeDestroy() {
        clearTimeout(this.timer);
        this.pause();
    }
}
</script>

<style lang="less" scoped>
.video {
    border-radius: 10px;
    overflow: hidden;
    height: fit-content;

    .video-content {
        position: relative;

        .video-content-xg {
            position: absolute;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;

            .video-like-count {
                position: absolute;
                top: 12px;
                left: 12px;
                color: @white;
                display: flex;
                align-items: center;
                gap: 5px;
                font-size: 16px;
                filter: drop-shadow(0 0 1px @gray-2);
            }

            .video-delete {
                position: absolute;
                top: 12px;
                right: 12px;
                color: @white;
                font-size: 16px;
                filter: drop-shadow(0 0 1px @gray-2);
            }
        }
    }

    .video-info {
        width: 100%;
        padding: 12px;
        background-color: @gray-4;

        .video-info__header {
            color: @gray-1;
            font-size: 15px;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 2;
            overflow: hidden;
            text-overflow: ellipsis;
            line-height: 26px;

            &.fix-height {
                height: 52px;
            }

            .desc {
                margin-right: 6px;
            }

            .topic {
                margin: 0 6px 0 0;
            }
        }

        .video-info__footer {
            color: @gray-2;
            font-size: 14px;
            margin-top: 6px;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            line-height: 22px;

            .follow {
                background-color: @transparent-dark;
                line-height: 1;
                padding: 2px 6px;
                border-radius: 4px;
                font-size: 12px;
            }

            .author {
                cursor: pointer;
            }

            .time {
                font-size: 12px;
            }
        }

        .video-info__topic {
            display: flex;
            flex-wrap: wrap;
            margin-top: 4px;
            height: 28px;

            &.no-warp {
                flex-wrap: nowrap;
                white-space: nowrap;
                text-overflow: ellipsis;
                overflow: hidden;
                max-width: 100%;
            }
        }
    }
}
</style>