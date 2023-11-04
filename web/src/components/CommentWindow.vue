<template>
    <!-- 设置评论背景 -->
    <div class="comment-window" ref="commentWindow" :style="{ '--bg': `url(${bg})`, }" @wheel.stop @click.stop
        @mousedown.stop @pointerdown.stop>
        <!-- 评论顶部 -->
        <header class="comment-window-header">
            <span>评论</span>
            <span class="close" @click="$emit('close')">
                <CloseIcon></CloseIcon>
            </span>
        </header>
        <!-- 评论内容: 如果没有评论且未加载中，显示没有评论 -->
        <main class="comment-window-main"
            :class="{ 'no-comment': !commentList.length && !loading, 'center': loading && !commentList.length }">
            <!-- 加载图标: 加载中且没有评论时出现 -->
            <Loading v-if="loading && !commentList.length"></Loading>
            <!-- 评论列表 -->
            <CommentItem v-for="item in commentList" :key="item.id" v-bind="item" @replyTo="changeReply"
                @refresh="getCommentList"></CommentItem>
        </main>
        <!-- 评论底部 -->
        <footer class="comment-window-footer">
            <!-- 回复给谁 -->
            <div class="comment-window-footer-reply" v-if="reply.cid">
                <p class="comment-window-footer-reply-content">
                    <span>回复</span>
                    <span>@{{ reply.username }}</span>
                    <span>:</span>
                    <span>{{ reply.content }}</span>
                </p>
                <!-- 关闭图标: 重置回复对象 -->
                <span class="close" @click="reply = {}">
                    <CloseIcon></CloseIcon>
                </span>
            </div>
            <!-- 输入框: 自适应高度，限制300个字符 -->
            <el-input type="textarea" autosize resize="none" v-model="input" maxlength="300" ref="input"
                @keydown.native.enter.prevent="submit" :disabled="!userInfo.uid" @click.native="toggleLogin">
            </el-input>
            <!-- 提交图标 -->
            <div class="comment-window-footer-submit" v-if="input.trim()" @click="submit">
                <SubmitIcon></SubmitIcon>
            </div>
        </footer>
    </div>
</template>

<script>
import CloseIcon from "@/icons/CloseIcon.vue";
import SubmitIcon from "@/icons/SubmitIcon.vue";
import { VideoCommentList, CommentVideo } from "@/api/video";
import flattenData from "@/utils/flattenData";
import CommentItem from "@/components/CommentItem.vue";
import Loading from "./Loading.vue";
import { mapState } from "vuex";
export default {
    name: "CommentWindow",
    components: {
        CloseIcon,
        SubmitIcon,
        CommentItem,
        Loading
    },
    computed: {
        ...mapState({
            userInfo: state => state.user.userInfo
        })
    },
    props: {
        // 评论背景: 视频封面
        bg: {
            type: String,
            default: ""
        },
        // 视频ID
        vid: {
            type: Number | String,
            required: true
        }
    },
    data() {
        return {
            // 评论内容
            input: "",
            // 评论列表
            commentList: [],
            // 数据是否加载中
            loading: false,
            // 回复对象
            reply: {
                cid: null,
                username: "",
                content: ""
            }
        }
    },
    methods: {
        // 切换登录
        toggleLogin() {
            if (!this.userInfo.uid)
                this.$bus.$emit("toggleLogin")
        },
        // 提交评论
        async submit() {
            // 发送评论
            const res = await CommentVideo({
                vid: this.vid,
                content: this.input,
                pid: this.reply.cid
            })
            // 评论失败
            if (!res) {
                return
            }
            // 评论成功
            if (res.status === 200) {
                // 清空输入框
                this.input = ""
                // 清空回复对象
                this.reply = {}
                // 刷新评论
                this.getCommentList()
            }
        },
        changeReply(data) {
            // 回复对象
            this.reply = data
            // 聚焦
            this.focus()
        },
        focus() {
            // 聚焦
            this.$refs.input.focus()
        },
        getCommentList() {
            // 节流
            if (this.loading) return
            this.loading = true
            // 获取评论
            VideoCommentList(this.vid).then(res => {
                // 获取失败
                if (!res) return
                // 获取成功
                if (res.status === 200) {
                    // 评论列表
                    this.commentList = res.data.data || []
                    // 评论列表扁平化
                    this.commentList.forEach(item => {
                        if (item.children)
                            // 每个父评论的子评论都扁平化，并且传入父评论的作者信息
                            item.children = flattenData(item.children, {
                                pid: item.cid,
                                pAuthor: item.author
                            })
                    })
                    // 评论长度
                    this.$emit("changeCommentLength", this.commentList.reduce((a, b) => a + (b.children || []).length + 1, 0))
                }
                // 数据加载完成
                this.loading = false
            })
        }
    },
    mounted() {
        // 获取评论
        this.getCommentList()
    }
}
</script>

<style lang="less" scoped>
.comment-window {
    width: 360px;
    position: relative;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    white-space: nowrap;

    &::after {
        content: '';
        background: var(--bg) no-repeat;
        background-size: cover;
        height: 100%;
        width: 100%;
        position: absolute;
        top: 0;
        right: 0;
        filter: blur(100px) brightness(1.5);
        transform: scale(1.2);
        z-index: -1;
        background-position: center center;
    }

    .comment-window-header {
        height: 48px;
        position: relative;
        display: flex;

        span {
            line-height: 48px;
            height: 100%;
            display: inline-block;
            font-size: 16px;
            font-weight: 500;
            margin-left: 20px;
            color: @white;
            border-bottom: 2px solid @red;
            cursor: pointer;

            &.close {
                margin-left: auto;
                border-bottom: unset;
                margin-right: 10px;
                display: flex;
                align-items: center;

                svg {
                    fill: @white;
                }
            }
        }

    }

    .comment-window-main {
        flex: 1;
        padding: 0 10px 0 16px;
        color: @white;
        overflow-y: scroll;

        &::-webkit-scrollbar {
            display: none;
        }

        &.no-comment {
            display: flex;
            align-items: center;
            justify-content: center;
            background: url(data:image/png;base64,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) no-repeat center center;
            background-size: 200px;

            &::after {
                content: "暂无评论";
                transform: translateY(70px);
                color: @gray-2;
                font-size: 14px;
            }
        }

        &.center {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
        }
    }

    .comment-window-footer {
        margin: 0 16px 16px;
        position: relative;

        .comment-window-footer-reply {
            background-color: @transparent-dark;
            border-radius: 6px;
            display: flex;
            align-items: center;
            position: relative;

            .comment-window-footer-reply-content {
                flex: 1;
                color: @gray-1;
                font-size: 14px;
                word-break: break-all;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                margin-left: 6px;
                line-height: 1;
            }

            .close {
                height: 24px;
                width: 24px;
                margin-left: auto;
                cursor: pointer;

                svg {
                    fill: @white;
                    height: 100%;
                    width: 100%;
                }
            }
        }

        .el-textarea {

            /deep/ .el-textarea__inner {
                padding: 11px 50px 11px 12px;
                border: none;
                background-color: @transparent-darker;
                color: @gray-1;
                font-size: 14px;
                line-height: 22px;

                &::-webkit-scrollbar {
                    display: none;
                }
            }
        }

        .comment-window-footer-submit {
            position: absolute;
            right: 5px;
            bottom: 0;
            cursor: pointer;
        }
    }
}
</style>