<template>
  <div>
    <a-card>
      <a-table bordered rowKey="id" :columns="columns" :pagination="pagination" :dataSource="recordlist"
               @change="handleTableChange">
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: '资源名',
    dataIndex: 'ResInfo.resname',
    width: '5%',
    key: 'ResInfo.resname',
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
    title: '用户',
    dataIndex: 'UserInfo.username',
    width: '5%',
    key: 'UserInfo.username',
    align: 'center'
  },
  {
    title: '操作员',
    dataIndex: 'ManagerInfo.username',
    width: '5%',
    key: 'ManagerInfo.username',
    align: 'center'
  },
  {
    title: 'time',
    dataIndex: 'time',
    width: '20%',
    key: 'time',
    align: 'center'
  }
]
export default {
  data () {
    return {
      recordlist: [],
      columns,
      queryParam: {
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
    this.GetBorrowRecord()
  },
  methods: {
    async GetBorrowRecord () {
      const { data: res } = await this.$http.get('record/borrow', {
        params: {
          pagesize: this.queryParam.defaultPageSize,
          pagenum: this.queryParam.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.recordlist = res.data
      this.pagination.total = res.total
    },
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.pagination = pager
      this.GetBorrowRecord()
    }
  }
}
</script>
<style scoped>
</style>
