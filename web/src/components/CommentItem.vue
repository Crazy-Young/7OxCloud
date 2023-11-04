<template>
    <!-- 单条评论 -->
    <div class="comment-item">
        <!-- 父评论用户头像 -->
        <div class="comment-item-left">
            <router-link :to="`/user/${author.id}`">
                <el-avatar :src="author.avatar" alt="" :size="32"></el-avatar>
            </router-link>
        </div>
        <div class="comment-item-right">
            <!-- 父评论用户名 -->
            <div class="comment-item-right-top">
                <router-link :to="`/user/${author.id}`">
                    <span class="comment-item-right-top-name">{{ author.username }}</span>
                </router-link>
            </div>
            <div class="comment-item-right-bottom">
                <!-- 父评论内容 -->
                <p class="comment-item-right-bottom-text">
                    <span>{{ content }}</span>
                </p>
                <!-- 父评论时间 -->
                <p class="comment-item-right-bottom-time">
                    <span>{{ formatTime(createdAt) }}</span>
                </p>
                <!-- 父评论回复按键以及删除 -->
                <p class="comment-item-right-bottom-reply">
                    <span @click="replyTo(cid, author.username, content)">
                        <CommentHollowIcon></CommentHollowIcon>回复
                    </span>
                    <!-- 当前评论作者与登录用户相同时显示删除 -->
                    <span @click="DeleteComment(cid)" v-if="author.id === userInfo.uid">
                        <i class="el-icon-delete"></i>删除评论
                    </span>
                </p>
                <!-- 子评论 -->
                <ul class="comment-item-right-bottom-list">
                    <!-- 遍历所有子评论 -->
                    <li class="comment-item-right-bottom-list-item" v-for="c in children.slice(0, sliceNum)" :key="c.cid">
                        <!-- 子评论用户头像 -->
                        <div class="comment-item-right-bottom-list-item-left">
                            <router-link :to="`/user/${c.author.id}`">
                                <el-avatar :src="c.author.avatar" alt="" :size="20"></el-avatar>
                            </router-link>
                        </div>
                        <!-- 子评论用户名和内容 -->
                        <div class="comment-item-right-bottom-list-item-right">
                            <div class="comment-item-right-bottom-list-item-right-top">
                                <!-- 子评论谁回复给谁 -->
                                <p class="comment-item-right-top">
                                    <router-link :to="`/user/${c.author.id}`">
                                        <span class="comment-item-right-bottom-list-item-right-top-name">
                                            {{ c.author.username }}
                                        </span>
                                    </router-link>
                                    <!-- 如果回复的是父评论，那么不显示回复给谁 -->
                                    <span class="comment-item-right-bottom-list-item-right-top-to" v-if="c.pid !== cid">
                                        <i class="el-icon-caret-right"></i>
                                    </span>
                                    <router-link :to="`/user/${c.pAuthor.id}`" v-if="c.pid !== cid">
                                        <span class="comment-item-right-bottom-list-item-right-top-name">
                                            {{ c.pAuthor.username }}
                                        </span>
                                    </router-link>
                                </p>
                            </div>
                            <div class="comment-item-right-bottom-list-item-right-bottom">
                                <!-- 子评论内容 -->
                                <p class="comment-item-right-bottom-text">
                                    <span>{{ c.content }}</span>
                                </p>
                                <!-- 子评论时间 -->
                                <p class="comment-item-right-bottom-time">
                                    <span>{{ formatTime(c.createdAt) }}</span>
                                </p>
                                <!-- 子评论回复按键以及删除 -->
                                <p class="comment-item-right-bottom-reply">
                                    <span @click="replyTo(c.cid, c.author.username, c.content)">
                                        <CommentHollowIcon></CommentHollowIcon>回复
                                    </span>
                                    <!-- 当前评论作者与登录用户相同时显示删除 -->
                                    <span @click="DeleteComment(c.cid)" v-if="c.author.id === userInfo.uid">
                                        <i class="el-icon-delete"></i>删除评论
                                    </span>
                                </p>
                            </div>
                        </div>
                    </li>
                </ul>
                <!-- 点击查看更多评论或者折叠评论 -->
                <div class="comment-item-more" @click="moreComment" v-if="children && children.length > sliceNum">
                    <i class="el-icon-caret-right"></i><span>点击查看更多评论</span>
                </div>
                <div class="comment-item-more" @click="lessComment" v-else-if="children && children.length > 0">
                    <i class="el-icon-caret-top"></i><span>点击折叠评论</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import formatTime from '@/utils/formatTime'
import CommentHollowIcon from '@/icons/CommentHollowIcon.vue'
import { mapState } from 'vuex'
import { DeleteComment } from '@/api/video'

export default {
    name: "CommentItem",
    components: {
        CommentHollowIcon
    },
    computed: {
        ...mapState({
            userInfo: state => state.user.userInfo
        })
    },
    props: {
        // 评论作者
        author: {
            type: Object,
            default() {
                return {}
            }
        },
        // 评论ID
        cid: {
            type: Number
        },
        // 评论内容
        content: {
            type: String
        },
        // 评论时间
        createdAt: {
            type: Number
        },
        // 子评论
        children: {
            type: Array,
            default() {
                return []
            }
        }
    },
    data() {
        return {
            // 查看更多评论的数量
            sliceNum: 0
        }
    },
    methods: {
        formatTime,
        moreComment() {
            // 每次点击查看更多评论的时候，sliceNum加3
            this.sliceNum = this.sliceNum + 3
        },
        lessComment() {
            // 每次点击折叠评论的时候，sliceNum重置为0
            this.sliceNum = 0
        },
        replyTo(cid, username, content) {
            // 回复评论
            this.$emit('replyTo', {
                cid,
                username,
                content
            })
        },
        async DeleteComment(cid) {
            // 发送删除评论的请求
            const res = await DeleteComment(cid)
            if (!res) return
            if (res.status === 200) {
                // 删除成功
                this.$message.success(res.data.msg)
                // 刷新评论
                this.$emit('refresh')
            }
        }
    }
}
</script>

<style lang="less" scoped>
.comment-item {
    display: flex;
    padding: 10px 0;
    width: 100%;

    .comment-item-left {
        width: 32px;
        height: 32px;
        border-radius: 50%;
        overflow: hidden;
        flex-shrink: 0;
    }

    .comment-item-right {
        padding-left: 10px;

        &-more {
            margin-top: 5px;
            font-size: 0.8rem;
            color: @gray-3;

            i {
                font-size: 0.8rem;
                color: @gray-3;
                margin-right: .3rem;
            }
        }

        .comment-item-right-top {
            white-space: break-spaces;

            .comment-item-right-bottom-list-item-right-top-to {
                margin: 0 5px;
                font-size: 12px;
            }

            &>a {
                flex: 1;
                max-width: 10rem;

                .comment-item-right-top-name {
                    font-size: 13px;
                    color: @gray-1;
                    font-weight: 700;
                    line-height: 22px;
                }
            }
        }

        .comment-item-right-bottom {

            .comment-item-right-bottom-text {
                font-size: 14px;
                color: @white;
                white-space: break-spaces;
                word-break: break-all;
                line-height: 20px;
            }

            .comment-item-right-bottom-time {
                font-size: 12px;
                color: @gray-1;
                margin-top: 5px;
            }

            .comment-item-right-bottom-reply {
                font-size: 12px;
                display: flex;
                align-items: center;
                color: @gray-1;
                margin-top: 6px;
                cursor: pointer;
                width: fit-content;

                svg {
                    fill: @gray-1;
                }

                span {
                    display: flex;
                    align-items: center;
                    margin-right: 10px;
                    font-size: 12px;
                    line-height: 20px;

                    i {
                        font-size: 14px;
                    }
                }
            }

            .comment-item-right-bottom-list {
                margin-top: 5px;

                .comment-item-right-bottom-list-item {
                    display: flex;
                    padding: 10px 0;

                    .comment-item-right-bottom-list-item-left {
                        width: 20px;
                        height: 20px;
                        border-radius: 50%;
                        overflow: hidden;
                        flex-shrink: 0;

                        img {
                            width: 100%;
                            height: 100%;
                        }
                    }

                    .comment-item-right-bottom-list-item-right {
                        flex: 1;
                        padding-left: 10px;

                        .comment-item-right-bottom-list-item-right-top {

                            .comment-item-right-bottom-list-item-right-top-name {
                                font-size: 0.8rem;
                                font-weight: 700;
                                color: @white;
                                word-break: break-all;
                            }

                            .comment-item-right-bottom-list-item-right-top-time {
                                font-size: 0.6rem;
                                color: @gray-3;
                            }
                        }

                        .comment-item-right-bottom-list-item-right-bottom {
                            margin-top: 5px;

                            .comment-item-right-bottom-list-item-right-bottom-text {
                                font-size: 0.8rem;
                                color: @white;
                            }
                        }
                    }
                }
            }
        }
    }

    .comment-item-more {
        color: @gray-1;
        font-size: 12px;
        cursor: pointer;
    }
}
</style>