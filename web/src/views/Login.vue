<template>
  <div class="container" :style="{backgroundImage:`url(${img})`}">
    <div class="loginBox">
      <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="Username">
            <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)"/>
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input-password v-model="formdata.password" placeholder="Password"
                            type="password" v-on:keyup.enter="login"
          >
            <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)"/>
          </a-input-password>
        </a-form-model-item>

        <a-form-model-item>
          <a-button class="loginButton" type="primary" style="margin: 10px" @click="login">登录</a-button>
          <a-button class="resButton" type="link" style="margin-right: 15px" @click="showModal()">注册</a-button>
          <a-modal
            title="欢迎注册"
            :visible="visible"
            @ok="handleOk"
            @cancel="handleCancel"
            :mask="false"
          >
            <a-form-model :model="User" ref="UserInfoRef" :rules="UserRules" :hide-required-mark="true">
              <a-form-model-item label="用户名" prop="username">
                <a-input style="width: 300px" v-model="User.username"></a-input>
              </a-form-model-item>
              <a-form-model-item label="密码" prop="password">
                <a-input type="password" style="width: 300px" v-model="User.password"></a-input>
              </a-form-model-item>
              <a-form-model-item label="确认密码" prop="checkpass">
                <a-input type="password" style="width: 300px" v-model="User.checkpass"></a-input>
              </a-form-model-item>
              <a-form-model-item label="权限" prop="username">
                <a-select v-model="User.role" style="width: 200px" placeholder="请选择权限" @change="roleChange">
                  <a-select-option v-for="item in Rolelist" :key="item.role" :value="item.role">{{ item.name }}
                  </a-select-option>
                </a-select>
              </a-form-model-item>
            </a-form-model>
          </a-modal>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      UserInfo: [],
      User: {
        username: '',
        password: '',
        checkpass: '',
        role: ''
      },
      Rolelist: [
        {
          role: 1,
          name: '管理员'
        },
        {
          role: 2,
          name: '普通用户'
        }
      ],
      formdata: {
        username: '',
        password: ''
      },
      UserRules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          },
          {
            min: 4,
            max: 12,
            message: '长度必须在4-12个字符之间',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          },
          {
            min: 6,
            max: 20,
            message: '长度必须在6-20个字符之间',
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          },
          {
            validator: (rule, value, callback) => {
              if (this.User.checkpass !== this.User.password) {
                callback(new Error('两次输入密码不同'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        role: [
          {
            required: true,
            message: '请选择权限',
            trigger: 'blur'
          }
        ]
      },
      rules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          },
          {
            min: 4,
            max: 12,
            message: '长度必须在4-12个字符之间',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          },
          {
            min: 6,
            max: 20,
            message: '长度必须在6-20个字符之间',
            trigger: 'blur'
          }
        ]
      },
      img: require('../static/moutient.jpg'),
      visible: false
    }
  },
  methods: {
    resetFrom () {
      this.$refs.loginFormRef.resetFields()
    },
    login () {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return this.$message.error('输入信息错误，请重新输入')
        const { data: res } = await this.$http.post('/login', this.formdata)
        if (res.status !== 200) return this.$message.error(res.message)
        this.UserInfo = res.data
        window.sessionStorage.setItem('UserInfo', JSON.stringify(this.UserInfo))
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('reslist')
      })
    },
    showModal () {
      this.visible = true
    },
    res () {
      this.$refs.UserInfoRef.validate(async valid => {
        if (!valid) return this.$message.error('输入信息错误，请重新输入')
        const { data: res } = await this.$http.post('/user/add', this.User)
        if (res.status !== 200) return this.$message.error(res.message)
        this.visible = false
        return this.$message.success('注册成功')
      })
    },
    handleOk (e) {
      this.res()
    },
    handleCancel (e) {
      this.visible = false
    },
    roleChange (value) {
      this.UserInfo.role = value
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  background-size: cover;
}

.loginBox {
  width: 400px;
  height: 220px;
  position: absolute;
  top: 45%;
  left: 70%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
  background: #fdfdfd;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: -4%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginButton {
  left: -10px;
  width: 80%;
}

.resButton {
  left: -10px;
  width: 0%;
}

</style>
