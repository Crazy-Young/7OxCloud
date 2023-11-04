export default [
    {
        path: "/",
        component: () => import("@/views/Base"),
        meta: {
            isNavChild: true
        },
        children: [{
            path: "",
            name: "Home",
            component: () => import("@/views/Base/Home"),
            meta: {
                title: "首页",
                isNav: true,
                index: 0,
            }
        }, {
            path: "recommend",
            name: "Recommend",
            component: () => import("@/views/Base/Recommend"),
            meta: {
                title: "推荐",
                isNav: true,
                index: 1,
            }
        }, {
            path: "follow",
            name: "Follow",
            component: () => import("@/views/Base/Follow"),
            meta: {
                title: "关注",
                isNav: true,
                index: 2,
            }
        }, {
            path: "friend",
            name: "Friend",
            component: () => import("@/views/Base/Friend"),
            meta: {
                title: "朋友",
                isNav: true,
                index: 3,
            }
        }, {
            path: "me",
            name: "Me",
            component: () => import("@/views/Base/Me"),
            meta: {
                title: "我的",
                isNav: true,
                index: 15,
            }
        }, {
            path: "upload",
            name: "Upload",
            component: () => import("@/views/Base/Upload"),
            meta: {
                title: "上传",
                isNav: true,
                index: 5,
            }
        },
        {
            path: "hot",
            name: "Hot",
            component: () => import("@/views/Base/Hot"),
            meta: {
                isNav: true,
                index: 6,
                title: "热点"
            }
        },
        {
            path: "channel/:cid",
            name: "Channel",
            component: () => import("@/views/Base/Channel"),
            meta: {
                title: "频道",
                isNavChild: true
            },
        }, {
            path: "user/:uid(\\d+)",
            name: "User",
            component: () => import("@/views/Base/User"),
        }, {
            path: "topic/:tid(\\d+)",
            name: "Topic",
            component: () => import("@/views/Base/Topic")
        }, {
            path: "search",
            name: "Search",
            component: () => import("@/views/Base/Search")
        }]
    },
    {
        path: "/video/:vid(\\d+)",
        name: "Video",
        component: () => import("@/views/Video")
    }
]        