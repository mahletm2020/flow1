<template>
    <div>
      <h1>Login</h1>
      <form @submit.prevent="handleLogin">
        <input v-model="email" type="email" placeholder="Email" />
        <input v-model="password" type="password" placeholder="Password" />
        <button type="submit">Login</button>
      </form>
      <p v-if="error">{{ error }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  const email = ref('')
  const password = ref('')
  const error = ref('')
  
  const axios = useAxios()
  
  async function handleLogin() {
    try {
      const res = await axios.post('/login', {
        email: email.value,
        password: password.value,
      })
  
      const token = res.data.token
      localStorage.setItem('token', token)
  
      // redirect to dashboard
      await navigateTo('/dashboard')
    } catch (err) {
      console.error(err)
      error.value = 'Login failed: ' + err?.response?.data?.message || err.message
    }
  }
  </script>
  