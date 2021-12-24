import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login'
import Admin from '../views/Admin'
import ResList from '../components/res/Reslist'
import Addres from '../components/res/Addres'
import Catelist from '../components/cate/Catelist'
import Addcate from '../components/cate/Addcate'
import Userlist from '../components/user/Userlist'
import Applylist from '../components/borrow/Applylist'
import Borrowlist from '../components/borrow/Borrowlist'
import Applyres from '../components/borrow/Applyres'
import Expired from '../components/borrow/Expired'
import BorrowRecord from '../components/record/BorrowRecord'
import ReturnRecord from '../components/record/ReturnRecord'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: { title: '登录界面' }
  },
  {
    path: '/',
    name: 'admin',
    component: Admin,
    children: [
      {
        path: 'expired',
        component: Expired,
        meta: { title: '过期未归还' }
      },
      {
        path: 'reslist',
        component: ResList,
        props: true,
        meta: { title: '资源列表' }
      },
      {
        path: 'reslist/:id',
        component: ResList,
        props: true,
        meta: { title: '资源列表' }
      },
      {
        path: 'addres',
        meta: { title: '添加资源' },
        component: Addres
      },
      {
        path: 'addres/:id',
        meta: { title: '编辑资源' },
        component: Addres,
        props: true
      },
      {
        path: 'catelist',
        meta: { title: '分类列表' },
        component: Catelist
      },
      {
        path: 'addcate/:id',
        meta: { title: '编辑资源' },
        component: Addcate,
        props: true
      },
      {
        path: 'addcate',
        meta: { title: '添加资源' },
        component: Addcate,
        props: true
      },
      {
        path: 'userlist',
        meta: { title: '用户列表' },
        component: Userlist
      },
      {
        path: 'applylist',
        meta: { title: '已申请列表' },
        component: Applylist
      },
      {
        path: 'borrowlist',
        meta: { title: '已借出列表' },
        component: Borrowlist
      },
      {
        path: 'applyres/:id',
        meta: { title: '申请资源' },
        component: Applyres,
        props: true
      },
      {
        path: 'borrowrecord',
        component: BorrowRecord,
        meta: { title: '借出记录' },
        props: true
      },
      {
        path: 'returnrecord',
        component: ReturnRecord,
        meta: { title: '归还记录' },
        props: true
      }
    ]
  }
]

const router = new VueRouter({
  routes
})
router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})
router.beforeResolve((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path !== '/login') {
    next('/login')
  } else {
    next()
  }
})

export default router
