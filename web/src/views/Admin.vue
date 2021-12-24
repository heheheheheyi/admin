<template>
  <a-layout class="container">
    <a-layout-sider>
      <Nav></Nav>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="headerbtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content><router-view :key="$route.path"></router-view></a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import Header from '../components/admin/Header'
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'

export default {
  components: {
    Nav,
    Footer,
    Header
  }
}
</script>
<style scoped>
.container {
  height: 100%;
}

.headerbtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

</style>
