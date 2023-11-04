import request from './request'

/**
 * 发送短信
 * @param {Object} data 
 * @returns {Promise}
 * @example 注册
 * SendSMS({
 *     mobile: '13900000001',
 *     type: 1
 * })
 * @example 登录
 * SendSMS({
 *     mobile: '13900000001',
 *     type: 2
 * })
 * @example 修改密码
 * SendSMS({
 *     mobile: '13900000001',
 *     type: 3
 * })
 */
export const SendSMS = (data) => request.post('/base/send_sms', data)

/**
 * 发送图形验证码
 * @returns {Promise}
 */
export const SendCaptcha = () => request.get('/base/captcha')

/**
 * 删除文件
 * @param {String} url 
 * @returns {Promise}
 */
export const DeleteFile = (url) => request.delete(`/base/delete_file?url=${url}`)

/**
 * 上传文件
 * @param {FormData} data 
 * @returns {Promise}
 */
export const UploadFile = (data, signal) => request.post('/base/upload_file', data, { signal, headers: { 'Content-Type': 'multipart/form-data' } })
