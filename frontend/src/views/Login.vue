<template>
  <div id="login" class="content">
    <h2>Login</h2>

    <input type="text" class="login-field" v-model="email" placeholder="Email Address">
    <input type="password" class="login-field" v-model="password" placeholder="Password">
    <input type="password" class="login-field" v-model="password2" placeholder="Confirm Password" v-if="accountCreation">

    <button @click="login()" class="login-field" v-if="!accountCreation">Login</button>
    <button @click="createAccount()" class="login-field" v-if="accountCreation">Create Account</button>
    
    <br>

    <p v-if="!accountCreation" @click="toggleAccountCreation">
      <i>Create account instead</i>
    </p>
    <p v-if="accountCreation" @click="toggleAccountCreation">
      <i>Login instead</i>
    </p>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import axios from 'axios'

export default Vue.extend({
  name: 'Login',
  data: function() {
    return {
      email: "",
      password: "",
      password2: "",
      accountCreation: false
    }
  },

  created() {
    document.title = "Gnezdo Vorona - Login"
  },

  methods: {
    login() {
      const url = `${this.$store.getters.getURL}:8000/auth`
      const data = {
        Email: this.email,
        Password: this.password
      }

      axios.post(url, data)
      .then((res) => {
        if (res.status === 200) {
          this.$store.commit('login')
          this.$router.push("/account")
          return
        }
      })
      .catch((err) => {
        if (err.response.status === 401) {
          alert("Incorrect Email and/or Password.")
        } else {
          alert(err)
        }
      })
    },

    createAccount() {
      if (this.password !== this.password2) {
        alert("Passwords do not match!")
        return
      }

      const url = `${this.$store.getters.getURL}:8000/auth`
      const data = {
        Email: this.email,
        Password: this.password
      }

      axios.put(url, data)
      .then(() => {
        this.$store.commit('login')
      })
      .catch((err) => alert(err))
    },

    toggleAccountCreation() {
      this.accountCreation = !this.accountCreation
    }
  }
});
</script>

<style>
#login {
  width: 100%;
  height: 50vh;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.field-container {
  width: 85%;
  max-width: 24em;

  margin: 1em;
}

.login-field {
  width: calc(100% - 4em);
  max-width: 24em;

  padding: 1em;
  margin: 1em;

  border-style: solid;
  border-radius: 0.5em;
}
</style>