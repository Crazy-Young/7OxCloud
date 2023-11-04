<template>
    <!-- 话题tag -->
    <span class="topic-item" :style="{ '--size': size + 'px' }">
        <!-- 设置跳转 -->
        <router-link :to="{ name: 'Topic', params: { tid: id } }"># {{ description }}</router-link>
    </span>
</template>

<script>
export default {
    name: "Topic",
    props: {
        // 话题id
        id: {
            type: Number | String,
            required: true
        },
        // 话题描述
        description: {
            type: String,
            default: ""
        },
        // 话题tag大小
        size: {
            type: Number,
            default: 12
        }
    }
}
</script>

<style lang="less" scoped>
.topic-item {
    color: @gray-2;
    cursor: pointer;
    background-color: @transparent-dark;
    // 根据size设置字体以及各种边距大小
    font-size: var(--size);
    margin-right: calc(var(--size) * 2 / 3);
    margin-top: calc(var(--size) * 2 / 3);
    padding: calc(var(--size) / 6) calc(var(--size) / 2);
    border-radius: calc(var(--size) / 3);
    box-decoration-break: clone;
}
</style>