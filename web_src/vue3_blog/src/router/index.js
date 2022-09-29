import { createRouter, createWebHistory} from "vue-router";
const router = createRouter({
    history: createWebHistory(),
    routes:[
        {
            path: '/',
            name: 'home',
            component: () => import("../view/Home.vue")
        },
        {
            path:'/quession',
            name: '问答',
            component: ()=> import("../view/Quesion.vue")
        },
        {
            path:'/square',
            name: '广场',
            component: ()=> import("../view/Quesion.vue")
        },
        {
            path:'/bbs',
            name: '论坛',
            component: ()=> import("../view/Quesion.vue")
        }
    ]
});

export default router;