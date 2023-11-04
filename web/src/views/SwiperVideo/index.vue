<template>
    <!-- 视频布局: 轮播 -->
    <div class="swiper-video">
        <div class="xgplayer-container" ref="XgplayerContainer">
            <swiper :options="swiperOptions" ref="swiper" @mousedown.native="handleTouchStart" v-if="videoList.length">
                <swiper-slide v-for="(item, index) in videoList" :key="item.id">
                    <PlayerWithMethod v-bind="item" ref="XgPlayer" :autoplay="index === activeIndex" isUser>
                        <template #title>
                            <div class="xgplayer-title">
                                <span class="xgplayer-title-location">
                                    <i class="el-icon-location-outline"></i>
                                    <strong>{{ item.author.location }}</strong>
                                </span>
                                <p class="xgplayer-title-user">
                                    <span class="username">
                                        <router-link :to="`/user/${item.author.uid}`">
                                            @{{ item.author.username }}
                                        </router-link>
                                    </span>
                                    <span class="and">·</span>
                                    <span class="time">{{ formatTime(item.publishTime) }}</span>
                                </p>
                                <h3 class="xgplayer-title-text">{{ item.description }}</h3>
                                <p class="xgplayer-title-topics">
                                    <span v-for="topic in item.topics" class="topic" :key="topic.id">
                                        <router-link :to="`/topic/${topic.id}`">
                                            #{{ topic.description }}
                                        </router-link>
                                    </span>
                                </p>
                            </div>
                        </template>
                    </PlayerWithMethod>
                </swiper-slide>
            </swiper>
            <Loading v-else-if="!isDataEnd"></Loading>
            <el-empty v-else description="暂无视频" class="no-data"></el-empty>
        </div>
        <div class="xgplayer-playswitch">
            <div class="xgplayer-playswitch-tab">
                <div class="xgplayer-playswitch-tab-item xgplayer-playswitch-prev" @click="prev"
                    :class="{ disabled: activeIndex === 0 }">
                    <UpArrowIcon></UpArrowIcon>
                </div>
                <div class="xgplayer-playswitch-tab-item xgplayer-playswitch-next" @click="next"
                    :class="{ disabled: activeIndex === videoList.length - 1 || videoList.length === 0 }">
                    <DownArrowIcon></DownArrowIcon>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import XgPlayer from "@/components/XgPlayer";
import { swiper, swiperSlide } from 'vue-awesome-swiper'
import formatTime from "@/utils/formatTime";
import UpArrowIcon from "@/icons/UpArrowIcon.vue";
import DownArrowIcon from "@/icons/DownArrowIcon.vue";
import videoList from "@/mixins/videoList";
import UserOperation from "@/components/UserOperation.vue";
import CommentWindow from "@/components/CommentWindow.vue";
import PlayerWithMethod from '@/components/PlayerWithMethod.vue';
import { ViewVideo } from "@/api/video";
import Loading from "@/components/Loading.vue";
export default {
    name: "SwiperVideo",
    mixins: [videoList],
    props: {
        getVideoFn: Function,
        default: () => {
            return () => []
        }
    },
    components: {
        XgPlayer,
        swiper,
        swiperSlide,
        UpArrowIcon,
        DownArrowIcon,
        UserOperation,
        CommentWindow,
        PlayerWithMethod,
        Loading
    },
    computed: {
        swiperOptions() {
            return {
                direction: 'vertical',
                autoplay: false,
                observer: true,
                observeParents: true,
                on: {
                    slideChange: this.activePlay,
                    init: () => {
                        this.activePlay()
                    }
                }
            }
        }
    },
    data() {
        return {
            activeIndex: 0,
            isComment: false,
            commentText: '',
        }
    },
    methods: {
        formatTime,
        prev() {
            this.$refs.swiper?.swiper.slidePrev();
        },
        next() {
            this.$refs.swiper?.swiper.slideNext();
        },
        handleInteractive(event) {
            if (event.deltaY > 0 || event.key === 'ArrowDown') {
                this.next()
            } else if (event.deltaY < 0 || event.key === 'ArrowUp') {
                this.prev()
            }
        },
        activePlay() {
            this.$refs.XgPlayer?.forEach(item => {
                item.pause()
            })
            this.activeIndex = this.$refs.swiper?.swiper?.activeIndex || 0
            this.$refs.XgPlayer[this.activeIndex]?.play()
            ViewVideo(this.videoList[this.activeIndex].vid)
        },
        handleTouchStart(event) {
            this.$refs.swiper?.swiper.onTouchStart(event)
        },
        getVideoList() {
            return this.reqVideoList(() => this.getVideoFn(this.latestTime))
        }
    },
    mounted() {
        this.$refs.XgplayerContainer.addEventListener('wheel', this.handleInteractive)
        window.addEventListener('keydown', this.handleInteractive)
        this.getVideoList()
    },
    beforeDestroy() {
        this.$refs.XgplayerContainer.removeEventListener('wheel', this.handleInteractive)
        window.removeEventListener('keydown', this.handleInteractive)
    },
    watch: {
        activeIndex(newVal) {
            if (newVal === this.videoList.length - 1) {
                this.getVideoList()
            }
        }
    }
};
</script>

<style lang="less" scoped>
.swiper-video {
    display: flex;
    align-items: center;
    padding: 0 10px 20px 10px;
    height: 100%;

    .xgplayer-container {
        border-radius: 18px;
        position: relative;
        width: 100%;
        overflow: hidden;
        height: 100%;
        min-width: 600px;
        display: flex;
        justify-content: center;
        align-items: center;

        .swiper-slide {
            display: flex;
        }

        .swiper-container {
            height: 100%;
            width: 100%;
        }

        .xgplayer-item {
            position: relative;
            flex: 1 0 320px;

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

                    .topic {
                        margin-right: 4px;
                        margin-top: 4px;
                        background-color: @transparent-dark;
                        border-radius: 4px;
                        padding: 2px 6px;

                        a {
                            color: @gray-1;
                        }
                    }
                }
            }

            .xgplayer-operation {
                position: absolute;
                bottom: 80px;
                right: 20px;
            }
        }

        .xgplayer-comment {
            width: 0;
            transition: width .3s;

            &.active {
                width: 360px;
            }
        }
    }

    .xgplayer-playswitch {
        margin-left: 10px;

        .xgplayer-playswitch-tab {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            background-color: @gray-4;
            border-radius: 999px;

            .xgplayer-playswitch-tab-item {
                height: 44px;
                width: 44px;
                display: flex;
                justify-content: center;
                align-items: center;
                cursor: pointer;

                &.disabled {
                    cursor: not-allowed;
                    opacity: .3;
                }

                svg {
                    fill: @white;
                }
            }
        }

    }
}
</style>