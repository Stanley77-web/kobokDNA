<template>
  <img alt="Kobok DNA logo" src="./assets/kobok_DNA_logo_resized.png">
  
  <inputPage/>
</template>

<script>
import inputPage from './components/inputPage.vue'

export default {
  name: 'App',
  components: {
    inputPage
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
