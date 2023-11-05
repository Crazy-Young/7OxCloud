export default {
    data() {
        return {
            videoList: [],
            latestTime: null,
            isDataEnd: false,
            isRequestData: false
        }
    },
    methods: {
        async reqVideoList(getVideoFunction) {
            try {
                if(this.isRequestData || this.isDataEnd) {
                    return
                }
                this.isRequestData = true
                const res = await getVideoFunction()
                if(!res) {
                    return
                }
                if (res.status !== 200) {
                    this.$message.error(res.data.msg)
                    return
                }
                
                this.videoList.push(...(res.data.data.videos || []))
                this.latestTime = res.data.data.nextTime

                if (res.data.data.videos.length < 15 || !res.data.data.nextTime) {
                    this.isDataEnd = true
                }
                this.isRequestData = false
            } catch (error) {
                this.$message.error(error)
                this.isDataEnd = true
                this.isRequestData = false
            }
        }
    }
}