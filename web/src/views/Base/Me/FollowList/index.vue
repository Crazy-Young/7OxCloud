<template>
    <div id="follow-list">
        <!-- 遮罩层 -->
        <div class="mask" @click="$emit('close')"></div>
        <!-- 面板 -->
        <div class="panel">
            <!-- 关闭按钮 -->
            <div class="panel-close" @click="$emit('close')"></div>
            <!-- 头部: 切换关注/粉丝 -->
            <div class="panel__header">
                <div class="panel__header-title" @click="activeIndex = 0" :class="{ active: activeIndex === 0 }">关注({{
                    m_followCount }})</div>
                <div class="panel__header-title" @click="activeIndex = 1" :class="{ active: activeIndex === 1 }">粉丝({{
                    m_fanCount }})</div>
            </div>

            <!-- 搜索框 -->
            <div class="panel__search">
                <el-input clearable v-model="search" placeholder="搜索" @input="handleSearch">
                    <template #prefix>
                        <search-icon />
                    </template>
                </el-input>
            </div>

            <!-- 列表 -->
            <div class="panel-content">
                <div class="list-item" v-for="item in activeIndex === 0 ? followList : fanList" :key="item.id || item.uid">
                    <router-link :to="`/user/${item.id || item.uid}`"><el-avatar :size="60" :src="item.avatar"></el-avatar></router-link>
                    <div class="info">
                        <p class="name">
                            <router-link :to="`/user/${item.id || item.uid}`">{{ item.username }}</router-link>
                        </p>
                        <p class="desc">{{ item.signature }}</p>
                    </div>
                    <el-button class="btn" :class="{ active: item.isFollow }" @click="handleFollow(item)">
                        {{ item.isFollow && item.isFan ? '互相关注' : item.isFollow ? '已关注' : item.isFan ? '回关' : '关注' }}
                    </el-button>
                    <!-- 如果是粉丝，那么可以移除 -->
                    <el-button class="btn active" v-if="activeIndex === 1" @click="handleRemoveFan(item)">移除</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import SearchIcon from "@/icons/SearchIcon.vue";
import { FollowUser, RemoveFan, SearchFollow, SearchFan } from "@/api/user";
import { mapState, mapGetters } from "vuex";
import debounce from "@/utils/debounce";
export default {
    name: "FollowList",
    components: {
        SearchIcon
    },
    computed: {
        ...mapState({
            followList: (state) => state.user.followList,
            fanList: (state) => state.user.fanList
        }),
        ...mapGetters({
            followCount: "followCount",
            fanCount: "fanCount"
        })
    },
    data() {
        return {
            activeIndex: 0,
            search: "",
            m_followCount: 0,
            m_fanCount: 0
        };
    },
    created() {
        this.m_followCount = this.followCount
        this.m_fanCount = this.fanCount
    },
    methods: {
        handleFollow(item) {
            FollowUser({
                userId: item.id || item.uid,
                type: !item.isFollow ? 1 : 2
            }).then((res) => {
                if (res.status === 200) {
                    item.isFollow = !item.isFollow
                    this.m_followCount += item.isFollow ? 1 : -1
                    this.handleSearch()
                } else {
                    this.$message.error(res.data.msg)
                }
            })
        },
        handleRemoveFan(item) {
            RemoveFan(item.id || item.uid).then((res) => {
                if (res.status === 200) {
                    this.m_fanCount -= 1
                    this.handleSearch()
                } else {
                    this.$message.error(res.data.msg)
                }
            })
        },
        handleSearch: debounce(function () {
            if (this.search.trim()) {
                if (this.activeIndex === 0) {
                    SearchFollow(this.search).then((res) => {
                        if (res.status === 200) {
                            this.$store.commit('SET_USER_FOLLOW', res.data.data)
                        } else {
                            this.$message.error(res.data.msg)
                        }
                    })
                } else {
                    SearchFan(this.search).then((res) => {
                        if (res.status === 200) {
                            this.$store.commit('SET_USER_FAN', res.data.data)
                        } else {
                            this.$message.error(res.data.msg)
                        }
                    })
                }
            } else {
                this.$store.dispatch('UserFollowList')
                this.$store.dispatch('UserFanList')
            }
        })
    }
};
</script>

<style lang="less" scoped>
#follow-list {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 999;

    .mask {
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        background-color: @mask;
    }

    .panel {
        position: absolute;
        top: 50%;
        left: 50%;
        display: flex;
        border-radius: 4px;
        height: 648px;
        box-sizing: content-box;
        padding: 36px 40px;
        flex-direction: column;
        transform: translate(-50%, -50%);
        width: 560px;
        background: @gray-4;
        color: @white;

        .panel-close {
            background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJAAAACQCAYAAADnRuK4AAAACXBIWXMAACxLAAAsSwGlPZapAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAPTSURBVHgB7d29ThRRGIDh7wyLWGlFQoJ7A4rBBgs7CxPs8QqwsRMra3u8Am/AYGOnnR2yGxNZxM4KsgnQURrY43y7IQgL7PzPnHPeJzFCmGLIvMz/OSsCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoCxGPLGwsNSemjJrVuxzY+SOtWZPrN0UidZ7va19aQBdxygyqyaSZRHbtlaOxcqXJq1jWl4E9PDR4xfG2ncazuWfaUgnYld/b3d3pUaLi0sPrMinq9ZRDcS83vnZ2RDHOB/QcM/Tku83LaN/6SciK3VFNCkeNdobRc9c2xNF4rgokjeTltEN14o34P14Q0rFksSj9OfG2Im/S9M4H1C8D11OtFgNESWN54yev4ljnA8o6cY5W7aqiNLGo9Is2xTu74FEUp0zVBFRlnhG4itHx7gf0PAyOJ0yI8oez/AQtimOcT6g01vyYXgFk1IZEeWJZ7j3mYnWxTHOB/Sr292L70Ws1B1RrnhMvO7TdrW35d7NRG/uROc6dOS8T5Q7npasbHfrvdGZlTcBqToiCjke5VVAqsqIQo9HeReQqiIi4hnxMiBVZkTEc87bgFQZERHPRV4HpIqMiHjGeR+QKiKiaf2aeMYEEZDKG5H+TzzjgglI5XvUkIHn8aigAlKVRRRAPCq4gFTpEQUSjwoyIFVaRAHFo4INSBUeUWDxqKADUoVFFGA8KviAVO6IAo1HtQS52YHIyV8JEoewgg5hdQ9erAsn0QWeRIcYEZfxBV/GhxYRNxJLEFJEPMooSSgR8TC1RCFE5MPQ5kTyvgzWlMGLTRNEQEW8SdiEwYtNxCutN7l0h7nOwYtNxUv117nm8QQRXcSwnqtMeLZFROcYWHhZwgejRDTC0Ob/pXyqTkRMrnAu4ysZoUfE9C4q5/s8IUfkRUBNGDEaakReTDQeTQ033D1Jq+A3CXMPXmSi8erpRONNiEdtx3uQPHesJbIvxTE+TDT+RNIq8R3mXBHZZJOmN4kPE42n2/tU8AJ85oiMbYtjnA8o1UaqcPREloiy7LXq5sMhLNlE4zUMvUkdkUk/aXrdnA9oMC3vJ26gGsdtpYrIwYnGp8Rxh/3+8ezc/FF03af2NGDQ38FB/2hubv5bvC5P42/HL/HjdTy18nbnR8e5jzpwPiB1eNDfnW3Pb5iBfuaW6InoTPxv34r5GP9Vr/U6nT9SM40oXsevw3UUuSsakr7pOJDPcjt65WI8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJX+ATc9oF5n57RZAAAAAElFTkSuQmCC);
            background-size: 100% 100%;
            cursor: pointer;
            height: 36px;
            position: absolute;
            right: 10px;
            top: 10px;
            width: 36px;
            z-index: 9999;
            filter: brightness(4);
        }

        .panel__header {
            background: transparent;
            border-radius: 4px 4px 0 0;
            width: 100%;
            display: flex;
            position: relative;

            &::after {
                content: "";
                position: absolute;
                height: 1px;
                transform: scaleY(0.2);
                background-color: @gray-2;
                left: 0;
                right: 0;
                bottom: 0;
            }

            .panel__header-title {
                font-size: 18px;
                font-weight: 500;
                line-height: 34px;
                cursor: pointer;
                line-height: 26px;
                margin-bottom: 22px;
                margin-right: 40px;
                margin-top: -10px;
                color: @gray-2;

                &:hover {
                    color: @gray-1;
                }

                &.active {
                    color: @white;
                }
            }

        }

        .panel__search {
            margin: 12px 0;

            .el-input {
                height: 32px;
                line-height: 32px;

                /deep/ .el-input__inner {
                    height: 32px;
                    line-height: 32px;
                    width: 100%;
                    background-color: @transparent-dark;
                    border: unset;
                    padding-left: 36px;
                    color: @gray-1;
                }

                /deep/ .el-input__prefix {
                    display: flex;
                    align-items: center;
                    padding-left: 4px;
                }

                /deep/ .el-input__suffix {
                    display: flex;
                    align-items: center;
                }
            }
        }

        .panel-content {
            flex: 1;
            border-radius: 12px;
            min-height: 320px;
            display: flex;
            flex-direction: column;
            overflow-y: auto;
            padding-right: 10px;

            &::-webkit-scrollbar {
                display: none;
            }

            .list-item {
                width: 100%;
                display: flex;
                align-items: center;
                margin: 17px 0;

                .info {
                    flex: 1;
                    display: flex;
                    flex-direction: column;
                    margin: 0 20px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;

                    .name {
                        margin: 4px 0;
                        font-size: 16px;
                        font-weight: 500;
                        line-height: 24px;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                        color: @gray-1;
                    }

                    .desc {
                        margin: 4px 0;
                        font-size: 12px;
                        font-weight: 400;
                        line-height: 20px;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        color: @gray-1;
                        white-space: nowrap;
                    }
                }

                .btn {
                    width: 72px;
                    height: 32px;
                    font-size: 14px;
                    border-radius: 4px;
                    line-height: 32px;
                    background-color: @red;
                    border: unset;
                    color: @white;
                    padding: 0;

                    &:hover {
                        opacity: 0.8;
                    }

                    &.active {
                        background-color: @transparent-dark;
                    }
                }
            }
        }
    }


}
</style>