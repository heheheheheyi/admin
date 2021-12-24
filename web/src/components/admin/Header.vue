<template>
  <div>
    <a-avatar style="margin: 10px" icon="user"/>
    <a-badge :count="count" dot>
      <a-dropdown>
        <a class="ant-dropdown-link" @click="e => e.preventDefault()">
          {{ showUserNmae }}
          <a-icon type="down"/>
        </a>
        <a-menu slot="overlay">
          <a-menu-item>
            <a-badge :count="count" dot>
              <a type="button" @click="clickmessage">消息</a>
            </a-badge>
            <a-modal
              title="消息"
              :visible="visible"
              @ok="handleOk"
              @cancel="handleCancel"
              :mask="false"
              :footer="null"
            >
              <div v-if="noreadcontent!=null">
                <h3 :style="{ margin: '16px 0' }">
                  未读消息
                </h3>
                <!--                <a-list bordered :data-source="noreadcontent">-->
                <!--                  <a-list-item slot="renderItem" v-for="item in noreadcontent" :key="item.id">-->
                <!--                    {{ item.content }}-->
                <!--                  </a-list-item>-->
                <!--                </a-list>-->
                <a-list item-layout="horizontal" :data-source="noreadcontent">
                  <div v-for="item in noreadcontent" :key="item.id">
                    <a-list-item slot="renderItem">
                      <a-list-item-meta>
                        <span slot="title">{{ item.UserInfo.username }}</span>
                        <a slot="description">{{ item.content }}</a>
                        <a-avatar
                          slot="avatar"
                          src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                        />
                      </a-list-item-meta>
                    </a-list-item>
                  </div>
                </a-list>
              </div>
              <div v-if="contents!=null">
                <h3 :style="{ margin: '16px 0' }">
                  已读消息
                </h3>
                <!--                <a-list bordered :data-source="content">-->
                <!--                  <a-list-item slot="renderItem" v-for="item in content" :key="item.id">-->
                <!--                    {{ item.content }}-->
                <!--                  </a-list-item>-->
                <!--                </a-list>-->
                <a-list item-layout="horizontal" :data-source="contents">
                  <div v-for="item in contents" :key="item.id">
                    <a-list-item slot="renderItem">
                      <a-list-item-meta>
                        <span slot="title">{{ item.UserInfo.username }}</span>
                        <span slot="description">{{ item.content }}</span>
                        <a-avatar
                          slot="avatar"
                          src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                        />
                      </a-list-item-meta>
                    </a-list-item>
                  </div>
                </a-list>
              </div>
            </a-modal>
          </a-menu-item>
          <a-menu-item>
            <a @click="loginout">退出</a>
          </a-menu-item>
        </a-menu>
      </a-dropdown>
    </a-badge>
  </div>
</template>

<script>

export default {
  name: 'Header',
  data () {
    return {
      UserInfo: [],
      Message: [],
      count: 0,
      noreadcontent: [],
      contents: [],
      visible: false,
      confirmLoading: false
    }
  },
  async created () {
    await this.getUserInfo()
    this.getCount()
  },
  methods: {
    loginout () {
      window.sessionStorage.clear()
      this.$router.push('/login')
    },
    async getCount () {
      const { data: res } = await this.$http.get(`chat/count/${this.UserInfo.id}`)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.count = res.count
    },
    clickmessage () {
      this.getmessage()
      this.getnoreadmessage()
      this.visible = true
    },
    async getmessage () {
      const { data: res } = await this.$http.get(`chat/${this.UserInfo.id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.contents = res.data
    },
    async getnoreadmessage () {
      const { data: res } = await this.$http.get(`chat/noread/${this.UserInfo.id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.noreadcontent = res.data
    },
    handleOk (e) {
      this.visible = false
    },
    handleCancel (e) {
      this.getCount()
      this.visible = false
    },
    getUserInfo () {
      this.UserInfo = JSON.parse(window.sessionStorage.getItem('UserInfo'))
    }
  },
  computed: {
    showUserNmae () {
      return this.UserInfo.username
    }
  }
}
</script>

<style scoped>

</style>
