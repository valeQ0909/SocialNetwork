<template>
  <NavBar></NavBar>
  <router-view></router-view>
</template>

<script>
import NavBar from './components/NavBar.vue';

export default {
  components: {
    NavBar,
  },
  setup() {

  }
}
</script>

<style>
body{
  /*background-image: url("../src/assets/images/background.jpg");*/
  background-size: cover;
  width: 100%;
}
</style>
