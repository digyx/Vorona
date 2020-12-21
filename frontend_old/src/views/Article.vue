<template>
  <div id="article">
    <h1>{{ this.title }}</h1>
    <h2>{{ this.subtitle }}</h2>
    <hr>
    {{ this.body }}
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';
// import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

export default defineComponent({
  name: 'Article',

  data: function() {
    return {
      subtitle: String,
      sidebar: Object,
      body: String,
    }
  },

  created() {
    axios.get(`https://localhost:8000/article/${this.$route.params.title}`)
    .then((res) => {
      this.subtitle = res.data.Subtitle
      this.sidebar = res.data.Didebar
      this.body = res.data.Body
    })
    .catch((err)=> alert(err))
  }
});
</script>
