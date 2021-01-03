<template>
  <div id="account" class="content">
    <div>
        <h1>Account</h1>
        Welcome to your account page!!
    </div>
    
    <button @click="logout()" id="logout">
        Logout
    </button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
// import axios from 'axios'

export default Vue.extend({
  name: 'Account',
  data: function() {
    return {
      email: "",
    }
  },

  created() {
    document.title = "Gnezdo Vorona - Account"
  },

  methods: {
      logout() {
          this.$store.commit('logout')
      }
  }
});
</script>

<style scoped>
#account {
    display: flex;
    flex-direction: column;
}

#logout {
    width: 10em;

    position: fixed;
    bottom: 2em;
    left: calc(50vw - 5em);

    padding: 1em;

    border-style: solid;
    border-radius: 0.5em;
}
</style>