<template>
    <div class="player-with-method" :style="{ height }">
        <XgPlayer class="xgplayer-item" :autoHide="false" :vid="vid" :playUrl="playUrl" :coverUrl="coverUrl"
            :autoplay="autoplay" ref="XgPlayer">
            <template>
                <slot></slot>
            </template>
            <template #title>
                <slot name="title"></slot>
            </template>
            <template #operation>
                <UserOperation :isUser="isUser" :vid="vid" :author="author" @handleComment="isComment = !isComment"
                    :likeCount="likeCount" :commentCount="commentCount" :collectCount="collectCount" :isLike="isLike"
                    :isCollect="isCollect" ref="UserOperation">
                </UserOperation>
            </template>
        </XgPlayer>
        <CommentWindow @close="isComment = false" :show="isComment" :class="{ active: isComment }" :bg="coverUrl"
            class="xgplayer-comment" :vid="vid" @changeCommentLength="changeCommentCount"></CommentWindow>
    </div>
</template>

<script>
import UserOperation from "@/components/UserOperation.vue";
import CommentWindow from "@/components/CommentWindow.vue";
import XgPlayer from "@/components/XgPlayer.vue";
export default {
    name: "PlayerWithMethod",
    components: {
        UserOperation,
        CommentWindow,
        XgPlayer
    },
    data() {
        return {
            isComment: false
        }
    },
    props: {
        vid: {
            type: Number | String,
            required: true
        },
        autoplay: {
            type: Boolean,
            default: false
        },
        coverUrl: {
            type: String,
            default: ""
        },
        playUrl: {
            type: String,
            default: ""
        },
        height: {
            type: String,
            default: "100%"
        },
        isUser: {
            type: Boolean,
            default: false
        },
        author: {
            type: Object,
            default: () => {
                return {}
            }
        },
        likeCount: {
            type: Number,
            default: 0
        },
        commentCount: {
            type: Number,
            default: 0
        },
        collectCount: {
            type: Number,
            default: 0
        },
        isLike: {
            type: Boolean,
            default: false
        },
        isCollect: {
            type: Boolean,
            default: false
        },
    },
    methods: {
        changeCommentCount(count) {
            this.$refs.UserOperation.changeCommentCount(count)
        },
        play() {
            this.$refs.XgPlayer.play()
        },
        pause() {
            this.$refs.XgPlayer.pause()
        }
    }
}
</script>

<style lang="less" scoped>
.player-with-method {
    border-radius: 12px;
    display: flex;
    overflow: hidden;
    position: relative;
    width: 100%;
    height: 100%;

    .xgplayer-operation {
        position: absolute;
        bottom: 70px;
        right: 20px;
    }

    .xgplayer-comment {
        width: 0;
        transition: width 0.3s;
        max-width: 420px;

        &.active {
            width: 50%;

            @media screen and (max-width: 928px) {
                width: 80%;
            }
        }
    }
}
</style>