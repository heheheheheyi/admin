<template>
  <div>
    <a-card>
      <a-form-model :model="cateInfo" ref="cateInfoRef" :rules="cateInfoRules" :hide-required-mark="true">
        <a-form-model-item label="分类名" prop="catename">
          <a-input style="width: 300px" v-model="cateInfo.catename"></a-input>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="danger" style="margin-right: 15px" @click="resOK(cateInfo.id)">{{
              cateInfo.id ? '更新' : '提交'
            }}
          </a-button>
          <a-button type="primary" @click="addCancle">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>

export default {
  props: ['id'],
  data () {
    return {
      cateInfo: {
        id: 0,
        catename: ''
      },
      headers: {},
      cateInfoRules: {
        catename: [{
          required: true,
          message: '请输入分类名',
          trigger: 'blur'
        }]
      }
    }
  },
  created () {
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getCateInfo(this.id)
    }
  },
  methods: {
    async getCateInfo (id) {
      const { data: res } = await this.$http.get(`cate/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.cateInfo = res.data
    },
    resOK (id) {
      this.$refs.cateInfoRef.validate(async (valid) => {
        if (id === 0) {
          const { data: res } = await this.$http.post('cate/add', this.cateInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('添加分类成功')
          this.$router.push('/catelist')
        } else {
          const { data: res } = await this.$http.put(`cate/${id}`, this.cateInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('更新分类成功')
          this.$router.push('/catelist')
        }
      })
    },
    addCancle () {
      this.$refs.cateInfoRef.resetFields()
      this.$router.push('/catelist')
    }
  }
}
</script>
