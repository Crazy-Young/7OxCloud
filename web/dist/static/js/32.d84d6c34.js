"use strict";(self["webpackChunk_7OxCloud"]=self["webpackChunk_7OxCloud"]||[]).push([[32],{4366:function(t,e,i){i.d(e,{Z:function(){return u}});var s=function(){var t=this,e=t._self._c;return e("span",{staticClass:"topic-item",style:{"--size":t.size+"px"}},[e("router-link",{attrs:{to:{name:"Topic",params:{tid:t.id}}}},[t._v("# "+t._s(t.description))])],1)},o=[],a={name:"Topic",props:{id:{type:Number|String,required:!0},description:{type:String,default:""},size:{type:Number,default:12}}},r=a,n=i(1001),l=(0,n.Z)(r,s,o,!1,null,"7b01258c",null),u=l.exports},2176:function(t,e,i){i.r(e),i.d(e,{default:function(){return f}});var s=function(){var t=this,e=t._self._c;return e("div",{staticClass:"video",on:{dblclick:function(e){return t.$router.go(-1)}}},[t.video.playUrl?e("PlayerWithMethod",t._b({attrs:{height:"100%",isUser:"",autoplay:""},scopedSlots:t._u([{key:"title",fn:function(){return[e("div",{staticClass:"xgplayer-title"},[e("span",{staticClass:"xgplayer-title-location"},[e("i",{staticClass:"el-icon-location-outline"}),e("strong",[t._v(t._s(t.video.author.location))])]),e("p",{staticClass:"xgplayer-title-user"},[e("span",{staticClass:"username"},[e("router-link",{attrs:{to:`/user/${t.video.author.uid}`}},[t._v(" @"+t._s(t.video.author.username)+" ")])],1),e("span",{staticClass:"and"},[t._v("·")]),e("span",{staticClass:"time"},[t._v(t._s(t.formatTime(t.video.publishTime)))])]),e("h3",{staticClass:"xgplayer-title-text"},[t._v(t._s(t.video.description))]),e("p",{staticClass:"xgplayer-title-topics"},t._l(t.video.topics,(function(i){return e("Topic",t._b({key:i.id},"Topic",i,!1))})),1)])]},proxy:!0}],null,!1,1396715989)},"PlayerWithMethod",t.video,!1),[[e("div",{staticClass:"back",on:{click:function(e){return t.$router.go(-1)}}},[e("CloseIcon")],1)]],2):e("Loading")],1)},o=[],a=i(2675),r=i(6655),n=i(4781),l=i(4366),u=i(5980),c=i(9884),d={name:"Video",data(){return{video:{}}},components:{CloseIcon:r.Z,Topic:l.Z,PlayerWithMethod:u.Z,Loading:c.Z},computed:{vid(){return this.$route.params.vid}},methods:{formatTime:n.Z,getVideoInfo(){(0,a._7)(this.vid).then((t=>{t&&(200===t.status?this.video=t.data.data:this.$message.error(t.data.msg))}))}},mounted(){this.getVideoInfo()}},p=d,v=i(1001),h=(0,v.Z)(p,s,o,!1,null,"317d4330",null),f=h.exports}}]);