<template>
    <div id="navbar">
      <h2 @click="home()" style="cursor: pointer">
        Gnezdo Vorona
      </h2>
      <div v-for="article in this.articles" :key="article.id">
        <!-- Change this to be custom code for both style and the ability to reload the page at the same time -->
        <router-link :to="{ name: 'articles', params: { title: article }}" class="nav-item">
          {{ article }}
        </router-link>
      </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'

export default Vue.extend({
  name: 'SiteIndex',
  
  data: function() {
    return {
      articles: ["Mephisto", "Leviathan"]
    }
  },

  created() {
    axios.get(`${this.$store.getters.getURL}:8000/articles`)
    .then((res) => this.articles = res.data)
    .catch((err) => alert(err))
  },

  methods: {
    home() {
      this.$router.push("/")
    }
  }
});
</script>

<style scoped>
#navbar {
  width: 14em;
  height: 100%;
  padding: 2em;
  text-align: left;
}

.nav-item {
  margin: 1em;
}
</style>