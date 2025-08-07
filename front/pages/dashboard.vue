<template>
    <div>
      <h1>Dashboard</h1>
      <p>Welcome, you are logged in.</p>
      <p>Token: {{ localStorage.getItem('token') }}</p>

      <button @click="logout">Logout</button>
    </div>
  </template>
  
  <script setup>
  function logout() {
    localStorage.removeItem('token')
    navigateTo('/login')
  }
  </script>
  