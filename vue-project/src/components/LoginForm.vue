<template>
    <form @submit.prevent="handleSubmit" class="login-form">
      <div class="form-group">
        <label for="email">メールアドレス:</label>
        <input type="email" id="email" v-model="email" required class="form-control">
      </div>
      <div class="form-group">
        <label for="password">パスワード:</label>
        <input type="password" id="password" v-model="password" required class="form-control">
      </div>
      <button type="submit" class="btn btn-primary">ログイン</button>
    </form>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue'
  
  export default defineComponent({
    name: 'LoginForm',
    emits: ['login'],
    setup(_, { emit }) {
      const email = ref('')
      const password = ref('')
  
      const handleSubmit = () => {
        emit('login', { username: email.value, password: password.value })
      }
  
      return {
        email,
        password,
        handleSubmit
      }
    }
  })
  </script>
  
  <style scoped>
  .login-form {
    max-width: 300px;
    margin: 0 auto;
  }
  .form-group {
    margin-bottom: 15px;
  }
  .form-control {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  .btn {
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .btn:hover {
    background-color: #0056b3;
  }
  </style>