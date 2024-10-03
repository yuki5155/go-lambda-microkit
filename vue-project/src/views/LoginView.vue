<template>
    <div class="login">
      <h1>ログイン</h1>
      <LoginForm @login="handleLogin" />
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue'
  import LoginForm from '@/components/LoginForm.vue'
  
  export default defineComponent({
    name: 'LoginView',
    components: {
      LoginForm
    },
    setup() {
      const errorMessage = ref('')
  
      const handleLogin = (credentials: { email: string; password: string }) => {
        // ここにログインロジックを実装します
        console.log('ログイン試行:', credentials.email, credentials.password)
        // 例: APIリクエストを送信し、応答を処理します
        // 成功した場合はユーザーをリダイレクトし、失敗した場合はエラーメッセージを表示します
        errorMessage.value = '' // リセット
        // 仮のエラーメッセージ
        errorMessage.value = 'ログインに失敗しました。認証情報を確認してください。'
      }
  
      return {
        errorMessage,
        handleLogin
      }
    }
  })
  </script>
  
  <style scoped>
  .login {
    max-width: 300px;
    margin: 0 auto;
  }
  .error-message {
    color: red;
    margin-top: 10px;
  }
  </style>