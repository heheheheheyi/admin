<template>
  <div>
    <a-card>
      <a-form-model :model="resInfo" ref="resInfoRef" :rules="resInfoRules" :hide-required-mark="true">
        <a-form-model-item label="资源名" prop="resname">
          <a-input style="width: 300px" v-model="resInfo.resname"></a-input>
        </a-form-model-item>
        <a-form-model-item label="分类" prop="cid">
          <a-select v-model="resInfo.cid" style="width: 200px" placeholder="请选择分类" @change="CateChange">
            <a-select-option v-for="item in Categorylist" :key="item.id" :value="item.id">{{ item.catename }}
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="图片" prop="img">
          <a-upload
            name="file"
            :multiple="false"
            :action="upUrl"
            :headers="headers"
            @change="upChange"
            list-type="picture"
            :default="filelist"
          >
            <a-button>
              <a-icon type="upload"/>
              点击上传
            </a-button>
            <br>
            <template v-if="id">
              <img :src="resInfo.img" style="width: 120px;height: 100px">
            </template>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="danger" style="margin-right: 15px" @click="resOK(resInfo.id)">{{
              resInfo.id ? '更新' : '提交'
            }}
          </a-button>
          <a-button type="primary" @click="addCancle">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'

export default {
  props: ['id'],
  data () {
    return {
      resInfo: {
        id: 0,
        resname: '',
        cid: undefined,
        img: ''
      },
      Categorylist: [],
      upUrl: Url + 'upload',
      headers: {},
      filelist: [],
      resInfoRules: {
        resname: [{
          required: true,
          message: '请输入资源名',
          trigger: 'blur'
        }],
        cid: [{
          required: true,
          message: '请选择分类',
          trigger: 'change'
        }],
        img: [{
          required: true,
          message: '请选择缩略图',
          trigger: 'blur'
        }]
      }
    }
  },
  created () {
    this.getCateList()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getResInfo(this.id)
    }
  },
  methods: {
    async getResInfo (id) {
      const { data: res } = await this.$http.get(`res/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.resInfo = res.data
    },
    async getCateList () {
      const { data: res } = await this.$http.get('cate')
      if (res.status !== 200) return this.$message.error(res.message)
      this.Categorylist = res.data
    },
    CateChange (value) {
      this.resInfo.cid = value
    },
    upChange (info) {
      if (info.file.status !== 'uploading') {
      }
      if (info.file.status === 'done') {
        this.$message.success('上传成功')
        const imgUrl = info.file.response.url
        this.resInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error('上传失败')
      }
    },
    resOK (id) {
      this.$refs.resInfoRef.validate(async (valid) => {
        if (id === 0) {
          const { data: res } = await this.$http.post('res/add', this.resInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('添加资源成功')
          this.$router.push('/reslist')
        } else {
          const { data: res } = await this.$http.put(`res/${id}`, this.resInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('更新资源成功')
          this.$router.push('/reslist')
        }
      })
    },
    addCancle () {
      this.$refs.resInfoRef.resetFields()
      this.$router.push('/reslist')
    }
  }
}
</script>
