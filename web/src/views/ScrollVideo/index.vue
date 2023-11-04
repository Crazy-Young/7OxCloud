<template>
    <!-- 视频布局: 滚动 -->
    <div class="scroll-video">
        <div class="search" id="search" ref="search">
            <div class="video-item" v-for="item in videoList" :key="item.vid">
                <article class="video-info">
                    <div class="video-info-header">
                        <router-link :to="`/user/${item.author.uid}`">
                            <el-avatar class="video-info-avatar" :size="40" :src="item.author.avatar"></el-avatar>
                        </router-link>
                        <div class="video-info-user">
                            <p class="video-info-username">
                                <router-link :to="`/user/${item.author.uid}`">{{ item.author.username }}</router-link>
                            </p>
                            <p class="video-info-time">{{ formatTime(item.publishTime) }}</p>
                        </div>
                        <el-button class="video-info-follow" :class="{ 'follow': item.author.isFollow }"
                            @click="handleFollow(item)" v-if="item.author.uid !== userInfo.uid">{{ item.author.isFollow
                                ? '已关注' : '关注' }}</el-button>
                    </div>
                    <p class="video-info-desc">
                        <router-link :to="`/video/${item.vid}`"><span>{{ item.description }}</span></router-link>
                    </p>
                </article>
                <PlayerWithMethod v-bind="item" height="560px" :isUser="false"></PlayerWithMethod>
                <div class="video-topics">
                    <Topic v-for="topic in item.topics" :key="topic.id" v-bind="topic" :size="14"></Topic>
                </div>
            </div>
        </div>
        <DataStatus :isDataEnd="isDataEnd" :length="videoList.length" ref="loading" @getVideoList="getVideoList"></DataStatus>
    </div>
</template>

<script>
import XgPlayer from '@/components/XgPlayer.vue';
import formatTime from '@/utils/formatTime';
import { mapState } from 'vuex';
import Topic from '@/components/Topic.vue';
import DataStatus from '@/components/DataStatus.vue';
import UserOperation from '@/components/UserOperation.vue';
import { FollowUser } from '@/api/user';
import PlayerWithMethod from '@/components/PlayerWithMethod.vue';
export default {
    name: "ScrollVideo",
    props: {
        videoList: {
            type: Array,
            default: () => []
        },
        isDataEnd: {
            type: Boolean,
            default: false
        }
    },
    components: {
        XgPlayer,
        Topic,
        DataStatus,
        UserOperation,
        PlayerWithMethod
    },
    methods: {
        formatTime,
        getVideoList() {
            this.$emit("getVideoList");
        },
        handleFollow(item) {
            FollowUser({
                userId: item.author.uid,
                type: !item.author.isFollow ? 1 : 2
            })
            item.author.isFollow = !item.author.isFollow
        },
        handleComment(vid) {
            this.$router.push({
                name: 'Video',
                params: {
                    vid
                },
                query: {
                    comment: true
                }
            })
        }
    },
    computed: {
        ...mapState({
            userInfo: state => state.user.userInfo
        })
    }
}
</script>

<style lang="less" scoped>
.scroll-video {
    position: relative;
    max-width: 960px;
    margin: auto;

    .search {
        position: relative;
        margin: 0 30px;
        min-width: 480px;
        padding-bottom: 30px;
        min-height: 100%;

        .video-item {
            margin: 10px 0 50px;

            .video-info {
                color: @white;

                .video-info-header {
                    display: flex;
                    align-items: center;

                    .video-info-avatar {
                        margin-right: 20px;
                    }

                    .video-info-user {
                        .video-info-username {
                            font-size: 16px;

                            a {
                                color: @white;
                            }
                        }

                        .video-info-time {
                            font-size: 14px;
                            color: @gray-2;
                        }
                    }

                    .video-info-follow {
                        margin-left: auto;
                        background-color: @red;
                        color: @white;
                        border: unset;
                        padding: 8px 0;
                        width: 64px;

                        &.follow {
                            background-color: @transparent-dark;
                        }
                    }
                }

                .video-info-desc {
                    font-size: 16px;
                    color: @gray-1;
                    margin: 10px 0;
                    line-height: 1.5;

                    a {
                        color: @white;
                    }
                }
            }
        }

        .video-topics {
            display: flex;
            flex-wrap: wrap;
        }

    }
}
</style>