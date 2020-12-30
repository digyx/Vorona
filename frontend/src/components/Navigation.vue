<template>
    <div id="navigation">
      <div id="mobile" v-if="this.mobile">
        <img src="../assets/bars-line.svg" @click="displayNavbar">
      </div>

      <div id="navbar" v-if="this.showNav">
        <img src="../assets/times-line.svg" v-if="this.mobile" @click="hideNavbar">
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
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'

export default Vue.extend({
  name: 'SiteIndex',
  
  data: function() {
    return {
      articles: ["Mephisto", "Leviathan"],
      mobile: false,
      showNav: true,
    }
  },

  created() {
    axios.get(`${this.$store.getters.getURL}:8000/articles`)
    .then((res) => this.articles = res.data)
    .catch((err) => alert(err))

    window.addEventListener("resize", this.checkResize)
    window.addEventListener("locationchange", this.checkResize)
  },

  methods: {
    home() {
      this.$router.push("/")
    },
    checkResize() {
      if (window.innerWidth < 1380) {
        this.mobile = true
        this.showNav = false
      } else {
        this.mobile = false
        this.showNav = true
      }
    },
    displayNavbar() {
      this.showNav = true
    },
    hideNavbar() {
      this.showNav = false
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

@media screen and (max-width: 1380px) {
  #navbar {
    position: fixed;
    background: white;
    box-shadow: 0 0 10px #000;
  }

  #mobile {
    height: 3em;
    width: 100%;
    padding: 0 0.5em;
    position: fixed;

    display: flex;
    justify-content: left;
    align-items: center;

    background: white;
    box-shadow: 0 0 3px #000;
  }
}
</style>