<template>
  <div id="article" v-show="this.show" class="content">
    <h1>{{ this.title }}</h1>
    <h2>{{ this.subtitle }}</h2>
    <div id="container">
      <div v-html="body" id="content"></div>
      <div id="sidebar">
        <img :src="'https://cdn.vorona.gg/'+ this.title + '.jpg'" id="headshot" onerror="this.style.display='none'">
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

export default Vue.extend({
  name: 'Article',

  data: function() {
    return {
      show: false,
      title: this.$route.params.title,
      subtitle: "Placeholder",
      sidebar: {Hello: "World"},
      body: "Placeholder",
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
      window.dispatchEvent(new Event("locationchange"))
      document.title = `Gnezdo Vorona - ${this.title}`

      axios.get(`${this.$store.getters.getURL}:8000/articles/${this.title}`)
      .then((res) => {
        this.subtitle = res.data.Subtitle
        this.sidebar = res.data.Sidebar
        this.body = res.data.Body
        this.show = true
      })
      .catch((err)=> alert(err))
    }
  }
});
</script>

<style>
p {
  text-align: justify;
}

#container {
  display: flex;
}

#content {
  margin: 1em 5em 1em 2em;
}

#sidebar {
  width: 320px;
  height: 100%;
  margin: 2em;
  padding: 1em;
  flex-shrink: 0;
  
  border-style: solid;
  border-width: 1px;
  border-radius: 0.5em;
}

@media screen and (max-width: 1040px){
  #container {
    flex-direction: column-reverse;
    align-items: center;
  }

  .flex-content {
    flex-direction: column;
  }

  #sidebar {
    width: calc(100% - 2em);
    max-width: calc(320px + 2em);
    margin: 1em;
    padding: 1em;
  }

  #headshot {
    max-width: 320px;
    width: 100%;
    height: auto;
  }

  #content {
    margin: 1em;
  }
}
</style>