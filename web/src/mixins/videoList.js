export default {
    data() {
        return {
            videoList: [],
            latestTime: null,
            isDataEnd: false
        }
    },
    methods: {
        async reqVideoList(getVideoFunction) {
            try {
                if(this.isDataEnd) {
                    return
                }
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

                if (res.data.data.videos.length < 15) {
                    this.isDataEnd = true
                }
            } catch (error) {
                this.$message.error(error)
                this.isDataEnd = true
            }
        }
    }
}