<template>
    <div class="xgplayer-operation" @pointerdown.stop @mousedown.stop @dblclick.stop>
        <div class="xgplayer-operation-item avatar" v-if="isUser">
            <router-link :to="`/user/${uid}`">
                <el-avatar :size="40" :src="author.avatar"></el-avatar>
            </router-link>
            <button class="follow" v-if="uid !== userInfo.uid && !m_isFollow" @click="handleFollow"></button>
        </div>
        <div class="xgplayer-operation-item like" @click="handleLike">
            <LikeIcon :isLike="m_isLike"></LikeIcon>
            <span>{{ m_likeCount }}</span>
        </div>

        <div class="xgplayer-operation-item comment" @click="handleComment">
            <CommentIcon></CommentIcon>
            <span>{{ m_commentCount }}</span>
        </div>

        <div class="xgplayer-operation-item collect" @click="handleCollect">
            <CollectIcon :isCollect="m_isCollect"></CollectIcon>
            <span>{{ m_collectCount }}</span>
        </div>
    </div>
</template>

<script>
import CollectIcon from "@/icons/CollectIcon";
import LikeIcon from "@/icons/LikeIcon";
import CommentIcon from "@/icons/CommentIcon";
import { LikeVideo, CollectVideo } from "@/api/video";
import { FollowUser } from "@/api/user";
import { mapState } from "vuex";
export default {
    name: "UserOperation",
    methods: {
        async handleLike() {
            const res = await LikeVideo({
                vid: this.vid,
                type: !this.m_isLike ? 1 : 2
            });
            if (!res) return
            if (res.status === 200) {
                this.m_likeCount += this.m_isLike ? -1 : 1
                this.m_isLike = !this.m_isLike
                this.$emit("handleLike", this.vid);
            }
        },
        handleComment() {
            this.$emit("handleComment", this.vid);
        },
        async handleCollect() {
            const res = await CollectVideo({
                vid: this.vid,
                type: !this.m_isCollect ? 1 : 2
            })
            if (!res) return
            if (res.status === 200) {
                this.m_collectCount += this.m_isCollect ? -1 : 1
                this.m_isCollect = !this.m_isCollect
                this.$emit("handleCollect", this.vid);
            }
        },
        async handleFollow() {
            const res = await FollowUser({
                userId: this.uid,
                type: !this.m_isFollow ? 1 : 2
            })
            if (!res) return
            if (res.status === 200) {
                this.m_isFollow = !this.m_isFollow
                this.$emit("handleFollow", this.uid);
            }
        },
        changeCommentCount(count) {
            this.m_commentCount = count
        }
    },
    components: {
        CollectIcon,
        LikeIcon,
        CommentIcon
    },
    props: {
        vid: {
            type: Number | String,
            required: true
        },
        author: {
            type: Object,
            default: () => { }
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
        isUser: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            m_isLike: this.isLike,
            m_isCollect: this.isCollect,
            m_isFollow: this.author.isFollow ?? true,
            m_likeCount: this.likeCount,
            m_collectCount: this.collectCount,
            m_commentCount: this.commentCount,
            uid: this.author.uid || this.author.id
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
.xgplayer-operation {
    color: white;
    pointer-events: all;
    filter: drop-shadow(0 0 1px @gray-2);

    .xgplayer-operation-item {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .avatar {
        position: relative;
        margin-bottom: 20px;

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
</style>