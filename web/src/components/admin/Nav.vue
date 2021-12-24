<template>
  <a-layout-sider breakpoint="lg">
    <div class="log">
      <span>{{ collapsed ? '管理系统' : '实验室管理系统' }}</span>
    </div>
    <a-menu theme="dark" mode="inline" @click="gotopage">
      <a-menu-item key="reslist">
        <a-icon type="dash"/>
        <span>资源管理</span></a-menu-item>
      <a-menu-item key="catelist">
        <a-icon type="small-dash"/>
        <span>分类管理</span></a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user"></a-icon>
        <span>用户列表</span></a-menu-item>
      <a-menu-item key="applylist">
        <a-icon type="exclamation"/>
        <a-badge :count="this.total" dot>
          <span>待确认借出</span>
        </a-badge>
      </a-menu-item>
      <a-menu-item key="borrowlist">
        <a-icon type="info"/>
        <span>已借出资源</span></a-menu-item>
      <a-menu-item key="Expired">
        <a-icon type="info"/>
        <span>过期未归还</span></a-menu-item>
      <a-menu-item key="borrowrecord">
        <a-icon type="info"/>
        <span>借出记录</span></a-menu-item>
      <a-menu-item key="returnrecord">
        <a-icon type="info"/>
        <span>归还记录</span></a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>
<script>
export default {
  data () {
    return {
      collapsed: false,
      total: 0,
      borrowlist: []
    }
  },
  created () {
    this.getCategoryList()
  },
  methods: {
    gotopage (item) {
      this.$router.push('/' +
        item.key
      ).catch((err) => err)
    },
    async getApplyBorrowList () {
      const { data: res } = await this.$http.get('res/applyborrow')
      if (res.status !== 200) return this.$message.error(res.message)
      this.total = res.total
    },
    async getCategoryList () {
      const { data: res } = await this.$http.get('cate')
      if (res.status !== 200) return this.$message.error(res.message)
      this.Categorylist = res.data
      this.getApplyBorrowList()
    }
  }
}
</script>
<style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #1c1c1c;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 15px;
  color: #ffffff;
}
</style>
