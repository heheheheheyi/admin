<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search allow-clear v-model="queryParam.resname" placeholder="输入资源名查找" enter-button
                          @search="clickSearch"/>
        </a-col>
      </a-row>
      <a-table bordered rowKey="id" :columns="columns" :pagination="pagination" :dataSource="userlist"
               @change="handleTableChange">
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" style="margin-right: 15px" @click="sendMessage(data.resname,data.uid)">提醒</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: '资源名',
    dataIndex: 'resname',
    width: '5%',
    key: 'resname',
    align: 'center'
  },
  {
    title: '分类',
    dataIndex: 'catename',
    width: '5%',
    key: 'catename',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '5%',
    key: 'username',
    align: 'center'
  },
  {
    title: '归还时间',
    dataIndex: 'return_time',
    width: '20%',
    key: 'return_time',
    align: 'center'
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  }
]
export default {
  data () {
    return {
      UserInfo: [],
      Message: {
        uid: 0,
        fid: 0,
        content: '',
        time: ''
      },
      searchText: '',
      onFilterDropdownVisibleChangesearchInput: null,
      searchedColumn: '',
      reslist: [],
      userlist: [],
      columns,
      queryParam: {
        resname: '',
        pagesize: 5,
        pagenum: 1
      },
      pagination: {
        pageSizeOptions: ['2', '5', '10'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: function (total) {
          return `共${total}条`
        }
      }
    }
  },
  async created () {
    this.getResList()
    this.getUserInfo()
  },
  methods: {
    async getResList () {
      const { data: res } = await this.$http.get('res/borrow', {
        params: {
          resname: this.queryParam.resname,
          pagesize: this.queryParam.defaultPageSize,
          pagenum: this.queryParam.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.reslist = res.data
      this.pagination.total = res.total
      this.userlist = []
      this.getname()
    },
    getname () {
      if (this.reslist !== null) {
        this.reslist.forEach((item) => {
          if (new Date().getTime() > new Date(item.return_time).getTime()) {
            var user = {
              id: item.id,
              uid: item.uid,
              resname: item.resname,
              catename: item.CateInfo.catename,
              username: item.UserInfo.username,
              return_time: item.return_time
            }
            this.userlist.push(user)
          }
        })
      }
    },
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.pagination = pager
      this.getResList()
    },
    clickSearch () {
      this.getResList()
    },
    async sendMessage (name, id) {
      this.Message.uid = this.UserInfo.id
      this.Message.fid = id
      this.Message.time = this.getTime()
      this.Message.content = `你借用的${name}未按期归还。`
      const { data: res } = await this.$http.post('chat/add', this.Message)
      if (res.status !== 200) return this.$message.error(res.message)
      this.$message.success('提醒成功')
    },
    getTime () {
      var date = new Date()
      var Y = date.getFullYear() + '-'
      var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-'
      var D = date.getDate() + ' '
      var h = date.getHours() + ':'
      var m = date.getMinutes() + ':'
      var s = date.getSeconds()
      return Y + M + D + h + m + s
    },
    getUserInfo () {
      this.UserInfo = JSON.parse(window.sessionStorage.getItem('UserInfo'))
    }
  }
}
</script>
<style scoped>
</style>
