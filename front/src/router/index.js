import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '@/layout'

/* Router Modules */

Vue.use(Router)

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
 roles: ['admin','editor']    control the page roles (you can set multiple roles)
 title: 'title'               the name show in sidebar and breadcrumb (recommend set)
 icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
 noCache: true                if set true, the page will no be cached(default is false)
 affix: true                  if set true, the tag will affix in the tags-view
 breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
 activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
 }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    meta: {
      title: '集群信息',
      icon: 'table'
    },
    redirect: '/dashboard',
    hidden: false,
    children: [
      {
        path: '/dashboard',
        component: () => import('@/views/tab/index'),
        name: '集群列表',
        meta: { title: '集群列表', icon: 'dashboard', affix: true }
      },
      {
        path: '/pods',
        hidden: true,
        component: () => import('@/views/table/pods'),
        name: 'Pod列表',
        meta: { title: 'Pod列表', icon: 'list' }
      }
    ]
  },
  {
    path: '/user',
    component: Layout,
    meta: {
      title: '用户管理',
      icon: 'peoples'
    },
    hidden: false,
    children: [
      {
        path: '/managers',
        component: () => import('@/views/table/managers'),
        name: '管理员设置',
        meta: { title: '管理员设置', icon: 'people', affix: true }
      },
      {
        path: '/whites',
        component: () => import('@/views/table/white.vue'),
        name: '白名单管理',
        meta: { title: '白名单管理', icon: 'user' }
      }
    ]
  }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  // const newRouter = createRouter()
  // router.matcher = newRouter.matcher // reset router
}

export default router
