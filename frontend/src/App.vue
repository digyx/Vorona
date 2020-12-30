<template>
  <div id="app">
    <div id="nav">
      <Navbar/>
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Navbar from '@/components/Navigation.vue'
export default Vue.extend({
  name: "App",
  components: {
    Navbar,
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;

  height: 100%;
  width: 100%;
  display: flex;
}

#nav {
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

.content {
  margin: 1em 1em 1em 19em;
}

@media screen and (max-width: 1380px) {
  .content {
    margin: 5em 2em 2em 2em;
  }
}
</style>
