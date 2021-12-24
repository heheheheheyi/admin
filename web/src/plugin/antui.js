import Vue from 'vue'
import {
  Avatar,
  Button,
  Card,
  Col,
  ConfigProvider,
  DatePicker,
  Dropdown,
  FormModel,
  Icon,
  Input,
  Layout,
  Menu,
  message,
  Modal,
  Row,
  Select,
  Switch,
  Table,
  Upload,
  Descriptions,
  Badge,
  List
} from 'ant-design-vue'

message.config({
  top: '50px',
  duration: 2,
  maxCount: 3
})

Vue.prototype.$message = message
Vue.prototype.$confirm = Modal.confirm

Vue.use(Button)
Vue.use(Layout)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Row)
Vue.use(Col)
Vue.use(Table)
Vue.use(ConfigProvider)
Vue.use(Modal)
Vue.use(Dropdown)
Vue.use(Select)
Vue.use(Switch)
Vue.use(Upload)
Vue.use(Avatar)
Vue.use(DatePicker)
Vue.use(Descriptions)
Vue.use(Badge)
Vue.use(List)
