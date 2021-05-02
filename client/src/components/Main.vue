<template>
  <h1>{{ title }}</h1>

 <button @click="httpOne">HTTP/1.1</button>
  <p>
    {{ responseHttpOne }}
  </p>

 <button @click="httpTwo">HTTP/2</button>
  <p>
    {{ responseHttpTwo }}
  </p>

  <button @click="httpThree">HTTP/3</button>
  <p>
    {{ responseHttpThree }}
  </p>

  <button @click="httpThreeNginx">HTTP/3 (NGINX)</button>
  <p>
    {{ responseHttpThreeNginx }}
  </p>
</template>

<script lang="ts">
const httpOne = 'http://localhost:6122';
const httpTwo = 'https://localhost:6123';
const httpThree = 'https://localhost:6124';
const httpThreeNginx = 'https://localhost:6125';

import axios from 'axios';
import { ref, defineComponent } from 'vue'
export default defineComponent({
  name: 'Main',
  props: {
    title: {
      type: String,
      required: true
    }
  },
  data: () => {
    return {
        responseHttpOne: "",
        responseHttpTwo: "",
        responseHttpThree: "",
        responseHttpThreeNginx: "",
    }
  },
  methods: {
   async httpOne() {
      this.responseHttpOne = await axios.get(httpOne).then(res => res.data)
   },
   async httpTwo() {
      this.responseHttpTwo = await axios.get(httpTwo).then(res => res.data)
   },
   async httpThree() {
      this.responseHttpThree = await axios.get(httpThree).then(res => res.data)
   },
   async httpThreeNginx() {
      this.responseHttpThreeNginx = await axios.get(httpThreeNginx).then(res => res.data)
   },
  }
})
</script>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
