<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search allow-clear v-model="queryParam.username" placeholder="输入用户名查找" enter-button
                          @search="getUserList"/>
        </a-col>
      </a-row>
      <a-table bordered rowKey="id" :columns="columns" :pagination="pagination" :dataSource="userlist"
               @change="handleTableChange">
        <span slot="role" slot-scope="role">{{role==1? '管理员':'普通用户'}}</span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="danger" @click="deleteUser(data.id)">删除</a-button>
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
    dataIndex: 'username',
    width: '10%',
    key: 'username',
    align: 'center'
  },
  {
    title: '借出数量',
    dataIndex: 'num',
    width: '10%',
    key: 'num',
    align: 'center'
  },
  {
    title: '权限',
    dataIndex: 'role',
    width: '10%',
    key: 'role',
    scopedSlots: { customRender: 'role' },
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
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created () {
    this.getUserList()
    this.getUserList()
  },
  methods: {
    async getUserList () {
      const { data: res } = await this.$http.get('user', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.defaultPageSize,
          pagenum: this.queryParam.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data
      this.userlist.num = 1
      this.pagination.total = res.total
    },
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.pagination = pager
      this.getUserList()
    },
    deleteUser (id) {
      var that = this
      this.$confirm({
        title: '确定要删除该分类吗?',
        content: '删除后无法恢复',
        async onOk () {
          const { data: res } = await that.$http.delete(`user/${id}`)
          if (res.status !== 200) return that.$message.error(res.message)
          that.$message.success('删除成功')
          that.getUserList()
        },
        onCancel () {
          that.$message.info('已取消删除')
        }
      })
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}

</style>
