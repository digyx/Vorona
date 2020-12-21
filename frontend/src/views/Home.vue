<template>
  <div id="home">
    <h1>The World of Gnezdo Vorona</h1>
    <hr>
    Please use the navigation bar to the left to view articles.
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Home',
  created() {
    document.title = "Gnezdo Vorona"
  }
});
</script>

<style scoped>
#home {
  margin: 1em 1em 1em 19em;
  text-align: justify;
}
</style>