<template>
  <div id="app">

    <navbar></navbar>
    <router-view :key="key"/>
    <img alt="Vue logo" src="./assets/logo.png">

    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import navbar from "@/components/navbar";
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    navbar
  },
  computed: {
    key() {
      return this.$route.name !== undefined
          ? this.$route.name + new Date()
          : this.$route + new Date();
    }
  }
}
</script>

<style>
body {
  margin: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
