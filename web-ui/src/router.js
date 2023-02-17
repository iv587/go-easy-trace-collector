import Vue from "vue";
import VueRouter from "vue-router";
import main from "./components/main";
import trace from "./components/trace";
import connection from "./components/connection";

Vue.use(VueRouter);


var routes = [
    {
        path: '/',
        component: main,
        redirect: '/trace',
        children: [
            {
                path: '/trace',
                component: trace,
            },
            {
                path: '/connection',
                component: connection
            }
        ]
    }
]

var router =  new VueRouter({
    routes
})
export default router;