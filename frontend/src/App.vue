<template>
  <div id="app">

    <navbar></navbar>
    <router-view/>
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
