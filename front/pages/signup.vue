<template>
    <div>
      <h1>Sign Up</h1>
      <form @submit.prevent="handleSignup">
        <input v-model="name" type="text" placeholder="Full Name" required />
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Password" required />
        <button type="submit">Create Account</button>
      </form>
  
      <p v-if="error">{{ error }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  const name = ref('')
  const email = ref('')
  const password = ref('')
  const error = ref('')
  
  const axios = useAxios()
  
  async function handleSignup() {
    try {
      const res = await axios.post('/signup', {
        name: name.value,
        email: email.value,
        password: password.value,
      })
  
      const token = res.data.token
      localStorage.setItem('token', token)
  
      // redirect to dashboard
      await navigateTo('/dashboard')
    } catch (err) {
      error.value = 'Signup failed: ' + (err?.response?.data?.message || err.message)
    }
  }
  </script>
  