<template>
  <div>
    <a-card>
      <p>资源名：{{ resInfo.resname }}</p>
      <p>分类：{{ resInfo.catename }}</p>
      <p>
        <img :src="resInfo.img" style="width: 120px;height: 100px">
      </p>
      <template>
        <div>
          <a-range-picker @change="onChange" :disabled-date="disabledDate"/>
        </div>
      </template>
      <a-form-model-item>
        <a-button type="danger" style="margin-right: 15px" @click="resOK(resInfo.id)">申请资源</a-button>
        <a-button type="primary" @click="applyCancle">取消</a-button>
      </a-form-model-item>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'
import moment from 'moment'

export default {
  props: ['id'],
  data () {
    return {
      UserInfo: [],
      resInfo: {
        resname: '',
        catename: '',
        img: ''
      },
      Categorylist: [],
      upUrl: Url + 'upload',
      headers: {},
      date: {
        username: '',
        borrow_time: '',
        return_time: '',
        uid: ''
      }
    }
  },
  created () {
    this.getUserInfo()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    this.getResInfo(this.id)
  },
  methods: {
    getUserInfo () {
      this.UserInfo = JSON.parse(window.sessionStorage.getItem('UserInfo'))
    },
    moment,
    async getResInfo (id) {
      const { data: res } = await this.$http.get(`res/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.resInfo = res.data
      this.resInfo.catename = res.data.CateInfo.catename
    },
    async resOK (id) {
      if (this.date.return_time === '') {
        return this.$message.error('请选择日期')
      }
      this.date.uid = this.UserInfo.id
      const { data: res } = await this.$http.post(`res/borrow/${id}`, this.date)
      if (res.status !== 200) return this.$message.error(res.message)
      this.$message.success('申请资源成功')
      this.$router.push('/reslist')
    },
    applyCancle () {
      this.$router.push('/reslist')
    },
    disabledDate (current) {
      // Can not select days before today and today
      return current && current < moment().endOf('day')
    },
    onChange (date, dateString) {
      this.date.borrow_time = dateString[0]
      this.date.return_time = dateString[1]
    }
  }
}
</script>
