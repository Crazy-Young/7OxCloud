<template>
    <div class="recommend">
        <div class="xgplayer-container" ref="XgplayerContainer">
            <swiper :options="swiperOptions" ref="swiper" @mousedown.native="handleTouchStart">
                <swiper-slide v-for="(item, index) in videoList" :key="item.id">
                    <XgPlayer class="xgplayer-item" v-bind="item" ref="XgPlayer" :ignores="['cssFullscreen']"
                        :autoplay="index === activeIndex">
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
                        <template #operation>
                            <div class="xgplayer-operation" @pointerdown.stop @mousedown.stop>
                                <div class="xgplayer-operation-item avatar">
                                    <router-link :to="`/user/${item.author.uid}`">
                                        <el-avatar :size="40" :src="item.author.avatar"></el-avatar>
                                    </router-link>
                                    <button class="follow" v-if="!item.author.isFollow"
                                        @click="handleFollow(item.author)"></button>
                                </div>
                                <div class="xgplayer-operation-item like" @click="handleLike(item)">
                                    <LikeIcon :isLike="item.isLike"></LikeIcon>
                                    <span>{{ item.likeCount }}</span>
                                </div>

                                <div class="xgplayer-operation-item comment">
                                    <CommentIcon></CommentIcon>
                                    <span>{{ item.commentCount }}</span>
                                </div>

                                <div class="xgplayer-operation-item collect" @click="handleCollect(item)">
                                    <CollectIcon :isCollect="item.isCollect"></CollectIcon>
                                    <span>{{ item.collectCount }}</span>
                                </div>
                            </div>
                        </template>
                    </XgPlayer>
                    <div class="xgplayer-comment">
                    </div>
                </swiper-slide>
            </swiper>
        </div>
        <div class="xgplayer-playswitch">
            <div class="xgplayer-playswitch-tab">
                <div class="xgplayer-playswitch-tab-item xgplayer-playswitch-prev" @click="prev"
                    :class="{ disabled: activeIndex === 0 }"><svg width="26" height="26" fill="none"
                        xmlns="http://www.w3.org/2000/svg" class="" viewBox="0 0 26 26">
                        <g filter="url(#newAbove_svg__filter0_d_1651_162618)">
                            <path
                                d="M7.269 16.316a1.393 1.393 0 010-1.97l5.056-5.055a1.393 1.393 0 011.97 0l.011.011 5.045 5.045a1.393 1.393 0 11-1.97 1.97l-4.071-4.072-4.071 4.071a1.393 1.393 0 01-1.97 0z"
                                fill="#fff"></path>
                        </g>
                    </svg></div>
                <div class="xgplayer-playswitch-tab-item xgplayer-playswitch-next" @click="next"
                    :class="{ disabled: activeIndex === videoList.length - 1 }">
                    <svg width="26" height="26" fill="none" xmlns="http://www.w3.org/2000/svg" class="" viewBox="0 0 26 26">
                        <g filter="url(#newBelow_svg__filter0_d_1646_162559)">
                            <path
                                d="M7.269 9.29a1.393 1.393 0 000 1.97l5.056 5.056a1.393 1.393 0 001.97 0l.011-.011 5.045-5.045a1.393 1.393 0 10-1.97-1.97l-4.071 4.072L9.239 9.29a1.393 1.393 0 00-1.97 0z"
                                fill="#fff"></path>
                        </g>
                    </svg>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import XgPlayer from "@/components/XgPlayer";
import { swiper, swiperSlide } from 'vue-awesome-swiper'
import { VideoListByTime } from '@/api/video'
import formatTime from "@/utils/formatTime";
import CollectIcon from "./CollectIcon";
import LikeIcon from "./LikeIcon";
import CommentIcon from "./CommentIcon";
export default {
    name: "Recommend",
    components: {
        XgPlayer,
        swiper,
        swiperSlide,
        CollectIcon,
        LikeIcon,
        CommentIcon
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
                }
            }
        }
    },
    data() {
        return {
            activeIndex: 0,
            videoList: [],
            latestTime: null
        }
    },
    methods: {
        formatTime,
        prev() {
            this.$refs.swiper.swiper.slidePrev();
        },
        next() {
            this.$refs.swiper.swiper.slideNext();
        },
        handleInteractive(event) {
            if (event.deltaY > 0 || event.key === 'ArrowDown') {
                this.next()
            } else if (event.deltaY < 0 || event.key === 'ArrowUp') {
                this.prev()
            }
        },
        activePlay() {
            this.$refs.XgPlayer.forEach(item => {
                if (item.player.paused) {
                    return
                }
                item.pause()
            })
            this.activeIndex = this.$refs.swiper.swiper.activeIndex
            this.$refs.XgPlayer[this.activeIndex].play()
        },
        handleTouchStart(event) {
            this.$refs.swiper.swiper.onTouchStart(event)
        },
        async getVideoList() {
            const res = await VideoListByTime(this.latestTime)
            if (res.status !== 200) {
                this.$message.error(res.data.msg)
                return
            }
            this.videoList.push(...res.data.data.videos)
            this.latestTime = res.data.data.nextTime
        },
        handleLike(item) {
            item.isLike = !item.isLike
        },
        handleCollect(item) {
            item.isCollect = !item.isCollect
        },
        handleFollow(item) {
            item.isFollow = !item.isFollow
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
    }
};
</script>

<style lang="less" scoped>
.recommend {
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
        min-width: 320px;

        .swiper-container {
            height: 100%;
        }

        .xgplayer-item {
            position: relative;

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
                        background-color: @nav-bg-hover;
                    }
                }

                .xgplayer-title-user {
                    font-size: 20px;
                    display: flex;
                    align-items: center;
                    margin-top: 8px;

                    .username {
                        max-width: calc(100% - 100px);
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;

                        a {
                            color: @white;
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
                        background-color: @nav-bg-hover;
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
                color: white;
                pointer-events: all;
                filter: drop-shadow(0 0 3px @gray-3);

                .xgplayer-operation-item {
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    margin-bottom: 10px;
                }

                .avatar {
                    position: relative;
                    margin-bottom: 30px;

                    .follow {
                        position: absolute;
                        bottom: 0;
                        left: 50%;
                        transform: translate(-50%, 50%);
                        height: 20px;
                        width: 20px;
                        border: none;
                        outline: none;
                        background-color: @red;
                        border-radius: 50%;
                        cursor: pointer;

                        &::before {
                            content: '';
                            position: absolute;
                            top: 50%;
                            left: 50%;
                            transform: translate(-50%, -50%);
                            width: 50%;
                            height: 10%;
                            background-color: @white;
                            border-radius: 999px;
                        }

                        &::after {
                            content: '';
                            position: absolute;
                            top: 50%;
                            left: 50%;
                            transform: translate(-50%, -50%);
                            width: 10%;
                            height: 50%;
                            background-color: @white;
                            border-radius: 999px;
                        }
                    }
                }
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
            }
        }

    }
}
</style>