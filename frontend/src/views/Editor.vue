<template>
  <div id="editor" class="content">
    <h1>Article Editor</h1>
    <hr style="width: 100%;">
    <div>
      <select name="article-list" id="article-list" v-model="articleSelected" @change="changeEditor">
        <option
          :value="articleName"
          :key="articleName.key"
          v-for="articleName in articleList"
        >
          {{ articleName }}
        </option>
        <option value="new">+ New Article</option>
      </select>
    </div>
    <div id="editor-container">
      <h3>Title</h3>
      <input type="text" v-model="article.Title" />
      
      <h3>Subtitle</h3>
      <input type="text" v-model="article.Subtitle" />
      
      <h3>Body</h3>
      <textarea id="body-editor" v-model="article.Body"></textarea>

      <button @click="saveArticle()">Save Changes</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
  name: "Account",
  data: function () {
    return {
      articleList: [],
      articleSelected: "",
      article: {
        ID: "",
        Title: "",
        Subtitle: "",
        Sidebar: {},
        Body: ""
      }
    };
  },

  created() {
    document.title = "Gnezdo Vorona - Editor";

    axios
      .get(`${this.$store.getters.getURL}:8000/articles`)
      .then((res) => (this.articleList = res.data));
  },

  methods: {
    changeEditor() {
      if (this.articleSelected == "new") {
        this.article.ID = "new"
        this.article.Title = ""
        this.article.Subtitle = ""
        this.article.Sidebar = {}
        this.article.Body = ""

        return
      }

      axios.get(`${this.$store.getters.getURL}:8000/articles/${this.articleSelected}?format=markdown`)
      .then((res) => {
        this.article = res.data
      })
    },

    saveArticle() {
      const url = `${this.$store.getters.getURL}:8000/articles/${this.article.ID}`
      const data = {
        Title: this.article.Title,
        Subtitle: this.article.Subtitle,
        Sidebar: this.article.Sidebar,
        Body: this.article.Body
      }

      axios.post(url, data)
      .then(() => {
        this.$forceUpdate()
      })
      .catch((err) => {
        switch (err.response.status) {
          case 409:
            alert("Article with this name already exists.")
            break
          
          case 500:
            alert("Error adding to database")
            break

          case 401:
            alert("You shouldn't be here...")
            break
            
          default:
            alert(err)
            break
        }
      })
    }
  }
});
</script>

<style scoped>
#editor {
  height: 90%;
  width: 100%;

  display: flex;
  flex-direction: column;
}

#article-list {
  width: 100%;
  max-width: 25em;
  padding: 0.5em;
}

#editor-container {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}

input[type=text], button {
  max-width: 24em;
  padding: 0.5em;
}

textarea {
  max-width: 80em;
  height: 100%;
  min-height: 24em;

  margin-bottom: 2em;
}
</style>