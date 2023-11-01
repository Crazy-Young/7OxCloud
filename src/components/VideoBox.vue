<template>
    <article class="video">
        <div class="video-content" :style="{
            paddingTop: `${ratio * 100}%`
        }">
            <XgPlayer :playUrl="playUrl" :coverUrl="coverUrl"
                :ignores="['fullscreen', 'definition', 'i18n', 'cssfullscreen', 'playbackrate']" muted :vid="vid"
                class="video-content-xg" ref="video" ctrlpos="bottom"></XgPlayer>
        </div>
        <div class="video-info">
            <div class="video-info__header" :class="{ 'fix-height': !waterfall }">
                <span>{{ description }}</span>
            </div>
            <div class="video-info__footer" v-if="author">
                <span>
                    <router-link class="author" :to="{ name: 'User', params: { uid: author.uid } }">
                        @ <span>{{ author.username }}</span>
                    </router-link>
                    <span> · </span>
                    <span class="time">{{ formatTime(publishTime) }}</span>
                </span>
            </div>
            <div class="video-info__topic" v-if="topics && topics.length" :class="{ 'no-warp': !waterfall }">
                <span class="topic" v-for="item in topics" :key="item.id">
                    <router-link :to="{ name: 'Topic', params: { tid: item.id } }"># {{ item.description }}</router-link>
                </span>
            </div>
        </div>
    </article>
</template>

<script>
import XgPlayer from './XgPlayer.vue';
import formatTime from '@/utils/formatTime';
export default {
    name: 'VideoBox',
    components: { XgPlayer },
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
        }
    },
    methods: {
        formatTime,
        play() {
            this.$refs.video.play();
        },
        pause() {
            this.$refs.video.pause();
        }
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
        }
    }

    .video-info {
        width: 100%;
        padding: 12px;
        background-color: @gray-4;

        .video-info__header {
            color: @gray-1;
            font-size: 15px;
            cursor: pointer;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 2;
            overflow: hidden;
            text-overflow: ellipsis;

            &.fix-height {
                height: 40px;
            }
        }

        .video-info__footer {
            color: @gray-2;
            font-size: 14px;
            margin-top: 6px;

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

            .topic {
                color: @gray-2;
                font-size: 12px;
                margin-right: 8px;
                margin-top: 8px;
                cursor: pointer;
                background-color: @nav-bg-hover;
                padding: 2px 6px;
                border-radius: 4px;
            }

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