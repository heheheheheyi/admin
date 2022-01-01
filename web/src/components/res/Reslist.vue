<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search allow-clear v-model="queryParam.resname" placeholder="输入资源名查找" enter-button
                          @search="clickSearch"/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="$router.push('/addres')">新增</a-button>
        </a-col>
        <a-col :span="6" :offset="0">
          <a-select :allowClear=true defaultValue="请选择分类" style="width: 200px" @change="CateChange">
            <a-select-option v-for="item in this.Categorylist" :key="item.id" :value="item.id">
              {{ item.catename }}
            </a-select-option>
          </a-select>
        </a-col>
        <a-col :span="4" :offset="0">
          <a-select :allowClear=true defaultValue="请选择状态" style="width: 200px" @change="statusChange">
            <a-select-option v-for="item in this.Statuslist" :key="item.status" :value="item.status">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-col>
      </a-row>
      <a-table bordered rowKey="id" :columns="columns" :pagination="pagination" :dataSource="reslist"
               @change="handleTableChange">
        <span slot="status" slot-scope="status">{{ checkStatus(status) }}</span>
        <span class="ResImg" slot="img" slot-scope="img">
          <img :src="img">
        </span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" style="margin-right: 15px" @click="showModal(data.id)">
              详情
            </a-button>
            <a-modal
              title="详情"
              :visible="visible"
              @ok="handleOk"
              @cancel="handleCancel"
              :mask="false"
              :footer="null"
            >
              <a-descriptions>
                <a-descriptions-item label="ID">
                  {{ Detailslist.id }}
                </a-descriptions-item>
                <br/>
                <br/>
                <a-descriptions-item label="资源名">
                  {{ Detailslist.resname }}
                </a-descriptions-item>
                <br/>
                <br/>
                <a-descriptions-item label="分类名">
                  {{ Detailslist.catename }}
                </a-descriptions-item>
                <br/>
                <br/>
                <a-descriptions-item label="图片">
                  <img :src="Detailslist.img" width="50" height="50">
                </a-descriptions-item>
                <br/>
                <br/>
                <a-descriptions-item label="状态">
                  {{ checkStatus(Detailslist.status) }}
                </a-descriptions-item>
                <br/>
                <br/>
                <a-descriptions-item v-if="Detailslist.status==3" label="借出人">
                  {{ Detailslist.username }}
                </a-descriptions-item>
                <br/>
                <br/>
                <a-descriptions-item v-if="Detailslist.status==3" label="借出时间">
                  {{ Detailslist.borrow_time }}
                </a-descriptions-item>
                <br/>
                <a-descriptions-item v-if="Detailslist.status==3" label="归还时间">
                  {{ Detailslist.return_time }}
                </a-descriptions-item>
              </a-descriptions>
            </a-modal>
            <a-button type="primary" style="margin-right: 15px" @click="$router.push(`/applyres/${data.id}`)">申请资源
            </a-button>
            <a-button type="primary" style="margin-right: 15px" @click="$router.push(`/addres/${data.id}`)">编辑
            </a-button>
            <a-button type="danger" @click="deleteRes(data.id)">删除</a-button>
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
    title: '状态',
    dataIndex: 'status',
    width: '5%',
    key: 'status',
    scopedSlots: { customRender: 'status' },
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
  props: ['id'],
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
      Statuslist: [
        {
          status: 1,
          name: '未借出'
        },
        {
          status: 2,
          name: '待确认'
        },
        {
          status: 3,
          name: '已借出'
        }
      ],
      Categorylist: [],
      cateId: 0,
      status: 0,
      reslist: [],
      columns,
      queryParam: {
        resname: '',
        pagesize: 5,
        pagenum: 1
      },
      visible: false,
      confirmLoading: false,
      Detailslist: []
    }
  },
  created () {
    this.getCateList()
    this.getResList()
    this.checkTime()
  },
  methods: {
    checkTime () {
      if (new Date().getTime() > new Date(this.hahatime).getTime()) {

      }
    },
    async getResList () {
      if (this.cateId > 0) {
        const { data: res } = await this.$http.get(`cateres/${this.cateId}`, {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      } else {
        const { data: res } = await this.$http.get('res', {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      }
    },
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.pagination = pager
      this.getResList()
    },
    deleteRes (id) {
      var that = this
      this.$confirm({
        title: '确定要删除该资源吗?',
        content: '删除后无法恢复',
        async onOk () {
          const { data: res } = await that.$http.delete(`res/${id}`)
          if (res.status !== 200) return that.$message.error(res.message)
          that.$message.success('删除成功')
          that.getResList()
        },
        onCancel () {
          that.$message.info('已取消删除')
        }
      })
    },
    async getCateList () {
      const { data: res } = await this.$http.get('cate')
      if (res.status !== 200) return this.$message.error(res.message)
      this.Categorylist = res.data
    },
    checkStatus (status) {
      if (status === 1) {
        return '未借出'
      }
      if (status === 2) {
        return '待确认'
      }
      if (status === 3) {
        return '已借出'
      }
    },
    async getBorrowList () {
      if (this.cateId > 0) {
        const { data: res } = await this.$http.get(`res/borrow/${this.cateId}`, {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      } else {
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
      }
    },
    async getApplyBorrowList () {
      if (this.cateId > 0) {
        const { data: res } = await this.$http.get(`res/applyborrow/${this.cateId}`, {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      } else {
        const { data: res } = await this.$http.get('res/applyborrow', {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      }
    },
    async getUnBorrowList () {
      if (this.cateId > 0) {
        const { data: res } = await this.$http.get(`res/unborrow/${this.cateId}`, {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      } else {
        const { data: res } = await this.$http.get('res/unborrow', {
          params: {
            resname: this.queryParam.resname,
            pagesize: this.queryParam.defaultPageSize,
            pagenum: this.queryParam.defaultCurrent
          }
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.reslist = res.data
        this.pagination.total = res.total
      }
    },
    clickSearch () {
      this.newlist()
    },
    async getOneRes (id) {
      const { data: res } = await this.$http.get(`res/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.Detailslist = res.data
      this.Detailslist.catename = res.data.CateInfo.catename
      this.Detailslist.username = res.data.UserInfo.username
    },
    async showModal (id) {
      await this.getOneRes(id)
      this.visible = true
    },
    handleOk (e) {
      this.visible = false
    },
    handleCancel (e) {
      this.visible = false
    },
    CateChange (value) {
      this.cateId = value
      this.newlist()
    },
    newlist () {
      if (this.status === 1) {
        this.getUnBorrowList()
      } else if (this.status === 2) {
        this.getApplyBorrowList()
      } else if (this.status === 3) {
        this.getBorrowList()
      } else {
        this.getResList()
      }
    },
    statusChange (value) {
      this.status = value
      this.newlist()
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
