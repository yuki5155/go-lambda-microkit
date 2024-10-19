<template>
    <div class="login">
      <h1>ログイン</h1>
      <LoginForm @login="handleLogin" :isLoading="isLoading" />
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue'
  import { useRouter } from 'vue-router'
  import LoginForm from '@/components/LoginForm.vue'
  import authService, { LoginCredentials } from '@/services/authService'
  import { setAuthToken, setUserData } from '@/utils/utils'
  
  export default defineComponent({
    name: 'LoginView',
    components: {
      LoginForm
    },
    setup() {
      const errorMessage = ref('')
      const isLoading = ref(false)
      const router = useRouter()
  
      const handleLogin = async (credentials: LoginCredentials) => {
        isLoading.value = true
        errorMessage.value = ''
  
        try {
          const response = await authService.login(credentials)
          setAuthToken(response.token)
          setUserData(response.user)
          router.push('/') // ログイン後のリダイレクト先
        } catch (error) {
          if (error instanceof Error) {
            errorMessage.value = error.message
          } else {
            errorMessage.value = 'ログイン中に予期せぬエラーが発生しました'
          }
        } finally {
          isLoading.value = false
        }
      }
  
      return {
        errorMessage,
        isLoading,
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