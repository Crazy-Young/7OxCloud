import request from './request'

export const SendSMS = (data) => request.post('/base/send_sms', data)
export const SendCaptcha = () => request.get('/base/captcha')

export const DeleteFile = (url) => request.delete(`/base/delete_file?url=${url}`)
export const UploadFile = (data, signal) => request.post('/base/upload_file', data, { signal, headers: { 'Content-Type': 'multipart/form-data' } })
