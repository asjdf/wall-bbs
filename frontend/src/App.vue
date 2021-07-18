<template>
  <div id="app">

    <navbar></navbar>
    <router-view :key="key"/>
  </div>
</template>

<script>
import navbar from "@/components/navbar";

export default {
  name: 'App',
  components: {
    navbar
  },
  computed: {
    key() {
      return this.$route.name !== undefined
          ? this.$route.name + new Date()
          : this.$route + new Date();
    }
  },
  created() {
    this.$store.state.hasToken = localStorage.hasToken;
    this.$store.state.right = localStorage.right;
    this.$store.state.uid = localStorage.uid;
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
