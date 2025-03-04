<template>
  <q-page class="q-pa-md flex flex-center column">
    <div class="text-h5 text-center q-mb-md">Admin Page</div>
    <div>
      <p v-if="loading">⏳ Loading...</p>
      <p v-else-if="error" class="text-red">⚠️ {{ error }}</p>
      <p v-else-if="currentUser">👤 {{ currentUser.first_name }} {{ currentUser.last_name }}</p>
      <p v-else>No user data</p>
    </div>
  </q-page>
</template>
<script setup lang="ts">

import { onMounted } from 'vue'
import { useUser } from '../../composable/useUser'

const { currentUser, loading, error, initializeUser } = useUser()

onMounted(() => {
  void initializeUser()
})
</script>
