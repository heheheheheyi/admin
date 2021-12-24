<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search allow-clear v-model="queryParam.catename" placeholder="输入分类名查找" enter-button
                          @search="getCateList"/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="$router.push('/addcate')">新增</a-button>
        </a-col>
      </a-row>
      <a-table bordered rowKey="id" :columns="columns" :pagination="pagination" :dataSource="catelist"
               @change="handleTableChange">
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" style="margin-right: 15px" @click="$router.push(`/addcate/${data.id}`)">编辑
            </a-button>
            <a-button type="danger" @click="deleteCate(data.id)">删除</a-button>
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
    dataIndex: 'catename',
    width: '10%',
    key: 'catename',
    align: 'center'
  },
  {
    title: '数量',
    dataIndex: 'num',
    width: '10%',
    key: 'num',
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
      catelist: [],
      columns,
      queryParam: {
        catename: '',
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created () {
    this.getCateList()
    this.getCateList()
  },
  methods: {
    async getCateList () {
      const { data: res } = await this.$http.get('cate', {
        params: {
          catename: this.queryParam.catename,
          pagesize: this.queryParam.defaultPageSize,
          pagenum: this.queryParam.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.catelist = res.data
      this.pagination.total = res.total
    },
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.pagination = pager
      this.getCateList()
    },
    deleteCate (id) {
      var that = this
      this.$confirm({
        title: '确定要删除该分类吗?',
        content: '删除后无法恢复',
        async onOk () {
          const { data: res } = await that.$http.delete(`cate/${id}`)
          if (res.status !== 200) return that.$message.error(res.message)
          that.$message.success('删除成功')
          that.getCateList()
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
