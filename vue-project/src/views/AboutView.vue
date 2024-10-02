<template>
  <div class="about">
    <h1>{{msg}}</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      msg: 'Hello, Vue!'
    }
  }
}
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
