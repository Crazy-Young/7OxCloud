import request from './request'

export const VideoCategoryList = () => request.get('/video/category_list')
export const VideoTopicList = () => request.get('/video/topic_list')
export const VideoPublishList = (uid) => request.get('/video/publish_list' + (uid ? `?uid=${uid}` : ''))
export const VideoPublish = (data) => request.post('/video/publish', data)
export const VideoInfo = (vid) => request.get('/video/get_video' + (vid ? `?vid=${vid}` : ''))
export const VideoListByTime = (last) => request.get('/video/feed' + (last ? `?latestTime=${last}` : ''))
export const VideoListByCategory = (categoryId, last) => request.get('/video/feed_by_category' + (categoryId ? `?categoryId=${categoryId}&latestTime=${last}` : ''))
export const VideoListByTopic = (tid, last) => request.get('/video/feed_by_topic' + (tid ? `?tid=${tid}&latestTime=${last}` : ''))