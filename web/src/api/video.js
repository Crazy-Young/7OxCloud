import request from './request'
import axios from 'axios'

/**
 * 获取视频分类
 * @returns {Promise}
 * @memberof Video
 */
export const VideoCategoryList = () => request.get('/video/category_list')

/**
 * 获取视频话题
 * @returns {Promise}
 * @memberof Video
 */
export const VideoTopicList = () => request.get('/video/topic_list')

/**
 * 获取视频列表
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoList = (last) => request.get('/video/feed' + (last ? `?latestTime=${last}` : ''))

/**
 * 视频发布
 * @param {Object} data
 * @returns {Promise}
 * @memberof Video
 * @example
 * VideoPublish({
     categoryId: 1,
     coverUrl: 'http://s38ddu7np.hn-bkt.clouddn.com/background.png',
     playUrl: 'http://s38ddu7np.hn-bkt.clouddn.com/15d8a652-aadc-449b-b4f9-004f7fecdd2a.mp4',
     description: '这是一个视频',
     topics: ['话题1', '话题2'],
 })
 */
export const VideoPublish = (data) => request.post('/video/publish', data)

/**
 * 获取视频详情
 * @param {String} vid 视频ID
 * @returns {Promise}
 * @memberof Video
 */
export const VideoInfo = (vid) => request.get('/video/get_video' + (vid ? `?vid=${vid}` : ''))

/**
 * 删除视频
 * @param {String} vid 视频ID
 * @returns {Promise}
 * @memberof Video
 */
export const VideoDelete = (vid) => request.delete('/video/delete_video' + (vid ? `?vid=${vid}` : ''))

/**
 * 根据发布时间获取视频
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByTime = (last) => request.get('/video/feed' + (last ? `?latestTime=${last}` : ''))

/**
 * 根据分类获取视频
 * @param {Number} categoryId 分类ID
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByCategory = (categoryId, last) => request.get('/video/feed_by_category' + (categoryId ? `?categoryId=${categoryId}&latestTime=${last}` : ''))

/**
 * 根据话题获取视频
 * @param {Number} tid 话题ID
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByTopic = (tid, last) => request.get('/video/feed_by_topic' + (tid ? `?tid=${tid}&latestTime=${last}` : ''))

/**
 * 根据关键字获取视频
 * @param {String} keyword 关键字
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListBySearch = (keyword, last) => request.get('/video/feed_by_search' + (keyword ? `?keyword=${keyword}&latestTime=${last}` : ''))

/**
 * 根据热门获取视频
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByHot = (last) => request.get('/video/hot_feed' + (last ? `?latestTime=${last}` : ''))

/**
 * 根据关注用户获取视频
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByFollow = (last) => request.get('/social/follow_feed' + (last ? `?latestTime=${last}` : ''))

/**
 * 根据推荐用户获取视频
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByRecommend = (last) => request.get('/social/feed_by_recommend' + (last ? `?latestTime=${last}` : ''))

/**
 * 根据好友获取视频
 * @param {Number} last 最新时间
 * @returns {Promise}
 * @memberof Video
 */
export const VideoListByFriend = (last) => request.get('/social/friend_feed' + (last ? `?latestTime=${last}` : ''))

// 设置取消作品，喜欢，收藏，观看历史板块的令牌
let UserVideoCancelToken = axios.CancelToken.source()

/**
 * 用户切换tab时，取消原来的视频请求
 * @memberof Video
 */
export const cancelUserVideo = () => {
    // 取消原来的请求
    UserVideoCancelToken.cancel()
    // 创建新的请求，否则无法再次发起
    UserVideoCancelToken = axios.CancelToken.source()
}

/**
 * 获取用户视频列表
 * @param {Number} uid 用户ID
 * @returns {Promise}
 * @memberof Video
 */
export const VideoPublishList = (uid) => request.get('/video/publish_list' + (uid ? `?uid=${uid}` : ''), {
    cancelToken: UserVideoCancelToken.token
})

/**
 * 点赞视频列表
 * @param {Number} uid 用户ID
 * @returns {Promise}
 * @memberof Video
 */
export const LikeVideoList = (uid) => request.get('/interaction/like_list' + (uid ? `?uid=${uid}` : ''), {
    cancelToken: UserVideoCancelToken.token
})

/**
 * 收藏视频列表
 * @param {Number} uid 用户ID
 * @returns {Promise}
 * @memberof Video
 */
export const CollectVideoList = (uid) => request.get('/interaction/collect_list' + (uid ? `?uid=${uid}` : ''), {
    cancelToken: UserVideoCancelToken.token
})



/**
 * 喜欢/取消喜欢视频
 * @param {Object} data
 * @returns {Promise}
 * @memberof Video
 * @example 喜欢
 * LikeVideo({
     type: 1,
     vid: 1
 })
 * @example 取消喜欢
 * LikeVideo({
     type: 2,
     vid: 1
 })
 */
export const LikeVideo = (data) => request.post('/interaction/like_video', data)

/**
 * 收藏/取消收藏视频
 * @param {Object} data
 * @returns {Promise}
 * @memberof Video
 * @example 收藏
 * CollectVideo({
     type: 1,
     vid: 1
 })
 * @example 取消收藏
 * CollectVideo({
     type: 2,
     vid: 1
 })
 */
export const CollectVideo = (data) => request.post('/interaction/collect_video', data)

/**
 * 评论视频
 * @param {Object} data
 * @returns {Promise}
 * @memberof Video
 * @example
 * CommentVideo({
     content: '这是一条评论',
     vid: '1',
     pid: 1
 })
 */
export const CommentVideo = (data) => request.post('/interaction/comment', data)

/**
 * 记录用户浏览视频
 * @param {Number} vid 视频ID
 * @returns {Promise}
 * @memberof Video
 */
export const ViewVideo = (vid) => request.get('/interaction/view_video' + (vid ? `?vid=${vid}` : ''))

/**
 * 删除评论
 * @param {Number} cid 评论ID
 * @returns {Promise}
 * @memberof Video
 */
export const DeleteComment = (cid) => request.delete('/interaction/delete_comment' + (cid ? `?cid=${cid}` : ''))

/**
 * 获取评论列表
 * @param {Number} vid 视频ID
 * @returns {Promise}
 * @memberof Video
 */
export const VideoCommentList = (vid) => request.get('/interaction/comment_list' + (vid ? `?vid=${vid}` : ''))