<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search allow-clear v-model="queryParam.borrowname" placeholder="输入资源名查找" enter-button
                          @search="getBorrowList"/>
        </a-col>
      </a-row>
      <a-table bordered rowKey="id" :columns="columns" :pagination="pagination" :dataSource="borrowlist"
               @change="handleTableChange">
        <span class="ResImg" slot="img" slot-scope="img">
          <img :src="img">
        </span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="danger" style="margin-right: 15px" @click="returnBorrow(data.id,data.resname,data.uid)">归还
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>
<script>

const columns = [
  {
    title: '名字',
    dataIndex: 'resname',
    width: '5%',
    key: 'resname',
    align: 'center'
  },
  {
    title: '分类',
    dataIndex: 'CateInfo.catename',
    width: '5%',
    key: 'CateInfo.catename',
    align: 'center'
  },
  {
    title: '图片',
    dataIndex: 'img',
    width: '10%',
    key: 'img',
    scopedSlots: { customRender: 'img' },
    align: 'center'
  },
  {
    title: '借出用户名',
    dataIndex: 'UserInfo.username',
    width: '10%',
    key: 'UserInfo.borrowname',
    align: 'center'
  },
  {
    title: '借出时间',
    dataIndex: 'borrow_time',
    width: '10%',
    key: 'borrow_time',
    align: 'center'
  },
  {
    title: '归还日期  ',
    dataIndex: 'return_time',
    width: '10%',
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
      pagination: {
        pageSizeOptions: ['2', '5', '10'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: function (total) {
          return `共${total}条`
        }
      },
      borrowlist: [],
      columns,
      queryParam: {
        borrowname: '',
        pagesize: 5,
        pagenum: 1
      },
      Message: {
        uid: 0,
        fid: 0,
        content: '',
        time: ''
      }
    }
  },
  created () {
    this.getBorrowList()
    this.getUserInfo()
  },
  methods: {
    getUserInfo () {
      this.UserInfo = JSON.parse(window.sessionStorage.getItem('UserInfo'))
    },
    async getBorrowList () {
      const { data: res } = await this.$http.get('res/borrow', {
        params: {
          resname: this.queryParam.resname,
          pagesize: this.queryParam.defaultPageSize,
          pagenum: this.queryParam.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.borrowlist = res.data
      this.pagination.total = res.total
    },
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.pagination = pager
      this.getBorrowList()
    },
    returnBorrow (id, name, uid) {
      var that = this
      this.$confirm({
        title: '确定该资源已经归还了吗?',
        content: '请仔细审核',
        async onOk () {
          const res = await that.$http.delete(`res/borrow/${id}/${that.UserInfo.id}`)
          if (res.status !== 200) return that.$message.error(res.message)
          that.$message.success('归还成功')
          that.sendMessage(name, uid)
          that.getBorrowList()
        },
        onCancel () {
          that.$message.info('已取消归还')
        }
      })
    },
    async sendMessage (name, id) {
      this.Message.uid = this.UserInfo.id
      this.Message.fid = id
      this.Message.time = this.getTime()
      this.Message.content = `你借用的${name}归还成功。`
      const { data: res } = await this.$http.post('chat/add', this.Message)
      if (res.status !== 200) return this.$message.error(res.message)
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
    }
  }
}

</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}

.ResImg {
  height: 100%;
  width: 100%;
}

.ResImg img {
  width: 150px;
  height: 120px;
}

</style>
