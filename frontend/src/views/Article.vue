<template>
  <div id="article">
    <h1>{{ this.title }}</h1>
    <h2>{{ this.subtitle }}</h2>
    <hr>
    <div id="container">
      <div v-html="body" id="content"></div>
      <div id="sidebar">
        <div v-for="(value, key) in this.sidebar" :key="key.id">
          <p>
            <b>{{ key }}</b>
            <br>
            {{ value }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios';
import { nextTick } from 'vue/types/umd';

export default Vue.extend({
  name: 'Article',

  data: function() {
    return {
      title: this.$route.params.title,
      subtitle: "|",
      sidebar: {Hello: "World"},
      body: "",
    }
  },

  created() {
    this.getPage()
  },

  beforeRouteUpdate(to, from, next) {
    this.title = to.params["title"]
    this.getPage()

    return next()
  },

  methods: {
    getPage() {
      axios.get(`${this.$store.getters.getURL}:8000/articles/${this.title}`)
    .then((res) => {
      this.subtitle = res.data.Subtitle
      this.sidebar = res.data.Sidebar
      this.body = res.data.Body
    })
    .catch((err)=> alert(err))
    }
  }
});
</script>

<style scoped>
#article {
  margin: 1em 1em 1em 19em;
  text-align: justify;
}

#container {
  display: flex;
}

#content {
  margin: 1em 5em 1em 2em;
}

#sidebar {
  width: 15em;
  margin: 2em;
  padding: 3em;
  flex-shrink: 0;
}
</style>